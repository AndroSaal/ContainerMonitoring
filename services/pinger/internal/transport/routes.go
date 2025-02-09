package api

import "github.com/gin-gonic/gin"

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	pingContainer := router.Group("/pingIP")
	{
		//пинг конкретных нескольки контейнеров по их ip
		pingContainer.POST("", h.pingIP)
		//пинг всех существующих контейнеров на машине
		pingContainer.POST("/all", h.pingAll)
	}

	return router
}
