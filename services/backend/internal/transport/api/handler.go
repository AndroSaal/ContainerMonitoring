package api

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/AndroSaal/ContainerMonitoring/services/backend/internal/entities"
	"github.com/AndroSaal/ContainerMonitoring/services/backend/internal/repository"
	"github.com/AndroSaal/ContainerMonitoring/services/backend/internal/service"
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

func (h *Handler) getAllContainersInfo(c *gin.Context) {
	fi := "api.handler.getAllContainersInfo"

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	pingInfoS, err := h.service.GetAllContainersPingInfo(ctx)

	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			newErrorResponse(c, http.StatusRequestTimeout, ErrServerDown)
		} else if errors.Is(err, repository.ErrNotFound) {
			newErrorResponse(c, http.StatusRequestTimeout, ErrNotFound.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	pingInfosResponse := transformInfos(pingInfoS)

	h.logger.Info("success finish", fi, pingInfosResponse)
	c.AbortWithStatusJSON(http.StatusOK, map[string]interface{}{
		"AllContainersPingInfo": *pingInfosResponse,
	})

}

func (h *Handler) getContainerInfo(c *gin.Context) {
	fi := "api.handler.getContainerInfo"

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	ip := c.Param("ip")
	if ip == "" {
		newErrorResponse(c, http.StatusBadRequest, "ip parameter does not exist in path")
		return
	}

	pingInfo, err := h.service.GetPingInfo(ctx, ip)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	pingInfosResponse := transformInfos(pingInfo)

	h.logger.Info("success finish", fi, pingInfosResponse)
	c.JSON(http.StatusOK, *pingInfosResponse)
}

func (h *Handler) addContainerInfo(c *gin.Context) {
	fi := "api.handler.addContainerInfo"
	var pingInfo entities.PingInfo
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	if err := c.ShouldBindJSON(&pingInfo); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.service.AddPingInfo(ctx, pingInfo)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.logger.Info("success finish", fi, pingInfo)
	c.JSON(http.StatusOK, "successfully added")
}
