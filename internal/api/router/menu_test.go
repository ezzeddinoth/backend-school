package router

import (
	"school/internal/api/handler"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockMenuHandler struct {
	handler.MenuHandler
	mock.Mock
}

func (m *mockMenuHandler) GetAll(c *gin.Context) {
	m.Called(c)
}

func (m *mockMenuHandler) Create(c *gin.Context) {
	m.Called(c)
}

func (m *mockMenuHandler) Update(c *gin.Context) {
	m.Called(c)
}

func (m *mockMenuHandler) Delete(c *gin.Context) {
	m.Called(c)
}

func getMenuExpectedRoutes(menuHandler handler.MenuHandler) (exp []expectedRoutes) {
	exp = []expectedRoutes{
		{
			Method:      "GET",
			Path:        "/api/menu",
			FuncName:    "GetAll",
			HandlerFunc: menuHandler.GetAll,
		},
		{
			Method:      "POST",
			Path:        "/api/menu",
			FuncName:    "Create",
			HandlerFunc: menuHandler.Create,
		},
		{
			Method:      "PUT",
			Path:        "/api/menu",
			FuncName:    "Update",
			HandlerFunc: menuHandler.Update,
		},
		{
			Method:      "DELETE",
			Path:        "/api/menu",
			FuncName:    "Delete",
			HandlerFunc: menuHandler.Delete,
		},
	}
	return
}

func setupTestMenuRouter() ([]expectedRoutes, *mockMenuHandler) {

	mockMenuHandler := new(mockMenuHandler)
	expectedRs := getMenuExpectedRoutes(mockMenuHandler)
	for _, e := range expectedRs {
		mockMenuHandler.On(e.FuncName, mock.Anything).Return(nil).Once()
	}

	return expectedRs, mockMenuHandler
}

func TestMenuExpectedRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)
	expected, mock := setupTestMenuRouter()
	r := &router{
		engine:      gin.Default(),
		address:     ":8083",
		releaseMode: false,

		menuHandler: mock,
	}

	apiGroup := r.engine.Group("/api")
	r.addMenuRoutes(apiGroup)

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
