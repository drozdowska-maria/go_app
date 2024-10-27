package controllers

import (
	"fmt"
	"my_app/src/domain"
	"my_app/src/infrastructure"

	"github.com/gin-gonic/gin"
)

const Hoo = "/hoos"

type HooController struct {
	service *domain.HooService
}

func (c *HooController) CreateHoo(ctx *gin.Context) {
	// TODO: implement me!
}

func (c *HooController) GetAllHoos(ctx *gin.Context) {
	// TODO: implement me!
}

func (c *HooController) GetHooById(ctx *gin.Context) {
	// TODO: implement me!
}

func (c *HooController) UpdateHoo(ctx *gin.Context) {
	// TODO: implement me!
}

func (c *HooController) DeleteHoo(ctx *gin.Context) {
	// TODO: implement me!
}

func NewHooAdapter() *HooController {
	return &HooController{}
}

func (c *HooController) SetPort(port infrastructure.Repository) {
	// c.service = domain.NewHooService(port)
}

func (c *HooController) Setup(r *gin.Engine, port infrastructure.Repository) error {
	group := r.Group(Hoo)

	porter, ok := port.(domain.HooCrud)
	if !ok {
		return fmt.Errorf("zly typ %s", Hoo)
	}

	c.service = domain.NewHooService(porter)

	group.GET("", c.GetAllHoos)
	group.GET("/:id", c.GetHooById)
	group.POST("", c.CreateHoo)
	group.PATCH("/:id", c.UpdateHoo)
	group.DELETE("/:id", c.DeleteHoo)

	return nil
}
