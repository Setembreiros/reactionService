package superlike_post

import (
	"reactionservice/internal/api"
	model "reactionservice/internal/model/domain"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

//go:generate mockgen -source=controller.go -destination=test/mock/controller.go

type SuperlikePostController struct {
	service Service
}

type Service interface {
	CreateSuperlikePost(superlike *model.SuperlikePost) error
}

func NewSuperlikePostController(service Service) *SuperlikePostController {
	return &SuperlikePostController{
		service: service,
	}
}

func (controller *SuperlikePostController) Routes(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/superlikePost", controller.CreateSuperlikePost)
}

func (controller *SuperlikePostController) CreateSuperlikePost(c *gin.Context) {
	log.Info().Msg("Handling Request POST SuperlikePost")
	var superlike model.SuperlikePost

	if err := c.BindJSON(&superlike); err != nil {
		log.Error().Stack().Err(err).Msg("Invalid Data")
		api.SendBadRequest(c, "Invalid Json Request")
		return
	}

	err := controller.service.CreateSuperlikePost(&superlike)
	if err != nil {
		api.SendInternalServerError(c, err.Error())
		return
	}

	api.SendOK(c)
}
