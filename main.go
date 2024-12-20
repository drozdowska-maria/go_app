package main

import (
	"fmt"
	"my_app/src/controllers"
	"my_app/src/infrastructure"
	"os"

	"github.com/gin-gonic/gin"
)

func setupGinRouting(ctrls map[string]infrastructure.Controller, repos map[string]infrastructure.Repository) error {
	r := gin.Default()

	for groupName, ad := range ctrls {
		switch groupName {
		case controllers.User:
			userController, ok := ad.(*controllers.UserController)
			if !ok {
				return fmt.Errorf("zly typ %s", controllers.User)
			}
			userController.Setup(r, repos[controllers.User])

		case controllers.Tata:
			tataController, ok := ad.(*controllers.TataController)
			if !ok {
				return fmt.Errorf("zly typ %s", controllers.Tata)
			}
			tataController.Setup(r, repos[controllers.Tata])

			// NEXT CASES - DO NOT DELETE THIS COMMENT

		}
	}

	r.Run()
	return nil
}

func main() {
	err := setupGinRouting(ctrls, repos)

	if err != nil {
		os.Exit(1)
	}

}
