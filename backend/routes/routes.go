package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"netflow-visualizer/controllers"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	api := app.Group("/api")

	api.Get("/db-stats", controllers.GetDBStats(db))
	api.Get("/db-logs", controllers.GetRecentLogs())

	api.Get("/anomaly-logs", controllers.GetAnomalyLogs(db))
	api.Get("/overview", controllers.GetOverviewStats(db))
	api.Get("/overview-flow-count", controllers.GetFlowHistory(db))
}
