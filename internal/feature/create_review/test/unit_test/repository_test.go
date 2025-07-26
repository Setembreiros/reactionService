package unit_test_create_review

import (
	"errors"
	"testing"

	database "reactionservice/internal/db"
	mock_database "reactionservice/internal/db/test/mock"
	"reactionservice/internal/feature/create_review"
	model "reactionservice/internal/model/domain"

	"github.com/stretchr/testify/assert"
)

var dbClient *mock_database.MockDatabaseClient
var createReviewRepository *create_review.CreateReviewRepository

func repositorySetUp(t *testing.T) {
	setUp(t)
	dbClient = mock_database.NewMockDatabaseClient(ctrl)
	createReviewRepository = create_review.NewCreateReviewRepository(database.NewDatabase(dbClient))
}

func TestCreateReviewInRepository_WhenItReturnsSuccess(t *testing.T) {
	repositorySetUp(t)
	review := &model.Review{
		Username: "usernameA",
		PostId:   "post1",
		Content:  "a minha review",
		Rating:   2,
	}
	expectedReviewId := uint64(5)
	dbClient.EXPECT().CreateReview(review).Return(expectedReviewId, nil)

	reviewId, err := createReviewRepository.CreateReview(review)

	assert.Nil(t, err)
	assert.Equal(t, expectedReviewId, reviewId)
}

func TestErrorOnCreateReviewInRepository_WhenCreateReviewFails(t *testing.T) {
	repositorySetUp(t)
	review := &model.Review{
		Username: "usernameA",
		PostId:   "post1",
		Content:  "a minha review",
		Rating:   2,
	}
	expectedReviewId := uint64(0)
	dbClient.EXPECT().CreateReview(review).Return(expectedReviewId, errors.New("some error"))

	reviewId, err := createReviewRepository.CreateReview(review)

	assert.NotNil(t, err)
	assert.Equal(t, expectedReviewId, reviewId)
}
