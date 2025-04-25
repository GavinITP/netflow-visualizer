package controllers

import (
	"bufio"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func GetRecentLogs(logFilePath string, n int) fiber.Handler {
	return websocket.New(func(conn *websocket.Conn) {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				file, err := os.Open(logFilePath)
				if err != nil {
					conn.WriteJSON(fiber.Map{"error": "Unable to open log file: " + err.Error()})
					continue
				}

				var lines []string
				scanner := bufio.NewScanner(file)
				for scanner.Scan() {
					lines = append(lines, scanner.Text())
				}
				file.Close()

				if scannerErr := scanner.Err(); scannerErr != nil {
					conn.WriteJSON(fiber.Map{"error": "Error reading log file: " + scannerErr.Error()})
					continue
				}

				if len(lines) > n {
					lines = lines[len(lines)-n:]
				}

				result := strings.Join(lines, "\n")
				if err := conn.WriteJSON(fiber.Map{"recent_logs": result}); err != nil {
					return
				}
			}
		}
	})
}
