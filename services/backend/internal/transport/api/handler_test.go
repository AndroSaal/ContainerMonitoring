package api

// import (
// 	"context"
// 	"log/slog"
// 	"net/http"
// 	"net/http/httptest"
// 	"os"
// 	"testing"

// 	"github.com/AndroSaal/ContainerMonitoring/services/backend/internal/entities"
// 	"github.com/gin-gonic/gin"
// 	"github.com/go-playground/assert"
// )

// type serviceMock struct{}

// func (m *serviceMock) AddPingInfo(ctx context.Context, pingfInfo entities.PingInfo) error {
// 	return nil
// }

// func (m *serviceMock) GetPingInfo(ctx context.Context, containerIP string) (entities.PingInfo, error) {
// 	return entities.PingInfo{}, nil
// }

// func (m *serviceMock) GetAllContainersPingInfo(ctx context.Context) ([]entities.PingInfo, error) {
// 	return []entities.PingInfo{}, nil
// }

// func TestHandler_getAllContainersInfo_Correct(t *testing.T) {
// 	handler := NewHandler(
// 		new(serviceMock),
// 		slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})),
// 	)

// 	// формируем тестовый запрос
// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)

// 	handler.getAllContainersInfo(c)

// 	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
// 	defer func() {
// 		if err := c.Request.Body.Close(); err != nil {
// 			t.Fatal()
// 		}
// 	}()
// }
