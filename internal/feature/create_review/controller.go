package create_review

import (
	"reactionservice/internal/api"
	model "reactionservice/internal/model/domain"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

//go:generate mockgen -source=controller.go -destination=test/mock/controller.go

type CreateReviewController struct {
	service Service
}

type Service interface {
	CreateReview(review *model.Review) error
}

func NewCreateReviewController(service Service) *CreateReviewController {
	return &CreateReviewController{
		service: service,
	}
}

func (controller *CreateReviewController) Routes(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/review", controller.CreateReview)
}

func (controller *CreateReviewController) CreateReview(c *gin.Context) {
	log.Info().Msg("Handling Request POST CreateReview")
	var review model.Review

	if err := c.BindJSON(&review); err != nil {
		log.Error().Stack().Err(err).Msg("Invalid Data")
		api.SendBadRequest(c, "Invalid Json Request")
		return
	}

	err := controller.service.CreateReview(&review)
	if err != nil {
		api.SendInternalServerError(c, err.Error())
		return
	}

	api.SendOK(c)
}
