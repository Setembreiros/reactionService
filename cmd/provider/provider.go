package provider

import (
	"reactionservice/infrastructure/kafka"
	"reactionservice/internal/api"
	"reactionservice/internal/bus"
)

type Provider struct {
	env string
}

func NewProvider(env string) *Provider {
	return &Provider{
		env: env,
	}
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
