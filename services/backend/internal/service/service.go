package service

import (
	"context"
	"log/slog"

	"github.com/AndroSaal/ContainerMonitoring/services/backend/internal/entities"
	"github.com/AndroSaal/ContainerMonitoring/services/backend/internal/repository"
)

type ServiceHandler interface {
	AddPingInfo(ctx context.Context, pingfInfo entities.PingInfo) (int, error)
	GetPingInfo(ctx context.Context, containerId int) (entities.PingInfo, error)
	GetAllContainersPingInfo(ctx context.Context) ([]entities.PingInfo, error)
}

type Service struct {
	repo repository.RepositoryHandler
	log  *slog.Logger
}

func NewService(repo repository.RepositoryHandler, log *slog.Logger) *Service {
	return &Service{
		repo: repo,
		log:  log,
	}
}

func (s *Service) AddPingInfo(ctx context.Context, pingfInfo entities.PingInfo) (int, error) {
	return 0, nil
}

func (s *Service) GetPingInfo(ctx context.Context, containerId int) (entities.PingInfo, error) {
	return entities.PingInfo{}, nil
}
func (s *Service) GetAllContainersPingInfo(ctx context.Context) ([]entities.PingInfo, error) {
	return nil, nil
}
