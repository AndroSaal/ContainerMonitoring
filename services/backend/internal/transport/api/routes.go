package api

import "github.com/gin-gonic/gin"

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	pingInfo := router.Group("/ping")
	{
		pingInfo.GET("", h.getAllContainersInfo)
		pingInfo.GET("/:id", h.getContainerInfo)
		pingInfo.POST("", h.addContainerInfo)
	}
	h.logger.Info("Routes inited")
	return router
}
