package controllers

import (
	"encoding/csv"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"netflow-visualizer/models"
)

func GetAnomalyLogs(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		search := c.Query("search")
		portFilter := 0
		if portStr := c.Query("port"); portStr != "" {
			p, err := strconv.Atoi(portStr)
			if err != nil {
				return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid port"})
			}
			portFilter = p
		}

		protoParam := strings.ToUpper(c.Query("protocol"))
		validProtos := map[string]bool{"TCP": true, "UDP": true, "ICMP": true, "OTHERS": true}
		if protoParam != "" && !validProtos[protoParam] {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid protocol"})
		}

		var recs []models.FileRecord
		if err := db.Order("id desc").Find(&recs).Error; err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "db query failed: " + err.Error()})
		}

		var selected []models.FileRecord
		recentCountStr := c.Query("recent_count")
		if recentCountStr == "" {
			countAnoms := 0
			for _, r := range recs {
				if strings.Contains(r.FileName, "/netflow/") {
					selected = append(selected, r)
					countAnoms++
					if countAnoms >= 20 {
						break
					}
				}
			}
		} else {
			recentCount, err := strconv.ParseUint(recentCountStr, 10, 64)
			if err != nil {
				return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid recent_count"})
			}
			var sumCount uint64
			for _, r := range recs {
				if !strings.Contains(r.FileName, "/netflow/") {
					continue
				}
				selected = append(selected, r)
				sumCount += r.Count
				if sumCount >= recentCount {
					break
				}
			}
			if sumCount < recentCount || len(selected) < 20 {
				selected = nil
				countAnoms := 0
				for _, r := range recs {
					if strings.Contains(r.FileName, "/netflow/") {
						selected = append(selected, r)
						countAnoms++
						if countAnoms >= 20 {
							break
						}
					}
				}
			}
		}

		var flows []models.AnomalyNetflow
		for _, rec := range selected {
			f, err := os.Open(rec.FileName)
			if err != nil {
				continue
			}
			r := csv.NewReader(f)
			if _, err := r.Read(); err != nil {
				f.Close()
				continue
			}

			for {
				row, err := r.Read()
				if err == io.EOF {
					break
				}
				if err != nil || len(row) < 18 {
					continue
				}

				input, _ := strconv.Atoi(row[3])
				output, _ := strconv.Atoi(row[4])
				dPkts, _ := strconv.ParseUint(row[5], 10, 64)
				dOctets, _ := strconv.ParseUint(row[6], 10, 64)

				var firstTime, lastTime time.Time
				if t, err := time.Parse(time.RFC3339, row[7]); err == nil {
					firstTime = t
				} else if secs, err2 := strconv.ParseInt(row[7], 10, 64); err2 == nil {
					firstTime = time.Unix(secs, 0)
				}
				if t2, err := time.Parse(time.RFC3339, row[8]); err == nil {
					lastTime = t2
				} else if secs2, err2 := strconv.ParseInt(row[8], 10, 64); err2 == nil {
					lastTime = time.Unix(secs2, 0)
				}

				srcPort, _ := strconv.Atoi(row[9])
				dstPort, _ := strconv.Atoi(row[10])
				tcpFlags := row[11]
				protCode, _ := strconv.Atoi(row[12])
				tosVal, _ := strconv.Atoi(row[13])
				srcAS, _ := strconv.Atoi(row[14])
				dstAS, _ := strconv.Atoi(row[15])
				srcMask, _ := strconv.Atoi(row[16])
				dstMask, _ := strconv.Atoi(row[17])

				if protoParam != "" {
					switch protoParam {
					case "TCP":
						if protCode != 6 {
							continue
						}
					case "UDP":
						if protCode != 17 {
							continue
						}
					case "ICMP":
						if protCode != 1 {
							continue
						}
					case "OTHERS":
						if protCode == 6 || protCode == 17 || protCode == 1 {
							continue
						}
					}
				}

				if search != "" && row[0] != search && row[1] != search {
					continue
				}
				if portFilter != 0 && srcPort != portFilter && dstPort != portFilter {
					continue
				}

				flows = append(flows, models.AnomalyNetflow{
					BaseNetflow: models.BaseNetflow{
						SrcAddr: row[0],
						DstAddr: row[1],
						NextHop: row[2],
						DPkts:   dPkts,
						DOctets: dOctets,
						SrcPort: srcPort,
						DstPort: dstPort,
						Prot:    row[12],
						Tos:     tosVal,
					},
					Input:    input,
					Output:   output,
					First:    firstTime,
					Last:     lastTime,
					TCPFlags: tcpFlags,
					SrcAS:    srcAS,
					DstAS:    dstAS,
					SrcMask:  srcMask,
					DstMask:  dstMask,
				})
			}
			f.Close()
		}

		return c.JSON(flows)
	}
}
