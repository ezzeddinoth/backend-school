package router

import (
	"school/internal/api/handler"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockStudentHandler struct {
	handler.StudentHandler
	mock.Mock
}

func (m *mockStudentHandler) GetAll(c *gin.Context) {
	m.Called(c)
}

func (m *mockStudentHandler) Create(c *gin.Context) {
	m.Called(c)
}

func (m *mockStudentHandler) Update(c *gin.Context) {
	m.Called(c)
}

func (m *mockStudentHandler) Delete(c *gin.Context) {
	m.Called(c)
}

func getStudentExpectedRoutes(studentHandler handler.StudentHandler) (exp []expectedRoutes) {
	exp = []expectedRoutes{
		{
			Method:      "GET",
			Path:        "/api/student",
			FuncName:    "GetAll",
			HandlerFunc: studentHandler.GetAll,
		},
		{
			Method:      "POST",
			Path:        "/api/student",
			FuncName:    "Create",
			HandlerFunc: studentHandler.Create,
		},
		{
			Method:      "PUT",
			Path:        "/api/student",
			FuncName:    "Update",
			HandlerFunc: studentHandler.Update,
		},
		{
			Method:      "DELETE",
			Path:        "/api/student",
			FuncName:    "Delete",
			HandlerFunc: studentHandler.Delete,
		},
	}
	return
}

func setupTestStudentRouter() ([]expectedRoutes, *mockStudentHandler) {

	mockStudentHandler := new(mockStudentHandler)
	expectedRs := getStudentExpectedRoutes(mockStudentHandler)
	for _, e := range expectedRs {
		mockStudentHandler.On(e.FuncName, mock.Anything).Return(nil).Once()
	}

	return expectedRs, mockStudentHandler
}

func TestStudentExpectedRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)
	expected, mock := setupTestStudentRouter()
	r := &router{
		engine:      gin.Default(),
		address:     ":8083",
		releaseMode: false,

		studentHandler: mock,
	}

	apiGroup := r.engine.Group("/api")
	r.addStudentRoutes(apiGroup)

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
