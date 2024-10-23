package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	_ "github.com/lib/pq"
	"github.com/rodolfoHOk/fullcycle.imersao19/go_transcoder/internal/converter"
	"github.com/rodolfoHOk/fullcycle.imersao19/go_transcoder/internal/rabbitmq"
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
	db, err := connectPostgres()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rabbitmqUrl := getEnvOrDefault("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/")
	rabbitmqClient, err := rabbitmq.NewRabbitClient(rabbitmqUrl)
	if err != nil {
		panic(err)
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
	}

	vc := converter.NewVideoConverter(db, rabbitmqClient)

	for d := range msgs {
		go func(delivery amqp.Delivery) {
			vc.Handle(delivery, conversionExchange, confirmationKey, confirmationQueue)
		}(d)
	}

}
