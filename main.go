package main

import (
	"fmt"
	a "my_app/src/adapters"
	"my_app/src/infrastructure"
	p "my_app/src/ports/v1"
	"os"

	"net/http"

	"github.com/gin-gonic/gin"
)

func initAdapters() map[string]infrastructure.Adapter {
	var adapters = map[string]infrastructure.Adapter{}

	adapters[a.User] = a.NewUserAdapter()

	return adapters
}

func initPorts() map[string]infrastructure.Port {
	var ports = map[string]infrastructure.Port{}

	ports[a.User] = p.NewUserRepository()

	return ports
}

func setupGinRouting(adapters map[string]infrastructure.Adapter, ports map[string]infrastructure.Port) error {
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

	for groupName, ad := range adapters {
		switch groupName {
		case a.User:
			userController, ok := ad.(*a.UserController)
			if !ok {
				return fmt.Errorf("zly typ %s", a.User)
			}
			userController.Setup(r, ports[a.User])
		}
	}

	r.Run()
	return nil
}

func main() {

	adapters := initAdapters()
	ports := initPorts()

	err := setupGinRouting(adapters, ports)

	if err != nil {
		os.Exit(1)
	}

}
