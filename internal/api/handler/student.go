package handler

import (
	"fmt"
	"log"
	"net/http"
	"school/internal/api/dto"
	"school/internal/service"

	"github.com/gin-gonic/gin"
)

type StudentHandler interface {
	GetAll(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type studentHandler struct {
	studentService service.StudentService
}

func NewStudentHandler() StudentHandler {
	return &studentHandler{
		studentService: service.NewStudentService(),
	}
}

// @Summary		get all students
// @Description	get all students
// @Tags			student
// @Produce		json
// @Success		200	{object}	[]dto.Student 					"OK"
// @Failure		401	{object}	dto.ResponseError				"Not Authorized"
// @Failure		403	{object}	dto.ResponseError				"Forbidden"
// @Failure		500	{object}	dto.ResponseError				"Internal Server Error"
// @Router			/student [get]
// @Security		OAuth2Password
// @Security		Bearer
func (h *studentHandler) GetAll(c *gin.Context) {
	students, err := h.studentService.GetAll()
	if err != nil {
		panic(fmt.Sprintf("could not get students: %v", err))
	}

	c.JSON(http.StatusOK, students)
}

// @Summary		create student entries
// @Description	create student entries
// @Tags			student
// @Accept			json
// @Produce			json
// @Param			student	body		dto.PostStudent				true	"create student entries"
// @Success		200			{object}	dto.Student		     		"OK"
// @Failure		400			{object}	dto.ResponseError			"Bad Request"
// @Failure		401			{object}	dto.ResponseError			"Not Authorized"
// @Failure		403			{object}	dto.ResponseError			"Forbidden"
// @Failure		404			{object}	dto.ResponseError			"Not Found"
// @Failure		500			{object}	dto.ResponseError			"Internal Server Error"
// @Router			/student [post]
// @Security		OAuth2Password
// @Security		Bearer
func (h *studentHandler) Create(c *gin.Context) {
	var student dto.PostStudent
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, dto.ResponseError{ErrorMessage: err.Error()})
		return
	}

	resp, err := h.studentService.Create(student)
	if err != nil {
		log.Printf("could not create student: %v", err)
		panic(err)
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary		update student entries
// @Description	update student entries
// @Tags			student
// @Accept			json
// @Produce			json
// @Param			student	body		dto.Student				true	"create student entries"
// @Success		200			{object}	dto.Student		     		"OK"
// @Failure		400			{object}	dto.ResponseError			"Bad Request"
// @Failure		401			{object}	dto.ResponseError			"Not Authorized"
// @Failure		403			{object}	dto.ResponseError			"Forbidden"
// @Failure		404			{object}	dto.ResponseError			"Not Found"
// @Failure		500			{object}	dto.ResponseError			"Internal Server Error"
// @Router			/student [put]
// @Security		OAuth2Password
// @Security		Bearer
func (h *studentHandler) Update(c *gin.Context) {
	var student dto.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, dto.ResponseError{ErrorMessage: err.Error()})
		return
	}

	resp, err := h.studentService.Update(student)
	if err != nil {
		log.Printf("could not create student: %v", err)
		panic(err)
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary		update student entries
// @Description	update student entries
// @Tags			student
// @Accept			json
// @Produce			json
// @Param			id		body		int				true		"create student entries"
// @Success		200			{object}	string		     			"OK"
// @Failure		400			{object}	dto.ResponseError			"Bad Request"
// @Failure		401			{object}	dto.ResponseError			"Not Authorized"
// @Failure		404			{object}	dto.ResponseError			"Not Found"
// @Failure		500			{object}	dto.ResponseError			"Internal Server Error"
// @Router			/student [delete]
// @Security		OAuth2Password
// @Security		Bearer
func (h *studentHandler) Delete(c *gin.Context) {
	var id int
	if err := c.ShouldBindJSON(&id); err != nil {
		c.JSON(http.StatusBadRequest, dto.ResponseError{ErrorMessage: err.Error()})
		fmt.Println(err)
		return
	}

	err := h.studentService.Delete(id)
	if err != nil {
		log.Printf("could not create student: %v", err)
		panic(err)
	}

	c.JSON(http.StatusOK, nil)
}
