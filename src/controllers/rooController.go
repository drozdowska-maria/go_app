package controllers

import (
	"fmt"
	"my_app/src/domain"
	"my_app/src/infrastructure"

	"github.com/gin-gonic/gin"
)

const Roo = "/roos"

type RooController struct {
	service *domain.RooService
}

func (c *RooController) CreateRoo(ctx *gin.Context) {
	// TODO: implement me!
}

func (c *RooController) GetAllRoos(ctx *gin.Context) {
	// TODO: implement me!
}

func (c *RooController) GetRooById(ctx *gin.Context) {
	// TODO: implement me!
}

func (c *RooController) UpdateRoo(ctx *gin.Context) {
	// TODO: implement me!
}

func (c *RooController) DeleteRoo(ctx *gin.Context) {
	// TODO: implement me!
}

func NewRooAdapter() *RooController {
	return &RooController{}
}

func (c *RooController) SetPort(port infrastructure.Repository) {
	// c.service = domain.NewRooService(port)
}

func (c *RooController) Setup(r *gin.Engine, port infrastructure.Repository) error {
	group := r.Group(Roo)

	porter, ok := port.(domain.RooCrud)
	if !ok {
		return fmt.Errorf("zly typ %s", Roo)
	}

	c.service = domain.NewRooService(porter)

	group.GET("", c.GetAllRoos)
	group.GET("/:id", c.GetRooById)
	group.POST("", c.CreateRoo)
	group.PATCH("/:id", c.UpdateRoo)
	group.DELETE("/:id", c.DeleteRoo)

	return nil
}
