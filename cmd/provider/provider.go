package provider

import (
	"reactionservice/infrastructure/atlas"
	"reactionservice/infrastructure/kafka"
	"reactionservice/internal/api"
	"reactionservice/internal/bus"
	database "reactionservice/internal/db"
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

func (p *Provider) ProvideAtlasCLient() (*atlas.AtlasClient, error) {
	return atlas.NewAtlasClient(p.connStr)
}

func (p *Provider) ProvideDb() (*database.Database, error) {
	return database.NewDatabase(p.connStr)
}

func (p *Provider) ProvideEventBus() (*bus.EventBus, error) {
	kafkaProducer, err := kafka.NewKafkaProducer(p.kafkaBrokers())
	if err != nil {
		return nil, err
	}

	return bus.NewEventBus(kafkaProducer), nil
}

func (p *Provider) ProvideApiEndpoint() *api.Api {
	return api.NewApiEndpoint(p.env, p.ProvideApiControllers())
}

func (p *Provider) ProvideApiControllers() []api.Controller {
	return []api.Controller{}
}

func (p *Provider) ProvideSubscriptions(database *database.Database) *[]bus.EventSubscription {
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
			"172.31.36.175:9092",
			"172.31.45.255:9092",
		}
	}
}
