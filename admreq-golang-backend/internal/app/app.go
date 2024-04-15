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

func NewApp() (*App, error) {
	a := &App{}
	err := a.initDeps()
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (a *App) Run(se chan error) {
	a.ServiceProvider.Logger.Info("Starting GRPC Servers at:", "addr", a.ServiceProvider.Config.GrpcUserServer.Address, "addr", a.ServiceProvider.Config.GrpcTsrServer.Address)
	a.runGRPCServer(se)

	a.ServiceProvider.Logger.Info("Starting GRPC GW Server at:", "addr", a.ServiceProvider.Config.GrpcGw.Address)
	a.runHttpServer(se)
}

func (a *App) GracefulStop() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	err := a.ServiceProvider.Closer.Close(ctx)
	if err != nil {
		return fmt.Errorf("eror while stopping app: %v", err)
	}
	return nil
}

func (a *App) initDeps() error {
	inits := []func() error{
		a.initServiceProvider,
		a.initGRPCServer,
		a.initHttpServer,
	}
	for _, f := range inits {
		err := f()
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *App) initServiceProvider() error {
	a.ServiceProvider = newServiceProvider()
	err := a.ServiceProvider.SetConfig()
	if err != nil {
		return err
	}
	err = a.ServiceProvider.SetLogger()
	if err != nil {
		return err
	}
	a.ServiceProvider.Closer = closer.NewAppCloser()
	err = a.ServiceProvider.SetDB()
	if err != nil {
		return err
	}
	a.ServiceProvider.Logger.Info("Service provider initialized")
	return nil
}

func (a *App) initGRPCServer() error {
	a.grpcUserServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	reflection.Register(a.grpcUserServer)
	tsr.RegisterUserServiceServer(a.grpcUserServer, a.ServiceProvider.UserApi())

	a.grpcTsrServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	reflection.Register(a.grpcTsrServer)
	tsr.RegisterTSRServiceServer(a.grpcTsrServer, a.ServiceProvider.TSRApi())

	a.ServiceProvider.Logger.Info("GRPC Server initialized")
	return nil
}

func (a *App) runGRPCServer(se chan error) {
	a.ServiceProvider.Closer.Add("UserGRPC", a.stopUserGRPCServer)
	a.ServiceProvider.Closer.Add("TSRGrpc", a.stopTSRGRPCServer)

	go func() {
		list, err := net.Listen("tcp", a.ServiceProvider.Config.GrpcUserServer.Address)
		if err != nil {
			a.ServiceProvider.Closer.Remove("UserGRPC")
			se <- err
		}
		err = a.grpcUserServer.Serve(list)
		if err != nil {
			a.ServiceProvider.Closer.Remove("UserGRPC")
			se <- err
		}
	}()

	go func() {
		list, err := net.Listen("tcp", a.ServiceProvider.Config.GrpcTsrServer.Address)
		if err != nil {
			a.ServiceProvider.Closer.Remove("TSRGrpc")
			se <- err
		}
		err = a.grpcTsrServer.Serve(list)
		if err != nil {
			a.ServiceProvider.Closer.Remove("TSRGrpc")
			se <- err
		}
	}()

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

func (a *App) initHttpServer() error {
	ctx := context.Background()
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
		Handler:        cors(mux),
	}
	return nil
}

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
		h.ServeHTTP(w, r)
	})
}

func (a *App) runHttpServer(se chan error) {
	a.ServiceProvider.Closer.Add("HTTP", a.httpServer.Shutdown)
	go func() {
		err := a.httpServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			a.ServiceProvider.Closer.Remove("HTTP")
			se <- err
		}
	}()

}
