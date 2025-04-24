package unlike_post

import (
	"reactionservice/internal/api"
	model "reactionservice/internal/model/domain"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

//go:generate mockgen -source=controller.go -destination=test/mock/controller.go

type DeleteLikePostController struct {
	service Service
}

type Service interface {
	DeleteLikePost(like *model.LikePost) error
}

func NewDeleteLikePostController(service Service) *DeleteLikePostController {
	return &DeleteLikePostController{
		service: service,
	}
}

func (controller *DeleteLikePostController) Routes(routerGroup *gin.RouterGroup) {
	routerGroup.DELETE("/likePost/:postId/:username", controller.DeleteLikePost)
}

func (controller *DeleteLikePostController) DeleteLikePost(c *gin.Context) {
	log.Info().Msg("Handling Request POST DeleteLikePost")

	like := &model.LikePost{}

	like.PostId = c.Param("postId")
	if like.PostId == "" {
		api.SendBadRequest(c, "Missing postId parameter")
		return
	}

	like.Username = c.Param("username")
	if like.Username == "" {
		api.SendBadRequest(c, "Missing username parameter")
		return
	}

	err := controller.service.DeleteLikePost(like)
	if err != nil {
		api.SendInternalServerError(c, err.Error())
		return
	}

	api.SendOK(c)
}
