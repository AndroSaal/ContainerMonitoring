package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/AndroSaal/ContainerMonitoring/services/backend/internal/entities"
	"github.com/AndroSaal/ContainerMonitoring/services/backend/internal/repository"
	"github.com/AndroSaal/ContainerMonitoring/services/backend/pkg/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	ipTable          string = "container_ip"
	idField          string = "id"
	ipField          string = "ip"
	lastSuccessField string = "last_success"
	pingTimeField    string = "ping_time"
	pingTable        string = "container_ping"
	statusField      string = "status"
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

func (p *PostgreDB) AddPingInfo(ctx context.Context, pingfInfo entities.PingInfo) error {

	//проверяем есть ли такой ip уже в базе
	querySelect := fmt.Sprintf(
		"SELECT %s FROM %s WHERE %s = $1",
		idField, ipTable, ipField,
	)
	rowSelect := p.db.QueryRow(querySelect, pingfInfo.IPAdress)
	p.log.Info("Executing query: " + querySelect)

	var ipFromDatabase string
	if err := rowSelect.Scan(&ipFromDatabase); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// если нет, то добавляем
			queryInsert := fmt.Sprintf(
				"INSERT INTO %s (%s, %s) VALUES ($1, $2)",
				ipTable, ipField, lastSuccessField,
			)
			_, err := p.db.Exec(queryInsert, pingfInfo.IPAdress, pingfInfo.LastSuccess)
			if err != nil {
				p.log.Error("Error inserting data to database(1): " + err.Error())
				return err
			}
			ipFromDatabase = pingfInfo.IPAdress
		} else {
			p.log.Error("unexpected error: " + err.Error())
			return err
		}
	}

	queryAddNewPingInfo := fmt.Sprintf(
		"INSERT INTO %s (%s, %s, %s) VALUES ($1, $2, $3)",
		pingTable, ipField, pingTimeField, statusField,
	)

	_, err := p.db.Exec(queryAddNewPingInfo, pingfInfo.IPAdress, pingfInfo.PingTime, pingfInfo.Status)
	if err != nil {
		p.log.Error("Error inserting data to database(2): " + err.Error())
		return err
	}

	return nil
}

func (p *PostgreDB) GetPingInfo(ctx context.Context, ipAdress string) (*[]entities.PingInfo, error) {
	fi := "repository/postgres GetPingInfo"
	// Получаем информацию о пингах из базы
	querySelect := fmt.Sprintf(
		"SELECT %s, %s FROM %s WHERE %s = $1 LIMIT 2",
		pingTimeField, statusField, pingTable, ipField,
	)

	rowsSelect, err := p.db.Query(querySelect, ipAdress)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			p.log.Error("No rows in database: " + err.Error())
			return nil, repository.ErrNotFound
		}
		p.log.Error("unexpected error: " + err.Error())
		return nil, err
	}

	// Заполняем полученную информацию в слайс
	info := make([]entities.PingInfo, 0)
	for rowsSelect.Next() {
		var pingInfo entities.PingInfo
		if err := rowsSelect.Scan(&pingInfo.PingTime, &pingInfo.Status); err != nil {
			p.log.Error("unexpected error: " + err.Error())
			return nil, err
		}
		p.log.Info(fmt.Sprintf("Got {ping} INFO %v", pingInfo))
		pingInfo.IPAdress = ipAdress
		info = append(info, pingInfo)
	}
	p.log.Info(fmt.Sprintf(fi+"Got %d ping infos", len(info)))
	p.log.Info(fmt.Sprintf(fi+"There are %v ping infos", info))

	defer p.closeSomething(rowsSelect.Close(), "can't close rows")

	return &info, nil
}

func (p *PostgreDB) GetAllContainersPingInfo(ctx context.Context) (*[]entities.PingInfo, error) {
	//Получае все ip адреса из базы
	querySelect := fmt.Sprintf(
		"SELECT %s FROM %s",
		ipField, ipTable,
	)

	p.log.Info("Executing query: " + querySelect)
	rowsSelect, err := p.db.Query(querySelect)
	if err != nil {
		p.log.Error("unexpected error: " + err.Error())
		return nil, err
	}

	sliceInfos := make([]entities.PingInfo, 0)
	for rowsSelect.Next() {
		var ip string
		if err := rowsSelect.Scan(&ip); err != nil {
			p.log.Error("unexpected error: " + err.Error())
			return nil, err
		}
		p.log.Info(fmt.Sprintf("Got ip %s", ip))

		info, err := p.GetPingInfo(ctx, ip)
		if err != nil {
			p.log.Error("unexpected error: " + err.Error())
			return nil, err

		}
		p.log.Info(fmt.Sprintf("Got info after ping ip: %v ", *info))

		sliceInfos = append(sliceInfos, *info...)
	}
	p.log.Info(fmt.Sprintf("Got %d ping infos", len(sliceInfos)))
	p.log.Info(fmt.Sprintf("There are all containers %v ping infos", sliceInfos))
	defer p.closeSomething(rowsSelect.Close(), "can't close rows")
	return &sliceInfos, nil
}

func (p *PostgreDB) CloseConnection() {
	if err := p.db.Close(); err != nil {
		p.log.Error("Error closing database: " + err.Error())
	} else {
		p.log.Info("Database closed")
	}
}

func (p *PostgreDB) closeSomething(err error, msg string) {
	if err != nil {
		p.log.Error("Unexpected err:" + msg + ": " + err.Error())
	}
}
