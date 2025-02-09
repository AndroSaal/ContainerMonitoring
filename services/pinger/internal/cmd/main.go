package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/AndroSaal/ContainerMonitoring/services/pinger/internal/server"
	"github.com/AndroSaal/ContainerMonitoring/services/pinger/internal/service"
	api "github.com/AndroSaal/ContainerMonitoring/services/pinger/internal/transport"
	"github.com/AndroSaal/ContainerMonitoring/services/pinger/pkg/config"
	mylog "github.com/AndroSaal/ContainerMonitoring/services/pinger/pkg/log"
)

func main() {
	// загрузка конфига и переменных окружения, паника в случае ошибки
	cfg := config.MustLoadConfig()

	// настройка и получение инстанса логгера паника в случае ошибки
	logger := mylog.MustNewLogger(cfg.Env)

	// инициализация всех слоев
	docker := service.NewDockerCllient()
	service := service.NewService(docker, logger)
	handler := api.NewHandler(service, logger)

	// инициализация сервера
	server := initServer(cfg.SrvConf, handler, logger)

	// запуск сервера
	if err := server.Run(); err != nil {
		logger.Error("cannot run server: " + err.Error())
	}

	// запуск сервера
	runServer(server)

	// graceful shutdown
	stopServer(server, cfg.SrvConf)

}

func initServer(cfg config.ServerConfig, handler *api.Handler, log *slog.Logger) *server.Server {
	server, err := server.NewServer(cfg, handler.InitRoutes(), log)
	if err != nil {
		log.Error("cannot init server: " + err.Error())
		panic("cannot init server: " + err.Error())
	}
	return server
}

func runServer(server *server.Server) {
	go func() {
		if err := server.Run(); !errors.Is(err, http.ErrServerClosed) {
			server.Logger.Error("cannot run server: " + err.Error())
			panic("cannot run server: " + err.Error())
		}
	}()
}

func stopServer(server *server.Server, cfg config.ServerConfig) {
	// обработка остановки по сигналу
	ctxSig, stop := signal.NotifyContext(
		context.Background(), os.Interrupt, syscall.SIGQUIT, syscall.SIGTERM,
	)
	defer stop()

	// обработка остановки по таймауту
	ctxTim, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
	defer cancel()

	// graceful shutdown
	for {
		select {
		case <-ctxTim.Done():
			server.Logger.Info("Server Stopped by timout")
			server.Stop(ctxTim)
			return
		case <-ctxSig.Done():
			server.Logger.Info("Server stopped by system signall")
			server.Stop(ctxSig)
			return
		}
	}

}
