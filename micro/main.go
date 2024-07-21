package main

import (
	"context"
	"database/sql"
	"log/slog"
	"micro/config"
	"micro/logger"
	"micro/message"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.MustLoad()

	log := logger.New(cfg.LogLevel)
	log.Debug("Logger is initializated")

	dbConn := connDB(cfg.DSN, log)

	msgRepo := message.NewRepo(dbConn, log)
	msgServ := message.NewService(msgRepo, log)
	msgHandler := message.NewHandler(msgServ, log)

	router := gin.Default()
	router.POST("/send", msgHandler.Receiver)
	router.GET("/stat", nil)

}

//  connStr := "postgresql://<username>:<password>@<database_ip>/todos?sslmode=disable"

func connDB(dsn string, logger *slog.Logger) *sql.DB {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Error("failed connect to db: " + err.Error())
		panic("failed to connect to db: " + err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		logger.Error("failed ping database: " + err.Error())
		panic("failed ping database: " + err.Error())
	}

	logger.Debug("db success connected")

	return db
}
