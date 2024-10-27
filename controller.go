package main

import (
	"my_app/src/controllers"
	"my_app/src/infrastructure"
)

var ctrls = map[string]infrastructure.Controller{
	controllers.User: controllers.NewUserAdapter(),
	controllers.Room: controllers.NewRoomAdapter(),
	controllers.Roo: controllers.NewRooAdapter(),
	controllers.Ro: controllers.NewRoAdapter(),
	controllers.Hoo: controllers.NewHooAdapter(),
	controllers.Tata: controllers.NewTataAdapter(),
}
