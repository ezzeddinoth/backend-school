package router

import (
	"school/internal/api/handler"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockVersionHandler struct {
	handler.Version
	mock.Mock
}

func (m *mockVersionHandler) GetVersion(c *gin.Context) {
	m.Called(c)
}

func getVersionExpectedRoutes(versionHandler handler.Version) (exp []expectedRoutes) {
	exp = []expectedRoutes{
		{
			Method:      "GET",
			Path:        "/api/version",
			FuncName:    "GetVersion",
			HandlerFunc: versionHandler.GetVersion,
		},
	}
	return
}

func setupTestVersionRouter() ([]expectedRoutes, *mockVersionHandler) {
	mockVersionHandler := new(mockVersionHandler)
	expectedRs := getVersionExpectedRoutes(mockVersionHandler)
	for _, e := range expectedRs {
		mockVersionHandler.On(e.FuncName, mock.Anything).Return(nil).Once()
	}

	return expectedRs, mockVersionHandler
}

func TestVersionExpectedRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)
	expected, mock := setupTestVersionRouter()
	r := &router{
		engine:         gin.Default(),
		address:        ":8081",
		releaseMode:    false,
		versionHandler: mock,
	}

	apiGroup := r.engine.Group("/api")
	r.addVersionRoutes(apiGroup)
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
}
