package handler

import (
	"log"
	"net/http"
	"school/internal/api/dto"
	"school/internal/service"

	"github.com/gin-gonic/gin"
)

type SubscribeHandler interface {
	BulkCreate(c *gin.Context)
}

type subscribeHandler struct {
	subscribeService service.SubscribeService
}

func NewSubscribeHandler() SubscribeHandler {
	return &subscribeHandler{
		subscribeService: service.NewSubscribeService(),
	}
}

// @Summary		create subscribe entries
// @Description	create subscribe entries
// @Tags			subscribe
// @Accept			json
// @Produce			json
// @Param			subscribe	body	[]dto.PostSubscribe				true	"create subscribe entries"
// @Success		200			{object}	dto.Subscribe		     		"OK"
// @Failure		400			{object}	dto.ResponseError			"Bad Request"
// @Failure		401			{object}	dto.ResponseError			"Not Authorized"
// @Failure		403			{object}	dto.ResponseError			"Forbidden"
// @Failure		404			{object}	dto.ResponseError			"Not Found"
// @Failure		500			{object}	dto.ResponseError			"Internal Server Error"
// @Router			/subscribe [post]
// @Security		OAuth2Password
// @Security		Bearer
func (h *subscribeHandler) BulkCreate(c *gin.Context) {
	var subscribes []dto.PostSubscribe
	if err := c.ShouldBindJSON(&subscribes); err != nil {
		c.JSON(http.StatusBadRequest, dto.ResponseError{ErrorMessage: err.Error()})
		return
	}

	resp, err := h.subscribeService.BulkCreate(subscribes)
	if err != nil {
		log.Printf("could not create subscribions: %v", err)
		panic(err)
	}

	c.JSON(http.StatusOK, resp)
}
