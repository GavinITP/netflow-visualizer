package controllers

import (
	"bufio"
	"bytes"
	"io"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"netflow-visualizer/models"
)

var (
	flowPool    = sync.Pool{New: func() interface{} { return make([]models.AnomalyNetflow, 0, 1024) }}
	lineBufPool = sync.Pool{New: func() interface{} { return make([]byte, 0, 1<<20) }}
)

func GetAnomalyLogs(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {

		search := c.Query("search")
		searchB := []byte(search)

		portFilter := 0
		if ps := c.Query("port"); ps != "" {
			p, err := strconv.Atoi(ps)
			if err != nil {
				return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid port"})
			}
			portFilter = p
		}

		protoParam := strings.ToUpper(c.Query("protocol"))
		if protoParam != "" {
			switch protoParam {
			case "TCP", "UDP", "ICMP", "OTHERS":
			default:
				return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid protocol"})
			}
		}

		var recs []models.FileRecord
		if err := db.Order("id desc").Find(&recs).Error; err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "db query failed"})
		}

		selected := selectAnomalyFiles(recs, c.Query("recent_count"))

		numCPU := runtime.NumCPU()
		sem := make(chan struct{}, numCPU)
		outCh := make(chan []models.AnomalyNetflow, len(selected))
		var wg sync.WaitGroup

		for _, rec := range selected {
			wg.Add(1)
			sem <- struct{}{}
			go func(fn string) {
				defer wg.Done()
				defer func() { <-sem }()
				outCh <- parseFile(fn, searchB, portFilter, protoParam)
			}(rec.FileName)
		}
		wg.Wait()
		close(outCh)

		all := make([]models.AnomalyNetflow, 0, len(selected)*1024)
		for fs := range outCh {
			all = append(all, fs...)
			flowPool.Put(fs[:0])
		}

		if rc := c.Query("recent_count"); rc != "" {
			if limit, err := strconv.Atoi(rc); err == nil && limit > 0 && len(all) > limit {
				all = all[:limit]
			}
		}

		return c.JSON(all)
	}
}

func selectAnomalyFiles(recs []models.FileRecord, recentCountStr string) []models.FileRecord {
	if recentCountStr == "" {
		sel := make([]models.FileRecord, 0, 10)
		for _, r := range recs {
			if strings.Contains(r.FileName, "/netflow/netflow/") && len(sel) < 10 {
				sel = append(sel, r)
			}
		}
		return sel
	}

	return selectAnomalyFiles(recs, "")
}

func parseFile(filename string, searchB []byte, portFilter int, protoParam string) []models.AnomalyNetflow {
	f, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer f.Close()

	rdr := bufio.NewReaderSize(f, 1<<20)
	flows := flowPool.Get().([]models.AnomalyNetflow)
	buf := lineBufPool.Get().([]byte)
	defer lineBufPool.Put(buf[:0])

	if _, err := rdr.ReadBytes('\n'); err != nil {
		return flows
	}

	for {
		line, err := rdr.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		buf = append(buf[:0], line...)

		var flds [18][]byte
		start := 0
		valid := true
		for i := 0; i < 17; i++ {
			idx := bytes.IndexByte(buf[start:], ',')
			if idx < 0 {
				valid = false
				break
			}
			flds[i] = buf[start : start+idx]
			start += idx + 1
		}
		if !valid {
			continue
		}
		flds[17] = buf[start:]

		if len(searchB) > 0 && !bytes.Equal(flds[0], searchB) && !bytes.Equal(flds[1], searchB) {
			continue
		}
		sPort, _ := strconv.Atoi(string(flds[9]))
		dPort, _ := strconv.Atoi(string(flds[10]))
		if portFilter != 0 && sPort != portFilter && dPort != portFilter {
			continue
		}
		pCode, _ := strconv.Atoi(string(flds[12]))
		if protoParam != "" {
			switch protoParam {
			case "TCP":
				if pCode != 6 {
					continue
				}
			case "UDP":
				if pCode != 17 {
					continue
				}
			case "ICMP":
				if pCode != 1 {
					continue
				}
			case "OTHERS":
				if pCode == 6 || pCode == 17 || pCode == 1 {
					continue
				}
			}
		}

		iVal, _ := strconv.Atoi(string(flds[3]))
		oVal, _ := strconv.Atoi(string(flds[4]))
		dPkts, _ := strconv.ParseUint(string(flds[5]), 10, 64)
		dOctets, _ := strconv.ParseUint(string(flds[6]), 10, 64)
		fTime := parseTime(string(flds[7]))
		lTime := parseTime(string(flds[8]))
		tFlags := string(flds[11])
		tos, _ := strconv.Atoi(string(flds[13]))
		sAS, _ := strconv.Atoi(string(flds[14]))
		dAS, _ := strconv.Atoi(string(flds[15]))
		sMask, _ := strconv.Atoi(string(flds[16]))
		dMask, _ := strconv.Atoi(string(flds[17]))

		flows = append(flows, models.AnomalyNetflow{
			BaseNetflow: models.BaseNetflow{
				SrcAddr: string(flds[0]), DstAddr: string(flds[1]), NextHop: string(flds[2]),
				DPkts: dPkts, DOctets: dOctets, SrcPort: sPort, DstPort: dPort,
				Prot: string(flds[12]), Tos: tos,
			}, Input: iVal, Output: oVal, First: fTime, Last: lTime,
			TCPFlags: tFlags, SrcAS: sAS, DstAS: dAS, SrcMask: sMask, DstMask: dMask,
		})
	}
	return flows
}

func parseTime(s string) time.Time {
	if t, err := time.Parse(time.RFC3339, s); err == nil {
		return t
	}
	if secs, err := strconv.ParseInt(s, 10, 64); err == nil {
		return time.Unix(secs, 0)
	}
	return time.Time{}
}
