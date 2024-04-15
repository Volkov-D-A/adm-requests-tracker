package app

import (
	"context"
	"fmt"

	"github.com/volkov-d-a/adm-requests-tracker/internal/adapters/api"
	storage "github.com/volkov-d-a/adm-requests-tracker/internal/adapters/db/postgres"
	"github.com/volkov-d-a/adm-requests-tracker/internal/service"
	pg "github.com/volkov-d-a/adm-requests-tracker/pkg/PG"
	"github.com/volkov-d-a/adm-requests-tracker/pkg/closer"
	"github.com/volkov-d-a/adm-requests-tracker/pkg/config"
	"github.com/volkov-d-a/adm-requests-tracker/pkg/logger"
)

type serviceProvider struct {
	userApi     *api.UserApi
	userService api.UserService
	userStorage service.UserStorage
	tsrApi      *api.TSRApi
	tsrService  api.TSRService
	tsrStorage  service.TSRStorage
	Config      *config.Config
	Logger      *logger.Logger
	Closer      *closer.Closer
	db          *pg.PG
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) SetConfig() error {
	conf, err := config.GetConfig()
	if err != nil {
		return err
	}
	s.Config = conf
	return nil
}

func (s *serviceProvider) SetLogger() error {
	log, err := logger.GetLogger("dev")
	if err != nil {
		return err
	}
	s.Logger = log
	return nil
}

func (s *serviceProvider) SetDB() error {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", s.Config.PG.User, s.Config.PG.Password, s.Config.PG.Host, s.Config.PG.Port, s.Config.PG.Database)
	db, err := pg.NewDB(dsn, s.Config.PG.MP)
	if err != nil {
		return err
	}
	s.db = db
	s.Logger.Info("Database connection established")
	s.Closer.Add("DB", s.CloseDB)
	return nil
}

func (s *serviceProvider) CloseDB(ctx context.Context) error {
	ok := make(chan struct{})
	go func() {
		s.db.Close()
		close(ok)
	}()

	select {
	case <-ok:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (s *serviceProvider) TsrStorage() service.TSRStorage {
	if s.tsrStorage == nil {
		s.tsrStorage = storage.NewTsrStorage(s.db)
	}
	return s.tsrStorage
}

func (s *serviceProvider) TsrService() api.TSRService {
	if s.tsrService == nil {
		s.tsrService = service.NewTSRService(
			s.TsrStorage(),
		)
	}
	return s.tsrService
}

func (s *serviceProvider) TSRApi() *api.TSRApi {
	cfg := &api.TSRConfig{
		Key: s.Config.Key,
	}
	if s.tsrApi == nil {
		s.tsrApi = api.NewTSRApi(
			s.TsrService(),
			cfg,
		)
	}
	return s.tsrApi
}

func (s *serviceProvider) UserStorage() service.UserStorage {
	if s.userStorage == nil {
		s.userStorage = storage.NewUserStorage(s.db)
	}
	return s.userStorage
}

func (s *serviceProvider) UserService() api.UserService {
	if s.userService == nil {
		s.userService = service.NewUserService(
			s.UserStorage(),
		)
	}
	return s.userService
}

func (s *serviceProvider) UserApi() *api.UserApi {
	cfg := &api.UserConfig{
		Key: s.Config.Key,
	}
	if s.userApi == nil {
		s.userApi = api.NewUserApi(
			s.UserService(),
			cfg,
		)
	}
	return s.userApi
}
