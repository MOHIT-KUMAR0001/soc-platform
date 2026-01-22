package ingest

import (
	"backend/internal/schema"
	"fmt"
	"regexp"
	"strconv"
)

type RawLogParser struct{}

func (p *RawLogParser) Process(e *schema.Event) {
	re := regexp.MustCompile(`^(\S+)\s+\S+\s+\S+\s+\[(.*?)\]\s+"(\S+)\s+(\S+)\s+.*?"\s+(\d{3})\s+(\d+|-)\s+".*?"\s+"(.*?)"`)

	match := re.FindStringSubmatch(e.Raw)

	statusCode, _ := strconv.Atoi(match[5])
	bytesSent, _ := strconv.Atoi(match[6])

	data := &schema.LogEntry{
		Timestamp:  match[2],
		SourceIP:   match[1],
		Method:     match[3],
		Path:       match[4],
		StatusCode: statusCode,
		Bytes:      bytesSent,
		UserAgent:  match[7],
		AgentID:    e.AgentID,
		Hostname:   e.Hostname,
	}

	e.ParseLog = data
}

var database schema.ScanIP

type FailedLoginDetector struct{}

func (p *FailedLoginDetector) Process(e *schema.Event) {

	count := &database.Source.Count

	if database.IP != e.ParseLog.SourceIP {
		database.IP = e.ParseLog.SourceIP
		database.Source.Count = 1
		database.Source.FirstSeen = e.ParseLog.Timestamp
		return
	}

	*count++

	if *count > 5 {
		fmt.Println("heavy straming dectected...")
	}

}
