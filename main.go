package main

import (
	"my_app/src/adapters"
	"my_app/src/ports/v1"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Adapter interface {
	NewAdapter() Adapter
}

func setupRoutes(r *gin.Engine, adapter map[string]Adapter) {
	var groups map[string]*gin.RouterGroup

	for groupName, _ := range adapter {
		groups[groupName] = r.Group(groupName)
	}

	userController := adapters.NewUserController(
		ports.NewUserRepository(),
	)

	for groupName, group := range groups {
		switch groupName {
		case adapters.User:
			group.GET("", userController.GetAllUsers)
			group.GET("/:id", userController.GetUserById)
			group.POST("", userController.CreateUser)
			group.PATCH("/:id", userController.UpdateUser)
			group.DELETE("/:id", userController.DeleteUser)
		}
	}

}

func main() {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong, ong kong"})
	})

	r.GET("/hello", func(c *gin.Context) {
		name := c.Query("name")
		if name == "" {
			name = "World"
		}
		c.JSON(http.StatusOK, gin.H{"message": "Hello " + name})
	})

	r.Run()
}
