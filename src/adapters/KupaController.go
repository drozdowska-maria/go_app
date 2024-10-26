package adapters

import (
	"fmt"
	"my_app/src/domain"
	"my_app/src/infrastructure"

	"github.com/gin-gonic/gin"
)

const Kupa = "//kupas"

type KupaController struct {
	service *domain.KupaService
}

func (c *KupaController) CreateKupa(ctx *gin.Context) {
	// TODO: implement me!
}

func (c *KupaController) GetAllKupas(ctx *gin.Context) {
	// TODO: implement me!
}

func (c *KupaController) GetKupaById(ctx *gin.Context) {
	// TODO: implement me!
}

func (c *KupaController) UpdateKupa(ctx *gin.Context) {
	// TODO: implement me!
}

func (c *KupaController) DeleteKupa(ctx *gin.Context) {
	// TODO: implement me!
}

func NewKupaAdapter() *KupaController {
	return &KupaController{}
}

func (c *KupaController) SetPort(port infrastructure.Port) {
	// c.service = domain.NewKupaService(port)
}

func (c *KupaController) Setup(r *gin.Engine, port infrastructure.Port) error {
	group := r.Group(Kupa)

	porter, ok := port.(domain.KupaPorter)
	if !ok {
		return fmt.Errorf("zly typ %s", Kupa)
	}

	c.service = domain.NewKupaService(porter)

	group.GET("", c.GetAllKupas)
	group.GET("/:id", c.GetKupaById)
	group.POST("", c.CreateKupa)
	group.PATCH("/:id", c.UpdateKupa)
	group.DELETE("/:id", c.DeleteKupa)

	return nil
}
