package exception

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"xyz-multifinance/model/web"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type DuplicateError struct {
	Contains string
	Message  string
}

func ErrorDuplicateMessage(err error, duplicateErrors []DuplicateError) string {
	for _, duplicateError := range duplicateErrors {
		if strings.Contains(err.Error(), duplicateError.Contains) {
			return duplicateError.Message
		}
	}

	return "record already exists"
}

func ErrorHandler(c *gin.Context, err interface{}, duplicateErrors []DuplicateError) (fatal bool) {
	if validationError(c, err) {
		return false
	}

	if sendToResponseError(c, err) {
		return false
	}

	if permissionDeniedError(c, err) {
		return false
	}

	if foreignKeyError(c, err) {
		return true
	}

	if recordNotFoundError(c, err) {
		return false
	}

	if unauthorizedError(c, err) {
		return false
	}

	if duplicateError(c, err, duplicateErrors) {
		return true
	}

	internalServerError(c, err)
	return true
}

func validationError(c *gin.Context, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		webResponse := web.WebResponse{
			Success: false,
			Message: "Bad Request",
			Data:    exception.Error(),
		}

		c.JSON(http.StatusBadRequest, webResponse)
		return true
	} else {
		return false
	}
}

func recordNotFoundError(c *gin.Context, err interface{}) bool {
	exception, ok := err.(error)
	if ok && exception.Error() == "record not found" {
		webResponse := web.WebResponse{
			Success: true,
			Message: "Record not found",
		}

		c.JSON(http.StatusOK, webResponse)
		return true
	}
	return false
}

func sendToResponseError(c *gin.Context, err interface{}) bool {
	exception, ok := err.(*ErrorSendToResponse)
	if ok {
		webResponse := web.WebResponse{
			Success: false,
			Message: exception.Error(),
		}

		c.JSON(http.StatusBadRequest, webResponse)
		return true
	}
	return false
}

func unauthorizedError(c *gin.Context, err interface{}) bool {
	exception, ok := err.(error)
	if ok && (errors.Is(exception, ErrUnauthorized) || errors.Is(exception, ErrRefreshTokenExpired)) {
		webResponse := web.WebResponse{
			Success: false,
			Message: exception.Error(),
		}

		c.JSON(http.StatusUnauthorized, webResponse)
		return true
	}
	return false
}

func permissionDeniedError(c *gin.Context, err interface{}) bool {
	exception, ok := err.(error)
	if ok && errors.Is(exception, ErrPermissionDenied) {
		webResponse := web.WebResponse{
			Success: false,
			Message: exception.Error(),
		}

		c.JSON(http.StatusForbidden, webResponse)
		return true
	}
	return false
}

func duplicateError(c *gin.Context, err interface{}, duplicateErrors []DuplicateError) bool {
	exception, ok := err.(error)
	if ok {
		if strings.Contains(exception.Error(), "Duplicate entry") || strings.Contains(exception.Error(), "already exists") {
			webResponse := web.WebResponse{
				Success: false,
				Message: ErrorDuplicateMessage(exception, duplicateErrors),
			}

			c.JSON(http.StatusBadRequest, webResponse)
			return true
		}
	}
	return false
}

func foreignKeyError(c *gin.Context, err interface{}) bool {
	exception, ok := err.(error)
	if ok {
		if strings.Contains(exception.Error(), "Error 1452: Cannot add or update a child row") {
			webResponse := web.WebResponse{
				Success: false,
				Message: "A foreign key constraint fails",
			}

			c.JSON(http.StatusBadRequest, webResponse)
			return true
		}
	}
	return false
}

func internalServerError(c *gin.Context, err interface{}) {
	webResponse := web.WebResponse{
		Success: false,
		Message: "Internal Server Error",
	}
	fmt.Println(err)
	c.JSON(http.StatusInternalServerError, webResponse)
}
