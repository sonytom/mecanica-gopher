package controller

import (
	"github.com/gin-gonic/gin"
	handleError "github.com/sonytom/mecanica-gopher/src/configuration"
)

func CreateUser(c *gin.Context) {

	err := handleError.NewBadRequestError("method error")

	c.JSON(err.Code, err)
}
