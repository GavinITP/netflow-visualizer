package main

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"netflow-visualizer/models"
	"netflow-visualizer/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, falling back to defaults")
	}

	dbPath := os.Getenv("DB_FILE_PATH")
	if dbPath == "" {
		dbPath = "./database/file_paths.db"
	}

	logsPath := os.Getenv("LOGS_FILE_PATH")
	if logsPath == "" {
		logsPath = "./logs/gorm.log"
	}

	if err := os.MkdirAll("./database", os.ModePerm); err != nil {
		log.Fatal("Failed to create database directory:", err)
	}

	logFile, err := os.OpenFile(logsPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	defer logFile.Close()

	newLogger := logger.New(
		log.New(logFile, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	if err := db.AutoMigrate(&models.FilePath{}); err != nil {
		log.Fatal("Failed to migrate database schema:", err)
	}

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET,POST,OPTIONS",
		AllowCredentials: true,
	}))

	routes.SetupRoutes(app, db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Fatal(app.Listen(":" + port))
}
