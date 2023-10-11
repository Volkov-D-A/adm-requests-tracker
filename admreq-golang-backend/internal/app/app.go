package app

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	tsr "github.com/volkov-d-a/adm-requests-tracker/internal/generated"
	"github.com/volkov-d-a/adm-requests-tracker/pkg/closer"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type App struct {
	httpServer      *http.Server
	grpcServer      *grpc.Server
	ServiceProvider *serviceProvider
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}
	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (a *App) Run() error {
	a.ServiceProvider.Logger.Info("Starting GRPC Server at:", "addr", a.ServiceProvider.Config.GrpcServer.Address)
	go func() {
		err := a.runGRPCServer()
		if err != nil {
			a.ServiceProvider.Logger.Error("Error while handling GRPC requests:", "err", err.Error())
		}
	}()
	a.ServiceProvider.Logger.Info("Starting GRPC GW Server at:", "addr", a.ServiceProvider.Config.GrpcGw.Address)
	go func() {
		err := a.runHttpServer()
		if err != nil {
			a.ServiceProvider.Logger.Error("Error while handling GRPC GW requests:", "err", err.Error())
		}
	}()
	return nil
}

func (a *App) GracefulStop() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	err := a.ServiceProvider.Closer.Close(ctx)
	if err != nil {
		return fmt.Errorf("Eror while stopping app: %v", err)
	}
	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(ctx context.Context) error{
		a.initServiceProvider,
		a.initGRPCServer,
		a.initHttpServer,
	}
	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.ServiceProvider = newServiceProvider()
	err := a.ServiceProvider.SetConfig()
	if err != nil {
		return err
	}
	err = a.ServiceProvider.SetLogger()
	if err != nil {
		return err
	}
	a.ServiceProvider.Closer = &closer.Closer{}
	a.ServiceProvider.Logger.Info("Service provider initialized")
	return nil
}

func (a *App) initGRPCServer(_ context.Context) error {
	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	reflection.Register(a.grpcServer)
	tsr.RegisterTsrServiceServer(a.grpcServer, a.ServiceProvider.TsrImplement())
	a.ServiceProvider.Logger.Info("GRPC Server initialized")
	return nil
}

func (a *App) runGRPCServer() error {
	list, err := net.Listen("tcp", a.ServiceProvider.Config.GrpcServer.Address)
	if err != nil {
		return err
	}

	err = a.grpcServer.Serve(list)
	if err != nil {
		return err
	}
	a.ServiceProvider.Closer.Add(a.stopGRPCServer)
	return nil
}

func (a *App) stopGRPCServer(ctx context.Context) error {
	ok := make(chan struct{})

	go func() {
		a.grpcServer.GracefulStop()
		close(ok)
	}()

	select {
	case <-ok:
		return nil
	case <-ctx.Done():
		a.grpcServer.Stop()
		return ctx.Err()
	}

}

func (a *App) initHttpServer(ctx context.Context) error {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := tsr.RegisterTsrServiceHandlerFromEndpoint(ctx, mux, a.ServiceProvider.Config.GrpcServer.Address, opts)
	if err != nil {
		return err
	}
	a.httpServer = &http.Server{
		Addr:           a.ServiceProvider.Config.GrpcGw.Address,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		Handler:        mux,
	}
	return nil
}

func (a *App) runHttpServer() error {
	err := a.httpServer.ListenAndServe()
	if err != nil {
		return fmt.Errorf("Error listening http server: %v", err)
	}
	a.ServiceProvider.Closer.Add(a.httpServer.Shutdown)
	return nil
}
