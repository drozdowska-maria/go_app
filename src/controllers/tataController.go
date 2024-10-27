package controllers

import (
	"fmt"
	"my_app/src/domain"
	"my_app/src/infrastructure"

	"github.com/gin-gonic/gin"
)

const Tata = "/tatas"

type TataController struct {
	service *domain.TataService
}

func (c *TataController) CreateTata(ctx *gin.Context) {
	// TODO: implement me!
}

func (c *TataController) GetAllTatas(ctx *gin.Context) {
	// TODO: implement me!
}

func (c *TataController) GetTataById(ctx *gin.Context) {
	// TODO: implement me!
}

func (c *TataController) UpdateTata(ctx *gin.Context) {
	// TODO: implement me!
}

func (c *TataController) DeleteTata(ctx *gin.Context) {
	// TODO: implement me!
}

func NewTataAdapter() *TataController {
	return &TataController{}
}

func (c *TataController) SetPort(port infrastructure.Repository) {
	// c.service = domain.NewTataService(port)
}

func (c *TataController) Setup(r *gin.Engine, port infrastructure.Repository) error {
	group := r.Group(Tata)

	porter, ok := port.(domain.TataCrud)
	if !ok {
		return fmt.Errorf("zly typ %s", Tata)
	}

	c.service = domain.NewTataService(porter)

	group.GET("", c.GetAllTatas)
	group.GET("/:id", c.GetTataById)
	group.POST("", c.CreateTata)
	group.PATCH("/:id", c.UpdateTata)
	group.DELETE("/:id", c.DeleteTata)

	return nil
}
