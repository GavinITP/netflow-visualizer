package controllers

import (
	"encoding/csv"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"netflow-visualizer/models"
)

func GetNetflowsFromDBFilePath(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var fileRecord models.FilePath
		if err := db.First(&fileRecord).Error; err != nil {
			return c.Status(http.StatusInternalServerError).
				JSON(fiber.Map{"error": "CSV file path not found"})
		}

		search := c.Query("search")
		portStr := c.Query("port")
		fromStr := c.Query("from")
		toStr := c.Query("to")
		limitStr := c.Query("limit")

		var portFilter int
		if portStr != "" {
			if p, err := strconv.Atoi(portStr); err == nil {
				portFilter = p
			}
		}

		var fromTime, toTime time.Time
		if fromStr != "" {
			if t, err := time.Parse(time.RFC3339, fromStr); err == nil {
				fromTime = t
			} else {
				return c.Status(http.StatusBadRequest).
					JSON(fiber.Map{"error": "invalid from timestamp"})
			}
		}
		if toStr != "" {
			if t, err := time.Parse(time.RFC3339, toStr); err == nil {
				toTime = t
			} else {
				return c.Status(http.StatusBadRequest).
					JSON(fiber.Map{"error": "invalid to timestamp"})
			}
		}

		var limit int
		if limitStr != "" {
			if n, err := strconv.Atoi(limitStr); err == nil && n > 0 {
				limit = n
			}
		}

		f, err := os.Open(fileRecord.Path)
		if err != nil {
			return c.Status(http.StatusInternalServerError).
				JSON(fiber.Map{"error": "Unable to open CSV file"})
		}
		defer f.Close()

		r := csv.NewReader(f)
		if _, err := r.Read(); err != nil {
			return c.Status(http.StatusInternalServerError).
				JSON(fiber.Map{"error": "Failed to read CSV header"})
		}

		var netflows []models.Netflow
		for {
			row, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil || len(row) != 11 {
				continue
			}

			id, _ := strconv.ParseUint(row[0], 10, 32)
			ts, _ := time.Parse(time.RFC3339, row[1])
			dPkts, _ := strconv.ParseUint(row[5], 10, 64)
			dOctets, _ := strconv.ParseUint(row[6], 10, 64)
			srcPort, _ := strconv.Atoi(row[7])
			dstPort, _ := strconv.Atoi(row[8])
			tos, _ := strconv.Atoi(row[10])

			flow := models.Netflow{
				ID:        uint(id),
				Timestamp: ts,
				SrcAddr:   row[2],
				DstAddr:   row[3],
				NextHop:   row[4],
				DPkts:     dPkts,
				DOctets:   dOctets,
				SrcPort:   srcPort,
				DstPort:   dstPort,
				Prot:      row[9],
				Tos:       tos,
			}

			if search != "" && flow.SrcAddr != search && flow.DstAddr != search {
				continue
			}
			if portFilter != 0 && flow.SrcPort != portFilter && flow.DstPort != portFilter {
				continue
			}
			if !fromTime.IsZero() && flow.Timestamp.Before(fromTime) {
				continue
			}
			if !toTime.IsZero() && flow.Timestamp.After(toTime) {
				continue
			}

			netflows = append(netflows, flow)
		}

		if limit > 0 && len(netflows) > limit {
			netflows = netflows[len(netflows)-limit:]
		}

		return c.JSON(netflows)
	}
}
