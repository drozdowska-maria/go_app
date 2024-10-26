package adapters

import (
	"fmt"
	"my_app/src/domain"
	"my_app/src/infrastructure"

	"github.com/gin-gonic/gin"
)

const Dupa = "/dupas"

type DupaController struct {
	service *domain.DupaService
}

func (c *DupaController) CreateDupa(ctx *gin.Context) {
	// TODO: implement me!
}

func (c *DupaController) GetAllDupas(ctx *gin.Context) {
	// TODO: implement me!
}

func (c *DupaController) GetDupaById(ctx *gin.Context) {
	// TODO: implement me!
}

func (c *DupaController) UpdateDupa(ctx *gin.Context) {
	// TODO: implement me!
}

func (c *DupaController) DeleteDupa(ctx *gin.Context) {
	// TODO: implement me!
}

func NewDupaAdapter() *DupaController {
	return &DupaController{}
}

func (c *DupaController) SetPort(port infrastructure.Port) {
	// c.service = domain.NewDupaService(port)
}

func (c *DupaController) Setup(r *gin.Engine, port infrastructure.Port) error {
	group := r.Group(Dupa)

	porter, ok := port.(domain.DupaPorter)
	if !ok {
		return fmt.Errorf("zly typ %s", Dupa)
	}

	c.service = domain.NewDupaService(porter)

	group.GET("", c.GetAllDupas)
	group.GET("/:id", c.GetDupaById)
	group.POST("", c.CreateDupa)
	group.PATCH("/:id", c.UpdateDupa)
	group.DELETE("/:id", c.DeleteDupa)

	return nil
}
