package controllers

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"gorm.io/gorm"

	"netflow-visualizer/models"
)

func GetOverviewStats(db *gorm.DB) fiber.Handler {
	return websocket.New(func(conn *websocket.Conn) {
		for {
			var recs []models.FileRecord
			if err := db.Order("id DESC").Limit(10).Find(&recs).Error; err != nil {
				conn.WriteJSON(fiber.Map{"error": "Failed to query file_records: " + err.Error()})
				return
			}

			var totalBytes, totalPackets, totalFlowCount, anomalyFlowCount uint64
			var tcpSum, udpSum, icmpSum, otherSum uint64
			for _, r := range recs {
				totalBytes += r.SumDOctets
				totalPackets += r.SumDPkts
				totalFlowCount += r.Count
				// anomaly rows contain "/netflow/" in filename
				if strings.Contains(r.FileName, "/netflow/") {
					anomalyFlowCount += r.Count
				}
				tcpSum += r.ProtTCPCount
				udpSum += r.ProtUDPCount
				icmpSum += r.ProtICMPCount
				otherSum += r.ProtOtherCount
			}

			totalProto := tcpSum + udpSum + icmpSum + otherSum
			protoDist := make(map[string]float64)
			if totalProto > 0 {
				protoDist["tcp"] = float64(tcpSum) / float64(totalProto)
				protoDist["udp"] = float64(udpSum) / float64(totalProto)
				protoDist["icmp"] = float64(icmpSum) / float64(totalProto)
				protoDist["other"] = float64(otherSum) / float64(totalProto)
			}

			stats := fiber.Map{
				"total_bytes":           totalBytes,
				"total_packets":         totalPackets,
				"total_flow_count":      totalFlowCount,
				"anomaly_flow_count":    anomalyFlowCount,
				"protocol_distribution": protoDist,
			}

			if err := conn.WriteJSON(stats); err != nil {
				return
			}

			select {
			case <-time.After(5 * time.Second):
			}
		}
	})
}
