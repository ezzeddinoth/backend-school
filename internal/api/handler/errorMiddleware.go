package handler

import (
	"fmt"
	"net/http"
	"school/internal/api/dto"
	"school/internal/exception"

	"github.com/gin-gonic/gin"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if badRequest, ok := err.(*exception.BadRequest); ok {
					c.JSON(http.StatusBadRequest, dto.ResponseError{ErrorMessage: badRequest.Error()})
				} else if notFound, ok := err.(*exception.NotFound); ok {
					c.JSON(http.StatusNotFound, dto.ResponseError{ErrorMessage: notFound.Error()})
				} else if uniqueErr, ok := err.(*exception.UniqueContraint); ok {
					c.JSON(http.StatusConflict, dto.ResponseError{ErrorMessage: uniqueErr.Error()})
				} else if foreignKey, ok := err.(*exception.ForeignKeyConstraint); ok {
					c.JSON(http.StatusConflict, dto.ResponseError{ErrorMessage: foreignKey.Error()})
				} else if foreignKey, ok := err.(*exception.CollectionElementLimitExceeded); ok {
					c.JSON(http.StatusConflict, dto.ResponseError{ErrorMessage: foreignKey.Error()})
				} else {
					c.JSON(http.StatusInternalServerError, dto.ResponseError{ErrorMessage: fmt.Sprintf("Unknown error: %v", err)})
				}
			}
		}()

		c.Next()
	}
}
