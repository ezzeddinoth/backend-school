package router

import (
	"school/internal/api/handler"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockSubscribeHandler struct {
	handler.SubscribeHandler
	mock.Mock
}

func (m *mockSubscribeHandler) BulkCreate(c *gin.Context) {
	m.Called(c)
}

func getSubscribeExpectedRoutes(subscribeHandler handler.SubscribeHandler) (exp []expectedRoutes) {
	exp = []expectedRoutes{
		{
			Method:      "POST",
			Path:        "/api/subscribe",
			FuncName:    "BulkCreate",
			HandlerFunc: subscribeHandler.BulkCreate,
		},
	}
	return
}

func setupTestSubscribeRouter() ([]expectedRoutes, *mockSubscribeHandler) {

	mockSubscribeHandler := new(mockSubscribeHandler)
	expectedRs := getSubscribeExpectedRoutes(mockSubscribeHandler)
	for _, e := range expectedRs {
		mockSubscribeHandler.On(e.FuncName, mock.Anything).Return(nil).Once()
	}

	return expectedRs, mockSubscribeHandler
}

func TestSubscribeExpectedRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)
	expected, mock := setupTestSubscribeRouter()
	r := &router{
		engine:      gin.Default(),
		address:     ":8083",
		releaseMode: false,

		subscribeHandler: mock,
	}

	apiGroup := r.engine.Group("/api")
	r.addSubscribeRoutes(apiGroup)

	routes := r.engine.Routes()

	assert.Equal(t, len(expected), len(routes))
	var c *gin.Context
	for _, e := range expected {
		for _, r := range routes {
			if e.Path == r.Path && e.Method == r.Method {
				e.HandlerFunc(c)
				break
			}
		}
	}
	mock.AssertExpectations(t)

	endpointPermission := handler.GetEndpointPermissions()
	for _, r := range routes {
		routerInEndpointPermission := false
		for _, ep := range endpointPermission {
			if r.Path == ep.URL && r.Method == ep.Method {
				routerInEndpointPermission = true
				break
			}
		}
		assert.Equal(t, true, routerInEndpointPermission)
	}

}
