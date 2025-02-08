package repository

import (
	"context"
	"log/slog"

	"github.com/AndroSaal/ContainerMonitoring/services/backend/internal/entities"
)

// интерфейс для слоя репозитория
type RepositoryHandler interface {
	AddPingInfo(ctx context.Context, pingfInfo entities.PingInfo) (int, error)
	GetPingInfo(ctx context.Context, containerId int) (entities.PingInfo, error)
	GetAllContainersPingInfo(ctx context.Context) ([]entities.PingInfo, error)
}

// интерфейс для реляционной БД
type RelDB interface {
	AddPingInfo(ctx context.Context, pingfInfo entities.PingInfo) (int, error)
	GetPingInfo(ctx context.Context, containerId int) (entities.PingInfo, error)
	GetAllContainersPingInfo(ctx context.Context) ([]entities.PingInfo, error)
	CloseConnection()
}

// имплементация интерфейса слоя репозиторев
type Repository struct {
	relDB RelDB
	log   *slog.Logger
}

func NewRepository(relDB RelDB, log *slog.Logger) *Repository {
	return &Repository{
		relDB: relDB,
		log:   log,
	}
}

func (r *Repository) AddPingInfo(ctx context.Context, pingfInfo entities.PingInfo) (int, error) {
	r.log.Info("AddPingInfo", "containerId", pingfInfo.ContainerId)
	return r.relDB.AddPingInfo(ctx, pingfInfo)

}

func (r *Repository) GetPingInfo(ctx context.Context, containerId int) (entities.PingInfo, error) {
	r.log.Info("GetPingInfo", "containerId", containerId)
	return r.relDB.GetPingInfo(ctx, containerId)

}

func (r *Repository) GetAllContainersPingInfo(ctx context.Context) ([]entities.PingInfo, error) {
	r.log.Info("GetAllContainersPingInfo")
	return r.relDB.GetAllContainersPingInfo(ctx)
}
