package handler

import (
	"fmt"
	"log"
	"net/http"
	"school/internal/api/dto"
	"school/internal/service"

	"github.com/gin-gonic/gin"
)

type MenuHandler interface {
	GetAll(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type menuHandler struct {
	menuService service.MenuService
}

func NewMenuHandler() MenuHandler {
	return &menuHandler{
		menuService: service.NewMenuService(),
	}
}

// @Summary		get all menus
// @Description	get all menus
// @Tags			menu
// @Produce		json
// @Success		200	{object}	[]dto.Menu 					"OK"
// @Failure		401	{object}	dto.ResponseError				"Not Authorized"
// @Failure		403	{object}	dto.ResponseError				"Forbidden"
// @Failure		500	{object}	dto.ResponseError				"Internal Server Error"
// @Router			/menu [get]
// @Security		OAuth2Password
// @Security		Bearer
func (h *menuHandler) GetAll(c *gin.Context) {
	menus, err := h.menuService.GetAll()
	if err != nil {
		panic(fmt.Sprintf("could not get menus: %v", err))
	}

	c.JSON(http.StatusOK, menus)
}

// @Summary		create menu entries
// @Description	create menu entries
// @Tags			menu
// @Accept			json
// @Produce			json
// @Param			menu	body		dto.PostMenu				true	"create menu entries"
// @Success		200			{object}	dto.Menu		     		"OK"
// @Failure		400			{object}	dto.ResponseError			"Bad Request"
// @Failure		401			{object}	dto.ResponseError			"Not Authorized"
// @Failure		403			{object}	dto.ResponseError			"Forbidden"
// @Failure		404			{object}	dto.ResponseError			"Not Found"
// @Failure		500			{object}	dto.ResponseError			"Internal Server Error"
// @Router			/menu [post]
// @Security		OAuth2Password
// @Security		Bearer
func (h *menuHandler) Create(c *gin.Context) {
	var menu dto.PostMenu
	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, dto.ResponseError{ErrorMessage: err.Error()})
		return
	}

	resp, err := h.menuService.Create(menu)
	if err != nil {
		log.Printf("could not create menu: %v", err)
		panic(err)
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary		update menu entries
// @Description	update menu entries
// @Tags			menu
// @Accept			json
// @Produce			json
// @Param			menu	body		dto.Menu				true	"create menu entries"
// @Success		200			{object}	dto.Menu		     		"OK"
// @Failure		400			{object}	dto.ResponseError			"Bad Request"
// @Failure		401			{object}	dto.ResponseError			"Not Authorized"
// @Failure		403			{object}	dto.ResponseError			"Forbidden"
// @Failure		404			{object}	dto.ResponseError			"Not Found"
// @Failure		500			{object}	dto.ResponseError			"Internal Server Error"
// @Router			/menu [put]
// @Security		OAuth2Password
// @Security		Bearer
func (h *menuHandler) Update(c *gin.Context) {
	var menu dto.Menu
	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, dto.ResponseError{ErrorMessage: err.Error()})
		return
	}

	resp, err := h.menuService.Update(menu)
	if err != nil {
		log.Printf("could not create menu: %v", err)
		panic(err)
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary		update menu entries
// @Description	update menu entries
// @Tags			menu
// @Accept			json
// @Produce			json
// @Param			id		body		int				true		"create menu entries"
// @Success		200			{object}	string		     			"OK"
// @Failure		400			{object}	dto.ResponseError			"Bad Request"
// @Failure		401			{object}	dto.ResponseError			"Not Authorized"
// @Failure		404			{object}	dto.ResponseError			"Not Found"
// @Failure		500			{object}	dto.ResponseError			"Internal Server Error"
// @Router			/menu [delete]
// @Security		OAuth2Password
// @Security		Bearer
func (h *menuHandler) Delete(c *gin.Context) {
	var id int
	if err := c.ShouldBindJSON(&id); err != nil {
		c.JSON(http.StatusBadRequest, dto.ResponseError{ErrorMessage: err.Error()})
		return
	}

	err := h.menuService.Delete(id)
	if err != nil {
		log.Printf("could not create menu: %v", err)
		panic(err)
	}

	c.JSON(http.StatusOK, nil)
}
