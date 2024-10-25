package main

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"

	_ "github.com/lib/pq"
	"github.com/rodolfoHOk/fullcycle.imersao19/go_transcoder/internal/converter"
	"github.com/rodolfoHOk/fullcycle.imersao19/go_transcoder/pkg/log"
	"github.com/rodolfoHOk/fullcycle.imersao19/go_transcoder/pkg/rabbitmq"
	"github.com/streadway/amqp"
)

func connectPostgres() (*sql.DB, error) {
	user := getEnvOrDefault("POSTGRES_USER", "user")
	password := getEnvOrDefault("POSTGRES_PASSWORD", "password")
	dbname := getEnvOrDefault("POSTGRES_DB", "converter")
	host := getEnvOrDefault("POSTGRES_HOST", "postgres")
	sslmode := getEnvOrDefault("POSTGRES_SSLMODE", "disable")

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=%s", user, password, dbname, host, sslmode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		slog.Error("Error connecting to database", slog.String("connStr", connStr))
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		slog.Error("Error pinging database", slog.String("connStr", connStr))
		return nil, err
	}

	slog.Info("Connected to Postgres successfully")
	return db, nil
}

func getEnvOrDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func main() {
	isDebug := getEnvOrDefault("DEBUG", "false") == "true"
	logger := log.NewLogger(isDebug)
	slog.SetDefault(logger)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	db, err := connectPostgres()
	if err != nil {
		return
	}
	defer db.Close()

	rabbitmqUrl := getEnvOrDefault("RABBITMQ_URL", "amqp://guest:guest@rabbitmq:5672/")
	rabbitmqClient, err := rabbitmq.NewRabbitClient(ctx, rabbitmqUrl)
	if err != nil {
		slog.Error("Failed to connect to RabbitMQ", slog.String("error", err.Error()))
		return
	}
	defer rabbitmqClient.Close()

	conversionExchange := getEnvOrDefault("CONVERSION_EXCHANGE", "conversion_exchange")
	conversionKey := getEnvOrDefault("CONVERSION_KEY", "conversion")
	conversionQueue := getEnvOrDefault("CONVERSION_QUEUE", "video_conversion_queue")
	confirmationKey := getEnvOrDefault("CONFIRMATION_KEY", "finish_conversion")
	confirmationQueue := getEnvOrDefault("CONFIRMATION_QUEUE", "finish_conversion_queue")
	msgs, err := rabbitmqClient.ConsumeMessages(conversionExchange, conversionKey, conversionQueue)
	if err != nil {
		slog.Error("Failed to consume messages", slog.String("error", err.Error()))
		return
	}

	vc := converter.NewVideoConverter(db, rabbitmqClient)
	var wg sync.WaitGroup
	go func() {
		for d := range msgs {
			wg.Add(1)
			go func(delivery amqp.Delivery) {
				defer wg.Done()
				vc.HandleMessage(ctx, delivery, conversionExchange, confirmationKey, confirmationQueue)
			}(d)
		}
	}()

	slog.Info("Waiting for messages from RabbitMQ")
	<-signalChan
	slog.Info("Shutdown signal received, finalizing processing...")

	cancel()

	wg.Wait()

	slog.Info("Processing completed, exiting...")
}
