package postgres

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/AndroSaal/ContainerMonitoring/services/backend/internal/entities"
	"github.com/AndroSaal/ContainerMonitoring/services/backend/pkg/config"
	"github.com/jmoiron/sqlx"
)

type PostgreDB struct {
	db  *sqlx.DB
	log *slog.Logger
}

func NewPostgresDB(log *slog.Logger, cfg config.DBConfig) *PostgreDB {
	dbConn, err := sqlx.Connect("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Dbname, cfg.Sslmode))

	if err != nil {
		log.Error("Error connecting to database: " + err.Error())
		panic("error connecting databse: " + err.Error())
	}
	log.Info("Connected to database PostgreSQL: " + cfg.Dbname)

	return &PostgreDB{
		db:  dbConn,
		log: log,
	}
}

func (p *PostgreDB) AddPingInfo(ctx context.Context, pingfInfo entities.PingInfo) (int, error) {
	return 0, nil
}

func (p *PostgreDB) GetPingInfo(ctx context.Context, containerId int) (entities.PingInfo, error) {
	return entities.PingInfo{}, nil
}

func (p *PostgreDB) GetAllContainersPingInfo(ctx context.Context) ([]entities.PingInfo, error) {
	return []entities.PingInfo{}, nil
}

func (p *PostgreDB) CloseConnection() {
	if err := p.db.Close(); err != nil {
		p.log.Error("Error closing database: " + err.Error())
	} else {
		p.log.Info("Database closed")
	}
}
