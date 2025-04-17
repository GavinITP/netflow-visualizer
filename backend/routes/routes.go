package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"netflow-visualizer/controllers"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	api := app.Group("/api")

	api.Get("/db-stats", controllers.GetDBStats(db))
	api.Get("/db-logs", controllers.GetRecentLogs("./logs/gorm.log", 50))
	api.Get("/netflows", controllers.GetNetflowsFromDBFilePath(db))
	api.Get("/netflow-stats", controllers.GetNetflowStatsFromDB(db))
}
