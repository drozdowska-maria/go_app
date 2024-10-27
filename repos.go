package main

import (
	"my_app/src/controllers"
	"my_app/src/infrastructure"
	repositories "my_app/src/repositories/v1"
)

var repos = map[string]infrastructure.Repository{
	controllers.User: repositories.NewUserRepository(),
	controllers.Room: repositories.NewRoomRepository(),
	controllers.Roo: repositories.NewRooRepository(),
	controllers.Ro: repositories.NewRoRepository(),
	controllers.Hoo: repositories.NewHooRepository(),
	controllers.Tata: repositories.NewTataRepository(),
}
