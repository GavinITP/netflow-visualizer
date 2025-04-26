package controllers

import (
	"bufio"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func GetRecentLogs() fiber.Handler {
	logFilePath := os.Getenv("LOGS_FILE_PATH")
	return websocket.New(func(conn *websocket.Conn) {
		var lastResult string

		send := func() error {
			file, err := os.Open(logFilePath)
			if err != nil {
				return conn.WriteJSON(fiber.Map{"error": "Unable to open log file: " + err.Error()})
			}
			defer file.Close()

			var lines []string
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				lines = append(lines, scanner.Text())
			}
			if scanErr := scanner.Err(); scanErr != nil {
				return conn.WriteJSON(fiber.Map{"error": "Error reading log file: " + scanErr.Error()})
			}

			if len(lines) > 50 {
				lines = lines[len(lines)-50:]
			}

			result := strings.Join(lines, "\n")
			if result == lastResult {
				return nil
			}
			lastResult = result
			return conn.WriteJSON(fiber.Map{"recent_logs": result})
		}

		if err := send(); err != nil {
			conn.Close()
			return
		}

		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			if err := send(); err != nil {
				conn.Close()
				return
			}
		}
	})
}
