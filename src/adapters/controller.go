package adapters

import (
	"my_app/src/domain"

	"github.com/gin-gonic/gin"
)

func NewUserController(port domain.UserPorter) *Controller {
	return &Controller{
		service: domain.NewUserService(port),
	}
}

type Controller struct {
	service *domain.UserService
}

func (uc *Controller) CreateUser(ctx *gin.Context) {
	// gin adaptation
	// wywolanie serwisu
}

func (uc *Controller) GetAllUsers(ctx *gin.Context) {
	// gin adaptation
	// wywolanie serwisu
}

func (uc *Controller) GetUserById(ctx *gin.Context) {
	// gin adaptation
	// wywolanie serwisu
}

func (uc *Controller) UpdateUser(ctx *gin.Context) {
	// gin adaptation
	// wywolanie serwisu
}

func (uc *Controller) DeleteUser(ctx *gin.Context) {
	// gin adaptation
	// wywolanie serwisu
}
