package app

import (
	api "github.com/volkov-d-a/adm-requests-tracker/internal/api/tsr"
	"github.com/volkov-d-a/adm-requests-tracker/internal/repository"

	"github.com/volkov-d-a/adm-requests-tracker/internal/service"

	"github.com/volkov-d-a/adm-requests-tracker/pkg/closer"
	"github.com/volkov-d-a/adm-requests-tracker/pkg/config"
	"github.com/volkov-d-a/adm-requests-tracker/pkg/logger"
)

type serviceProvider struct {
	tsrImplementation *api.Implementation
	userService       service.UserService
	userRepository    repository.UserRepository
	Config            *config.Config
	Logger            *logger.Logger
	Closer            *closer.Closer
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

func (s *serviceProvider) UserRepository() repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = repository.NewUserRepository()
	}
	return s.userRepository
}

func (s *serviceProvider) UserService() service.UserService {
	conf := &service.Config{
		Key: s.Config.Key,
	}
	if s.userService == nil {
		s.userService = service.NewUserService(
			s.UserRepository(),
			conf,
		)
	}
	return s.userService
}

func (s *serviceProvider) TsrImplement() *api.Implementation {
	if s.tsrImplementation == nil {
		s.tsrImplementation = api.NewImplementation(
			s.UserService(),
		)
	}
	return s.tsrImplementation
}
