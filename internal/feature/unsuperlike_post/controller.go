package unsuperlike_post

import (
	"reactionservice/internal/api"
	model "reactionservice/internal/model/domain"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

//go:generate mockgen -source=controller.go -destination=test/mock/controller.go

type DeleteSuperlikePostController struct {
	service Service
}

type Service interface {
	DeleteSuperlikePost(superlike *model.SuperlikePost) error
}

func NewDeleteSuperlikePostController(service Service) *DeleteSuperlikePostController {
	return &DeleteSuperlikePostController{
		service: service,
	}
}

func (controller *DeleteSuperlikePostController) Routes(routerGroup *gin.RouterGroup) {
	routerGroup.DELETE("/superlikePost/:postId/:username", controller.DeleteSuperlikePost)
}

func (controller *DeleteSuperlikePostController) DeleteSuperlikePost(c *gin.Context) {
	log.Info().Msg("Handling Request POST DeleteSuperlikePost")

	superlike := &model.SuperlikePost{}

	superlike.PostId = c.Param("postId")
	if superlike.PostId == "" {
		api.SendBadRequest(c, "Missing postId parameter")
		return
	}

	superlike.Username = c.Param("username")
	if superlike.Username == "" {
		api.SendBadRequest(c, "Missing username parameter")
		return
	}

	err := controller.service.DeleteSuperlikePost(superlike)
	if err != nil {
		api.SendInternalServerError(c, err.Error())
		return
	}

	api.SendOK(c)
}
