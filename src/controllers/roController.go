package controllers

import (
	"fmt"
	"my_app/src/domain"
	"my_app/src/infrastructure"

	"github.com/gin-gonic/gin"
)

const Ro = "/ros"

type RoController struct {
	service *domain.RoService
}

func (c *RoController) CreateRo(ctx *gin.Context) {
	// TODO: implement me!
}

func (c *RoController) GetAllRos(ctx *gin.Context) {
	// TODO: implement me!
}

func (c *RoController) GetRoById(ctx *gin.Context) {
	// TODO: implement me!
}

func (c *RoController) UpdateRo(ctx *gin.Context) {
	// TODO: implement me!
}

func (c *RoController) DeleteRo(ctx *gin.Context) {
	// TODO: implement me!
}

func NewRoAdapter() *RoController {
	return &RoController{}
}

func (c *RoController) SetPort(port infrastructure.Repository) {
	// c.service = domain.NewRoService(port)
}

func (c *RoController) Setup(r *gin.Engine, port infrastructure.Repository) error {
	group := r.Group(Ro)

	porter, ok := port.(domain.RoCrud)
	if !ok {
		return fmt.Errorf("zly typ %s", Ro)
	}

	c.service = domain.NewRoService(porter)

	group.GET("", c.GetAllRos)
	group.GET("/:id", c.GetRoById)
	group.POST("", c.CreateRo)
	group.PATCH("/:id", c.UpdateRo)
	group.DELETE("/:id", c.DeleteRo)

	return nil
}
