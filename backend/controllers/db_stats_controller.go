package controllers

import (
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"gorm.io/gorm"
)

func GetDBStats(db *gorm.DB) fiber.Handler {
	return websocket.New(func(c *websocket.Conn) {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()
		for {
			<-ticker.C
			var pageSize, pageCount, freelistCount, cacheSize, synchronous, schemaVersion, userVersion, foreignKeysEnabled int
			var journalMode, integrityCheck string

			db.Raw("PRAGMA page_size").Scan(&pageSize)
			db.Raw("PRAGMA page_count").Scan(&pageCount)
			db.Raw("PRAGMA freelist_count").Scan(&freelistCount)
			db.Raw("PRAGMA cache_size").Scan(&cacheSize)
			db.Raw("PRAGMA synchronous").Scan(&synchronous)
			db.Raw("PRAGMA schema_version").Scan(&schemaVersion)
			db.Raw("PRAGMA user_version").Scan(&userVersion)
			db.Raw("PRAGMA foreign_keys").Scan(&foreignKeysEnabled)
			db.Raw("PRAGMA journal_mode").Scan(&journalMode)
			db.Raw("PRAGMA integrity_check").Scan(&integrityCheck)

			dbSize := pageSize * pageCount
			dbFilePath := os.Getenv("DB_FILE_PATH")

			absPath, err := filepath.Abs(dbFilePath)
			if err != nil {
				absPath = "Error: " + err.Error()
			}

			sqlDB, err := db.DB()
			if err != nil {
				c.WriteJSON(fiber.Map{"error": "Failed to get underlying sql.DB"})
				return
			}
			dbStats := sqlDB.Stats()

			stats := fiber.Map{
				"page_size":        pageSize,
				"page_count":       pageCount,
				"db_size_bytes":    dbSize,
				"freelist_count":   freelistCount,
				"cache_size":       cacheSize,
				"synchronous":      synchronous,
				"schema_version":   schemaVersion,
				"user_version":     userVersion,
				"foreign_keys":     foreignKeysEnabled,
				"journal_mode":     journalMode,
				"integrity_check":  integrityCheck,
				"db_abs_path":      absPath,
				"open_connections": dbStats.OpenConnections,
			}

			if err := c.WriteJSON(stats); err != nil {
				return
			}
		}
	})
}
