package controllers

import (
	"os"
	"path/filepath"
	"reflect"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"gorm.io/gorm"
)

// GetDBStatsWS streams database stats over WebSocket, only sending when the underlying file changes
func GetDBStats(db *gorm.DB) fiber.Handler {
	return websocket.New(func(c *websocket.Conn) {
		var lastStats map[string]interface{}
		var lastModTime time.Time

		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		send := func() error {
			// Cheap check: file modification time
			dbPath := os.Getenv("DB_FILE_PATH")
			info, err := os.Stat(dbPath)
			if err != nil {
				return c.WriteJSON(fiber.Map{"error": "Failed to stat DB file: " + err.Error()})
			}
			modTime := info.ModTime()
			if modTime.Equal(lastModTime) {
				// No file change, skip heavy PRAGMAs
				return nil
			}
			lastModTime = modTime

			// Retrieve PRAGMA statistics
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

			absPath, err := filepath.Abs(dbPath)
			if err != nil {
				absPath = "Error: " + err.Error()
			}

			sqlDB, err := db.DB()
			if err != nil {
				return c.WriteJSON(fiber.Map{"error": "Failed to get underlying sql.DB"})
			}
			dStats := sqlDB.Stats()

			// Build payload
			current := map[string]interface{}{
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
				"open_connections": dStats.OpenConnections,
			}

			// Only send if payload changed
			if reflect.DeepEqual(current, lastStats) {
				return nil
			}
			lastStats = current
			return c.WriteJSON(current)
		}

		// Immediately send once
		if err := send(); err != nil {
			c.Close()
			return
		}

		// Periodically check every 5 seconds
		for range ticker.C {
			if err := send(); err != nil {
				c.Close()
				return
			}
		}
	})
}
