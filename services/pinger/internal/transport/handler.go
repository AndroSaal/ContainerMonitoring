package api

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/AndroSaal/ContainerMonitoring/services/pinger/internal/entities"
	"github.com/AndroSaal/ContainerMonitoring/services/pinger/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.ServiceHandler
	logger  *slog.Logger
}

func NewHandler(service service.ServiceHandler, logger *slog.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}

func (h *Handler) pingIP(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	var ips entities.Ips

	// получаем все айпи, которые нужно пингануть, 400 в случае ошибки
	if err := c.ShouldBindJSON(&ips); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// запускаем пинг, возвращаем 500 в случае ошибки
	if err := h.service.StartPing(ctx, time.Second*5, ips.Ips); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "OK")

}

func (h *Handler) pingAll(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	ips, err := h.service.GetAllContainersIP(ctx)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err1 := h.service.StartPing(ctx, time.Second*5, ips)
	if err1 != nil {
		newErrorResponse(c, http.StatusBadRequest, err1.Error())
		return
	}

	c.JSON(http.StatusOK, "OK")

}
