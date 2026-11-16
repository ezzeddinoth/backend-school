package router

import (
	"net/http"
	"net/http/httptest"
	"slices"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

type expectedRoutes struct {
	Method      string
	Path        string
	FuncName    string
	HandlerFunc gin.HandlerFunc
}

func TestServerRoutes(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"value":"jwks"}`))
	}))
	defer server.Close()
	t.Setenv("KEYCLOAK_JWKS_URL", server.URL)
	t.Setenv("SCHOOL_ALLOWED_ORIGINS_LIST", "*")

	gin.SetMode(gin.TestMode)
	expected, r := setupTestRouter()

	r.SetupRouter()
	routes := r.engine.Routes()

	expectedString := []string{}
	for _, e := range expected {
		expectedString = append(expectedString, e.Path+e.Method)
	}

	routesString := []string{}
	for _, r := range routes {
		routesString = append(routesString, r.Path+r.Method)
	}

	for _, e := range expectedString {
		if !slices.Contains(routesString, e) {
			t.Fatalf("routes do not contains %s", e)
		}
	}

	for _, r := range routesString {
		if !slices.Contains(expectedString, r) {
			t.Fatalf("route %s is not expected", r)
		}
	}
	assert.Equal(t, len(expected), len(routes))
}

func setupTestRouter() ([]expectedRoutes, *router) {
	expected := []expectedRoutes{}
	expectedVersion, versionHandler := setupTestVersionRouter()
	expectedStudent, studentHandler := setupTestStudentRouter()
	expectedMenu, menuHandler := setupTestMenuRouter()
	expectedSubscribe, subscribeHandler := setupTestSubscribeRouter()

	expected = append(expected, expectedVersion...)
	expected = append(expected, expectedStudent...)
	expected = append(expected, expectedMenu...)
	expected = append(expected, expectedSubscribe...)

	return expected, &router{
		engine:      gin.Default(),
		address:     ":8083",
		releaseMode: true, //remove swagger

		versionHandler:   versionHandler,
		studentHandler:   studentHandler,
		menuHandler:      menuHandler,
		subscribeHandler: subscribeHandler,
	}
}
