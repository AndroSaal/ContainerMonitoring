package service

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"net"
	"net/http"
	"time"

	"github.com/AndroSaal/ContainerMonitoring/services/pinger/internal/entities"
)

type ServiceHandler interface {
	StartPing(ctx context.Context, interval time.Duration, ipAddresses []string) error
	SendPingInfo(ctx context.Context, inf entities.PingInfo) error
	GetAllContainersIP(ctx context.Context) ([]string, error)
}

type Service struct {
	DC  DockerClientHandler
	log *slog.Logger
}

type DockerClientHandler interface {
	GetAllContainersIP(ctx context.Context) ([]string, error)
}

func (s *Service) StartPing(ctx context.Context, interval time.Duration, ipAddresses []string) error {
	s.log.Info("Start pinging")

	//заводим тикер для интервального пинга
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		//проходим по всем ip адресам
		for _, ip := range ipAddresses {

			var (
				lastSuccessDate time.Time
				status          string = "failure"
			)

			// проверяем соединение - пингуем
			success := ping(ip)
			pingTime := time.Now()

			if success {
				lastSuccessDate = pingTime
				status = "success"
			}

			pingInfo := entities.PingInfo{
				IPAddress:   ip,
				PingTime:    pingTime,
				LastSuccess: lastSuccessDate,
				Status:      status,
			}

			if err := s.SendPingInfo(ctx, pingInfo); err != nil {
				s.log.Error("Send ping info error", "error", err)
				return err
			}
		}
	}
	return nil
}

func (s *Service) SendPingInfo(ctx context.Context, pingInfo entities.PingInfo) error {
	s.log.Info("Send ping info")
	jsonData, err := json.Marshal(pingInfo)
	if err != nil {
		slog.Error("Error while serializing data", "error", err)
		return err
	}

	resp, err := http.Post("http://backend-service:8080/ping", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		slog.Error("Error while sending data", "error", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		slog.Error("Unexpected status Code from backend-service, want 200, have :", "status", resp.StatusCode)
	}

	return nil
}

func (s *Service) GetAllContainersIP(ctx context.Context) ([]string, error) {
	return s.DC.GetAllContainersIP(ctx)
}

func NewService(doc DockerClientHandler, log *slog.Logger) *Service {
	return &Service{
		DC:  doc,
		log: log,
	}
}

func ping(ip string) bool {
	_, err := net.DialTimeout("ip4:icmp", ip, 1*time.Second)
	return err == nil
}
