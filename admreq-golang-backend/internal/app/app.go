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
	grpcUserServer  *grpc.Server
	grpcTsrServer   *grpc.Server
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
	a.ServiceProvider.Logger.Info("Starting GRPC Servers at:", "addr", a.ServiceProvider.Config.GrpcUserServer.Address, "addr", a.ServiceProvider.Config.GrpcTsrServer.Address)
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

func (a *App) initServiceProvider(ctx context.Context) error {
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
	err = a.ServiceProvider.SetDB()
	if err != nil {
		return err
	}
	a.ServiceProvider.Logger.Info("Service provider initialized")
	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	a.grpcUserServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	reflection.Register(a.grpcUserServer)
	tsr.RegisterUserServiceServer(a.grpcUserServer, a.ServiceProvider.UserApi())

	a.grpcTsrServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	reflection.Register(a.grpcTsrServer)
	tsr.RegisterTSRServiceServer(a.grpcTsrServer, a.ServiceProvider.TSRApi())

	a.ServiceProvider.Logger.Info("GRPC Server initialized")
	return nil
}

func (a *App) runGRPCServer() error {
	luser, err := net.Listen("tcp", a.ServiceProvider.Config.GrpcUserServer.Address)
	if err != nil {
		return err
	}

	err = a.grpcUserServer.Serve(luser)
	if err != nil {
		return err
	}
	a.ServiceProvider.Closer.Add(a.stopUserGRPCServer)

	ltsr, err := net.Listen("tcp", a.ServiceProvider.Config.GrpcTsrServer.Address)
	if err != nil {
		return err
	}

	err = a.grpcTsrServer.Serve(ltsr)
	if err != nil {
		return err
	}

	a.ServiceProvider.Closer.Add(a.stopTSRGRPCServer)
	return nil
}

func (a *App) stopUserGRPCServer(ctx context.Context) error {
	ok := make(chan struct{})
	go func() {
		a.grpcUserServer.GracefulStop()
		close(ok)
	}()

	select {
	case <-ok:
		return nil
	case <-ctx.Done():
		a.grpcUserServer.Stop()
		return ctx.Err()
	}

}

func (a *App) stopTSRGRPCServer(ctx context.Context) error {
	ok := make(chan struct{})
	go func() {
		a.grpcTsrServer.GracefulStop()
		close(ok)
	}()

	select {
	case <-ok:
		return nil
	case <-ctx.Done():
		a.grpcTsrServer.Stop()
		return ctx.Err()
	}

}

func (a *App) initHttpServer(ctx context.Context) error {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := tsr.RegisterUserServiceHandlerFromEndpoint(ctx, mux, a.ServiceProvider.Config.GrpcUserServer.Address, opts)
	if err != nil {
		return err
	}
	err = tsr.RegisterTSRServiceHandlerFromEndpoint(ctx, mux, a.ServiceProvider.Config.GrpcTsrServer.Address, opts)
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
