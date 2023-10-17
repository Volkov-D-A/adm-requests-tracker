package app

import (
	"github.com/volkov-d-a/adm-requests-tracker/internal/controller"
	"github.com/volkov-d-a/adm-requests-tracker/internal/repository"
	"github.com/volkov-d-a/adm-requests-tracker/internal/service"
	"github.com/volkov-d-a/adm-requests-tracker/pkg/closer"
	"github.com/volkov-d-a/adm-requests-tracker/pkg/config"
	"github.com/volkov-d-a/adm-requests-tracker/pkg/logger"
)

type serviceProvider struct {
	cn     *controller.TSRController
	ts     *service.TSRService
	tr     *repository.TSRRepository
	Config *config.Config
	Logger *logger.Logger
	Closer *closer.Closer
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

func (s *serviceProvider) TSRRepository() *repository.TSRRepository {
	if s.tr == nil {
		s.tr = repository.New()
	}
	return s.tr
}

func (s *serviceProvider) TSRService() *service.TSRService {
	conf := &service.Config{
		Key: s.Config.Key,
	}
	if s.ts == nil {
		s.ts = service.New(
			s.TSRRepository(),
			conf,
		)
	}
	return s.ts
}

func (s *serviceProvider) TSRController() *controller.TSRController {
	if s.cn == nil {
		s.cn = controller.New(
			s.TSRService(),
		)
	}
	return s.cn
}
