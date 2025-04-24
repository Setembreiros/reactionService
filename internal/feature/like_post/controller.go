package like_post

import (
	"reactionservice/internal/api"
	model "reactionservice/internal/model/domain"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

//go:generate mockgen -source=controller.go -destination=test/mock/controller.go

type LikePostController struct {
	service Service
}

type Service interface {
	CreateLikePost(like *model.LikePost) error
}

func NewLikePostController(service Service) *LikePostController {
	return &LikePostController{
		service: service,
	}
}

func (controller *LikePostController) Routes(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/likePost", controller.CreateLikePost)
}

func (controller *LikePostController) CreateLikePost(c *gin.Context) {
	log.Info().Msg("Handling Request POST LikePost")
	var like model.LikePost

	if err := c.BindJSON(&like); err != nil {
		log.Error().Stack().Err(err).Msg("Invalid Data")
		api.SendBadRequest(c, "Invalid Json Request")
		return
	}

	err := controller.service.CreateLikePost(&like)
	if err != nil {
		api.SendInternalServerError(c, err.Error())
		return
	}

	api.SendOK(c)
}
