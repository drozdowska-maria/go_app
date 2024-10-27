package controllers

import (
	"fmt"
	"my_app/src/domain"
	"my_app/src/infrastructure"

	"github.com/gin-gonic/gin"
)

const Room = "/rooms"

type RoomController struct {
	service *domain.RoomService
}

func (c *RoomController) CreateRoom(ctx *gin.Context) {
	// TODO: implement me!
}

func (c *RoomController) GetAllRooms(ctx *gin.Context) {
	// TODO: implement me!
}

func (c *RoomController) GetRoomById(ctx *gin.Context) {
	// TODO: implement me!
}

func (c *RoomController) UpdateRoom(ctx *gin.Context) {
	// TODO: implement me!
}

func (c *RoomController) DeleteRoom(ctx *gin.Context) {
	// TODO: implement me!
}

func NewRoomAdapter() *RoomController {
	return &RoomController{}
}

func (c *RoomController) SetPort(port infrastructure.Repository) {
	// c.service = domain.NewRoomService(port)
}

func (c *RoomController) Setup(r *gin.Engine, port infrastructure.Repository) error {
	group := r.Group(Room)

	porter, ok := port.(domain.RoomCrud)
	if !ok {
		return fmt.Errorf("zly typ %s", Room)
	}

	c.service = domain.NewRoomService(porter)

	group.GET("", c.GetAllRooms)
	group.GET("/:id", c.GetRoomById)
	group.POST("", c.CreateRoom)
	group.PATCH("/:id", c.UpdateRoom)
	group.DELETE("/:id", c.DeleteRoom)

	return nil
}
