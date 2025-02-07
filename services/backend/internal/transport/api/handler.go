package api

import (
	"log/slog"
	"net/http"

	"github.com/AndroSaal/ContainerMonitoring/services/backend/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.Service
	logger  slog.Logger
}

func NewHandler(service service.Service, logger slog.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}

func (h *Handler) getAllContainersInfo(c *gin.Context) {
	newErrorResponse(c, http.StatusAccepted, "hehe")
}

func (h *Handler) getContainerInfo(c *gin.Context) {
	newErrorResponse(c, http.StatusAccepted, "hehe")
}

func (h *Handler) addCreateContainerInfo(c *gin.Context) {
	newErrorResponse(c, http.StatusAccepted, "hehe")
}
