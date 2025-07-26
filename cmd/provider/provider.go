package provider

import (
	"reactionservice/infrastructure/database/migrator"
	"reactionservice/infrastructure/database/sql_db"
	"reactionservice/infrastructure/kafka"
	"reactionservice/internal/api"
	"reactionservice/internal/bus"
	database "reactionservice/internal/db"
	"reactionservice/internal/feature/create_review"
	"reactionservice/internal/feature/like_post"
	"reactionservice/internal/feature/superlike_post"
	"reactionservice/internal/feature/unlike_post"
	"reactionservice/internal/feature/unsuperlike_post"
	"reactionservice/internal/service"
)

type Provider struct {
	env     string
	connStr string
}

func NewProvider(env, connStr string) *Provider {
	return &Provider{
		env:     env,
		connStr: connStr,
	}
}

func (p *Provider) ProvideGooseCLient() (*migrator.GooseClient, error) {
	return migrator.NewGooseClient(p.connStr)
}

func (p *Provider) ProvideDb() (*sql_db.SqlDatabase, error) {
	return sql_db.NewDatabase(p.connStr)
}

func (p *Provider) ProvideEventBus() (*bus.EventBus, error) {
	kafkaProducer, err := kafka.NewKafkaProducer(p.kafkaBrokers())
	if err != nil {
		return nil, err
	}

	return bus.NewEventBus(kafkaProducer), nil
}

func (p *Provider) ProvideApiEndpoint(sqlClient *sql_db.SqlDatabase, bus *bus.EventBus) *api.Api {
	return api.NewApiEndpoint(p.env, p.ProvideApiControllers(sqlClient, bus))
}

func (p *Provider) ProvideApiControllers(sqlClient *sql_db.SqlDatabase, bus *bus.EventBus) []api.Controller {
	return []api.Controller{
		like_post.NewLikePostController(like_post.NewLikePostService(like_post.NewCreateLikePostRepository(database.NewDatabase(sqlClient)), bus)),
		unlike_post.NewDeleteLikePostController(unlike_post.NewDeleteLikePostService(unlike_post.NewDeleteLikePostRepository(database.NewDatabase(sqlClient)), bus)),
		superlike_post.NewSuperlikePostController(superlike_post.NewSuperlikePostService(superlike_post.NewCreateSuperlikePostRepository(database.NewDatabase(sqlClient)), bus)),
		unsuperlike_post.NewDeleteSuperlikePostController(unsuperlike_post.NewDeleteSuperlikePostService(unsuperlike_post.NewDeleteSuperlikePostRepository(database.NewDatabase(sqlClient)), bus)),
		create_review.NewCreateReviewController(create_review.NewCreateReviewService(service.GetTimeServiceInstance(), create_review.NewCreateReviewRepository(database.NewDatabase(sqlClient)), bus)),
	}
}

func (p *Provider) ProvideSubscriptions(database *sql_db.SqlDatabase) *[]bus.EventSubscription {
	return &[]bus.EventSubscription{}
}

func (p *Provider) ProvideKafkaConsumer(eventBus *bus.EventBus) (*kafka.KafkaConsumer, error) {
	brokers := p.kafkaBrokers()

	return kafka.NewKafkaConsumer(brokers, eventBus)
}

func (p *Provider) kafkaBrokers() []string {
	if p.env == "development" {
		return []string{
			"localhost:9093",
		}
	} else {
		return []string{
			"172.31.0.242:9092",
			"172.31.7.110:9092",
		}
	}
}
