package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"gorm.io/gorm"
)

type overviewRow struct {
	TotalBytes       uint64 `gorm:"column:total_bytes"`
	TotalPackets     uint64 `gorm:"column:total_packets"`
	TotalFlowCount   uint64 `gorm:"column:total_flow_count"`
	AnomalyFlowCount uint64 `gorm:"column:anomaly_flow_count"`
	TCPSum           uint64 `gorm:"column:tcp_sum"`
	UDPSum           uint64 `gorm:"column:udp_sum"`
	ICMPSum          uint64 `gorm:"column:icmp_sum"`
	OtherSum         uint64 `gorm:"column:other_sum"`
}

func GetOverviewStats(db *gorm.DB) fiber.Handler {
	return websocket.New(func(conn *websocket.Conn) {
		var last overviewRow
		if err := sendOverviewOnce(db, conn, &last, true); err != nil {
			return
		}

		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			if err := sendOverviewOnce(db, conn, &last, false); err != nil {
				return
			}
		}
	})
}

func sendOverviewOnce(db *gorm.DB, conn *websocket.Conn, lastPtr *overviewRow, force bool) error {
	var row overviewRow
	db.Raw(`
		SELECT
		  SUM(sum_dOctets) AS total_bytes,
		  SUM(sum_dPkts)   AS total_packets,
		  SUM(count)       AS total_flow_count,
		  SUM(CASE WHEN file_name LIKE '%/netflow/netflow/%' THEN count ELSE 0 END) AS anomaly_flow_count,
		  SUM(prot_tcp_count)  AS tcp_sum,
		  SUM(prot_udp_count)  AS udp_sum,
		  SUM(prot_icmp_count) AS icmp_sum,
		  SUM(prot_other_count) AS other_sum
		FROM (
		  SELECT * FROM file_records
		  ORDER BY id DESCF
		) AS latest
	`).Scan(&row)

	if !force && row == *lastPtr {
		return nil
	}
	*lastPtr = row

	totalProto := float64(row.TCPSum + row.UDPSum + row.ICMPSum + row.OtherSum)
	dist := map[string]float64{}
	if totalProto > 0 {
		dist["tcp"] = float64(row.TCPSum) / totalProto
		dist["udp"] = float64(row.UDPSum) / totalProto
		dist["icmp"] = float64(row.ICMPSum) / totalProto
		dist["other"] = float64(row.OtherSum) / totalProto
	}

	stats := fiber.Map{
		"total_bytes":           row.TotalBytes,
		"total_packets":         row.TotalPackets,
		"total_flow_count":      row.TotalFlowCount,
		"anomaly_flow_count":    row.AnomalyFlowCount,
		"protocol_distribution": dist,
	}

	return conn.WriteJSON(stats)
}
