package like_post

import (
	"reactionservice/internal/api"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

//go:generate mockgen -source=controller.go -destination=test/mock/controller.go

type LikePostController struct {
	service Service
}

type Service interface {
	LikePost(postId uint64) error
}

func NewLikePostController(service Service) *LikePostController {
	return &LikePostController{
		service: service,
	}
}

func (controller *LikePostController) Routes(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/like/:postId", controller.LikePost)
}

func (controller *LikePostController) LikePost(c *gin.Context) {
	log.Info().Msg("Handling Request POST LikePost")

	postId := c.Param("postId")
	if postId == "" {
		api.SendBadRequest(c, "Missing postId parameter")
		return
	}

	id, err := strconv.ParseUint(postId, 10, 64)
	if err != nil {
		log.Error().Stack().Err(err).Msgf("PostId %s couldn't be parsed", postId)
		api.SendBadRequest(c, "PostId couldn't be parsed. PostId should be a positive number")
		return
	}

	err = controller.service.LikePost(id)
	if err != nil {
		api.SendInternalServerError(c, err.Error())
		return
	}

	api.SendOK(c)
}
