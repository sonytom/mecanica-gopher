package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sonytom/mecanica-gopher/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {
	r.POST("/User", controller.CreateUser)
	r.GET("/User/:userId", controller.FindUserById)
	r.PUT("User/:userId", controller.UpdateUserById)
	r.DELETE("User/:UserId", controller.DeleteUserById)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
}
