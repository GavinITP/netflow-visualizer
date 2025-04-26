package controllers

import (
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"gorm.io/gorm"

	"netflow-visualizer/models"
)

type FlowHistoryEntry struct {
	Time  string `json:"time"`
	Count uint64 `json:"count"`
}

func GetFlowHistory(db *gorm.DB) fiber.Handler {
	return websocket.New(func(conn *websocket.Conn) {
		const layout = "20060102_150405"
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		var lastRsp []FlowHistoryEntry

		send := func() error {

			var latestRec struct {
				FileName string
				Count    uint64
			}
			if err := db.Model(&models.FileRecord{}).
				Select("file_name, count").
				Where("file_name LIKE ?", "%/netflow/%").
				Order("file_name DESC").
				Limit(1).
				Find(&latestRec).Error; err != nil {
				return err
			}

			base := filepath.Base(latestRec.FileName)
			name := strings.TrimSuffix(base, filepath.Ext(base))
			parts := strings.SplitN(name, "_", 2)
			latestTs := time.Now().Add(-10 * time.Second)
			if len(parts) == 2 {
				if ts, err := time.Parse(layout, parts[1]); err == nil {
					latestTs = ts.Truncate(time.Second)
				}
			}

			var recs []struct {
				FileName string
				Count    uint64
			}
			if err := db.Model(&models.FileRecord{}).
				Select("file_name, count").
				Where("file_name LIKE ? AND file_name < ?", "%/netflow/%", latestRec.FileName).
				Order("file_name DESC").
				Limit(9).
				Find(&recs).Error; err != nil {
				return err
			}

			counts := make(map[int64]uint64, len(recs)+1)
			counts[latestTs.Unix()] = latestRec.Count
			for _, r := range recs {
				base := filepath.Base(r.FileName)
				name := strings.TrimSuffix(base, filepath.Ext(base))
				parts := strings.SplitN(name, "_", 2)
				if len(parts) != 2 {
					continue
				}
				if ts, err := time.Parse(layout, parts[1]); err == nil {
					counts[ts.Truncate(time.Second).Unix()] = r.Count
				}
			}

			start := latestTs.Add(-9 * time.Second)
			rsp := make([]FlowHistoryEntry, 0, 10)
			for i := 0; i < 10; i++ {
				t := start.Add(time.Duration(i) * time.Second)
				rsp = append(rsp, FlowHistoryEntry{
					Time:  t.Format("15:04:05"),
					Count: counts[t.Unix()],
				})
			}

			if reflect.DeepEqual(rsp, lastRsp) {
				return nil
			}
			lastRsp = append([]FlowHistoryEntry(nil), rsp...)
			return conn.WriteJSON(rsp)
		}

		if err := send(); err != nil {
			conn.Close()
			return
		}
		for range ticker.C {
			if err := send(); err != nil {
				conn.Close()
				return
			}
		}
	})
}
