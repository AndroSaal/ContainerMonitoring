package service

import (
	"context"
	"log/slog"

	"github.com/AndroSaal/ContainerMonitoring/services/backend/internal/entities"
	"github.com/AndroSaal/ContainerMonitoring/services/backend/internal/repository"
)

type ServiceHandler interface {
	AddPingInfo(ctx context.Context, pingfInfo entities.PingInfo) error
	GetPingInfo(ctx context.Context, containerIP string) (entities.PingInfo, error)
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

func (s *Service) AddPingInfo(ctx context.Context, pingfInfo entities.PingInfo) error {
	return nil
}

func (s *Service) GetPingInfo(ctx context.Context, containerIP string) (entities.PingInfo, error) {
	return entities.PingInfo{}, nil
}
func (s *Service) GetAllContainersPingInfo(ctx context.Context) ([]entities.PingInfo, error) {
	return nil, nil
}

