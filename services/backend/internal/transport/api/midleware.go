package api

import (
	"fmt"
	"log/slog"

	"github.com/AndroSaal/ContainerMonitoring/services/backend/internal/entities"
	"github.com/gin-gonic/gin"
)

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	//возвращение ошибки внутри логгера (чтобы мы увидели)
	slog.Error(fmt.Sprintf("error at newErrorResponse (%d, %s)", statusCode, message))
	//возварщение ошибки в качестве ответа (чтобы увидел клиент)
	c.AbortWithStatusJSON(statusCode, entities.ErrorResponse{
		Reason: message,
	})
}

func transformInfos(i *[]entities.PingInfo) *[]entities.PingInfoResponse {
	r := make([]entities.PingInfoResponse, len(*i))
	for k, v := range *i {
		r[k].LastSuccess = fmt.Sprintf("%d-%d-%d", v.LastSuccess.Year(), v.LastSuccess.Month(), v.LastSuccess.Day())
		r[k].PingTime = fmt.Sprintf("%d:%d:%d", v.PingTime.Hour(), v.PingTime.Minute(), v.PingTime.Second())
		r[k].IPAdress = v.IPAdress
		r[k].Status = v.Status
	}

	return &r
}
