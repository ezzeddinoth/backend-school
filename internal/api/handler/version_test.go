package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestVersionHandler(t *testing.T) {
	v := &version{}
	gin.SetMode(gin.TestMode)
	gin.DefaultWriter = io.Discard
	router := gin.Default()
	router.GET("/api/version", v.GetVersion)

	r := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/version", nil)
	router.ServeHTTP(r, req)

	var versionResponse version
	json.Unmarshal(r.Body.Bytes(), &versionResponse)

	expected := "development_commit_1970-01-01T00:00:00Z"
	assert.Equal(t, expected, versionResponse.Version)
	assert.Equal(t, http.StatusOK, r.Code)
}
