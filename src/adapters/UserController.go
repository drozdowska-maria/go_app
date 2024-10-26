package adapters

import (
	"fmt"
	"my_app/src/domain"
	"my_app/src/infrastructure"

	"github.com/gin-gonic/gin"
)

const User = "/users"

type UserController struct {
	service *domain.UserService
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	// TODO: implement me!
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
	// TODO: implement me!
}

func (c *UserController) GetUserById(ctx *gin.Context) {
	// TODO: implement me!
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	// TODO: implement me!
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	// TODO: implement me!
}

func NewUserAdapter() *UserController {
	return &UserController{}
}

func (c *UserController) SetPort(port infrastructure.Port) {
	// c.service = domain.NewUserService(port)
}

func (c *UserController) Setup(r *gin.Engine, port infrastructure.Port) error {
	group := r.Group(User)

	porter, ok := port.(domain.UserPorter)
	if !ok {
		return fmt.Errorf("zly typ %s", User)
	}

	c.service = domain.NewUserService(porter)

	group.GET("", c.GetAllUsers)
	group.GET("/:id", c.GetUserById)
	group.POST("", c.CreateUser)
	group.PATCH("/:id", c.UpdateUser)
	group.DELETE("/:id", c.DeleteUser)

	return nil
}
