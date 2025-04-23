package controllers

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"gorm.io/gorm"

	"netflow-visualizer/models"
)

var serverStartTime = time.Now()

func GetNetflowStatsFromDB(db *gorm.DB) fiber.Handler {
	return websocket.New(func(conn *websocket.Conn) {
		var fp models.FilePath
		if err := db.First(&fp).Error; err != nil {
			conn.WriteJSON(fiber.Map{"error": "CSV file path not found"})
			conn.Close()
			return
		}

		readRows := func() ([]models.Netflow, error) {
			f, err := os.Open(fp.Path)
			if err != nil {
				return nil, err
			}
			defer f.Close()

			r := csv.NewReader(f)
			if _, err := r.Read(); err != nil {
				return nil, err
			}

			var rows []models.Netflow
			for {
				rec, err := r.Read()
				if err == io.EOF {
					break
				}
				if err != nil || len(rec) != 11 {
					continue
				}

				id, _ := strconv.ParseUint(rec[0], 10, 32)
				ts, _ := time.Parse(time.RFC3339, rec[1])
				dPkts, _ := strconv.ParseUint(rec[5], 10, 64)
				srcPort, _ := strconv.Atoi(rec[7])
				dstPort, _ := strconv.Atoi(rec[8])

				rows = append(rows, models.Netflow{
					ID:        uint(id),
					Timestamp: ts,
					SrcAddr:   rec[2],
					DstAddr:   rec[3],
					NextHop:   rec[4],
					DPkts:     dPkts,
					SrcPort:   srcPort,
					DstPort:   dstPort,
					Prot:      rec[9],
				})
			}
			return rows, nil
		}

		initialRows, err := readRows()
		if err != nil {
			conn.WriteJSON(fiber.Map{"error": "Unable to read CSV: " + err.Error()})
			conn.Close()
			return
		}
		initialCount := len(initialRows)
		prevCount := initialCount

		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		for range ticker.C {
			rows, err := readRows()
			if err != nil {
				conn.WriteJSON(fiber.Map{"error": "Unable to read CSV: " + err.Error()})
				continue
			}

			var totalPkts uint64
			protCounts := make(map[string]int)
			for _, nf := range rows {
				totalPkts += nf.DPkts
				protCounts[nf.Prot]++
			}

			currentCount := len(rows)
			packetsPerSec := currentCount - prevCount
			prevCount = currentCount
			activeAlerts := currentCount - initialCount

			protoDist := make(map[string]float64, len(protCounts))
			if currentCount > 0 {
				for p, cnt := range protCounts {
					protoDist[p] = float64(cnt) / float64(currentCount)
				}
			}

			elapsed := time.Since(serverStartTime)
			totalSecs := int(elapsed.Seconds())
			h := totalSecs / 3600
			m := (totalSecs % 3600) / 60
			s := totalSecs % 60

			stats := fiber.Map{
				"total_packets":      totalPkts,
				"packets_per_second": packetsPerSec,
				"active_alert":       activeAlerts,
				"uptime": fiber.Map{
					"hours":   h,
					"minutes": m,
					"seconds": s,
				},
				"protocol_distribution": protoDist,
			}

			if err := conn.WriteJSON(stats); err != nil {
				return
			}
		}
	})
}
