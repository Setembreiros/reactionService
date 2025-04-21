package integration_test_arrange

import (
	"context"
	"reactionservice/cmd/provider"
	database "reactionservice/internal/db"
	model "reactionservice/internal/model/domain"
	integration_test_assert "reactionservice/test/integration_test_common/assert"
	"testing"
)

func CreateTestDatabase(ctx context.Context) *database.Database {
	provider := provider.NewProvider("test", "postgres://postgres:artis@localhost:5432/artis?search_path=public&sslmode=disable")
	sqlDb, err := provider.ProvideDb()
	if err != nil {
		panic(err)
	}
	return database.NewDatabase(sqlDb)
}

func AddLikePost(t *testing.T, like *model.LikePost) {
	provider := provider.NewProvider("test", "postgres://postgres:artis@localhost:5432/artis?search_path=public&sslmode=disable")
	sqlDb, err := provider.ProvideDb()
	if err != nil {
		panic(err)
	}
	database := database.NewDatabase(sqlDb)

	err = database.Client.CreateLikePost(like)
	if err != nil {
		panic(err)
	}

	integration_test_assert.AssertLikePostExists(t, database, like)
}
