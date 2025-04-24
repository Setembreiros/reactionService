package main

import (
	"context"
	"os"
	"os/signal"
	"reactionservice/infrastructure/database/atlas"
	"sync"
	"syscall"

	"reactionservice/cmd/provider"
	"reactionservice/infrastructure/kafka"
	"reactionservice/internal/api"
	"reactionservice/internal/bus"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type App struct {
	Ctx              context.Context
	Cancel           context.CancelFunc
	Env              string
	ConnStr          string
	configuringTasks sync.WaitGroup
	runningTasks     sync.WaitGroup
}

func (app *App) startup() {
	app.configuringLog()

	log.Info().Msgf("Starting ReactionService service in [%s] enviroment...\n", app.Env)

	provider := provider.NewProvider(app.Env, app.ConnStr)

	migrator, err := provider.ProvideAtlasCLient()
	if err != nil {
		os.Exit(1)
	}
	database, err := provider.ProvideDb()
	if err != nil {
		os.Exit(1)
	}
	defer database.Client.Close()
	eventBus, err := provider.ProvideEventBus()
	if err != nil {
		os.Exit(1)
	}
	subscriptions := provider.ProvideSubscriptions(database)
	apiEnpoint := provider.ProvideApiEndpoint(database, eventBus)
	kafkaConsumer, err := provider.ProvideKafkaConsumer(eventBus)
	if err != nil {
		os.Exit(1)
	}

	app.runConfigurationTasks(migrator, subscriptions, eventBus)
	app.runServerTasks(kafkaConsumer, apiEnpoint)
}

func (app *App) configuringLog() {
	if app.Env == "development" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	log.Logger = log.With().Caller().Logger()
}

func (app *App) runConfigurationTasks(atlasCLient *atlas.AtlasClient, subscriptions *[]bus.EventSubscription, eventBus *bus.EventBus) {
	app.configuringTasks.Add(2)
	go app.applyMigrations(atlasCLient)
	go app.subcribeEvents(subscriptions, eventBus) // Always subscribe event before init Kafka
	app.configuringTasks.Wait()
}

func (app *App) runServerTasks(kafkaConsumer *kafka.KafkaConsumer, apiEnpoint *api.Api) {
	app.runningTasks.Add(2)
	go app.initKafkaConsumption(kafkaConsumer)
	go app.runApiEndpoint(apiEnpoint)

	blockForever()

	app.shutdown()
}

func (app *App) applyMigrations(atlasCLient *atlas.AtlasClient) {
	defer app.configuringTasks.Done()

	err := atlasCLient.ApplyMigrations(app.Ctx)
	if err != nil {
		log.Fatal().Stack().Err(err).Msgf("Failed to apply migrations")
	}
}

func (app *App) subcribeEvents(subscriptions *[]bus.EventSubscription, eventBus *bus.EventBus) {
	defer app.configuringTasks.Done()

	log.Info().Msg("Subscribing events...")

	for _, subscription := range *subscriptions {
		eventBus.Subscribe(&subscription, app.Ctx)
		log.Info().Msgf("%s subscribed", subscription.EventType)
	}

	log.Info().Msg("All events subscribed")
}

func (app *App) initKafkaConsumption(kafkaConsumer *kafka.KafkaConsumer) {
	defer app.runningTasks.Done()

	err := kafkaConsumer.InitConsumption(app.Ctx)
	if err != nil {
		log.Panic().Stack().Err(err).Msg("Kafka Consumption failed")
	}
	log.Info().Msg("Kafka Consumer Group stopped")
}

func (app *App) runApiEndpoint(apiEnpoint *api.Api) {
	defer app.runningTasks.Done()

	err := apiEnpoint.Run(app.Ctx)
	if err != nil {
		log.Panic().Err(err).Msg("Closing ReactionService Api failed")
	}
	log.Info().Msg("ReactionService Api stopped")
}

func blockForever() {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	<-signalCh
}

func (app *App) shutdown() {
	app.Cancel()
	log.Info().Msg("Shutting down ReactionService Service...")
	app.runningTasks.Wait()
	log.Info().Msg("ReactionService Service stopped")
}
