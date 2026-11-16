package handler

import (
	"net/http"
	"school/internal/service"

	"github.com/gin-gonic/gin"
)

type Version interface {
	GetVersion(c *gin.Context)
}
type version struct {
	Version string `json:"version"`
}

func NewVersion() Version {
	return &version{}
}

// @Summary			get version info
// @Description		get version info
// @Tags  			public endpoints
// @Accept			json
// @Produce			json
// @Success			200	{string}	string	"OK"
// @Failure			400	{string}	string	"Bad Request"
// @Failure			401	{string}	string	"Not Authorized"
// @Router			/version [get]
func (h *version) GetVersion(c *gin.Context) {
	c.JSON(http.StatusOK, version{Version: service.GetVersion()})
}
