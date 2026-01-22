package ingest

import (
	"backend/internal/schema"
	"fmt"
	"regexp"
	"strconv"
)

type RawLogParser struct{}

func (p *RawLogParser) Process(e *schema.Event) (*schema.LogEntry, error) {
	re := regexp.MustCompile(`^(\S+)\s+\S+\s+\S+\s+\[(.*?)\]\s+"(\S+)\s+(\S+)\s+.*?"\s+(\d{3})\s+(\d+|-)\s+".*?"\s+"(.*?)"`)

	match := re.FindStringSubmatch(e.Raw)
	if len(match) < 8 {
		return &schema.LogEntry{}, fmt.Errorf("raw log did not match pattern")
	}

	statusCode, _ := strconv.Atoi(match[5])
	bytesSent, _ := strconv.Atoi(match[6])

	return &schema.LogEntry{
		Timestamp:  match[2],
		SourceIP:   match[1],
		Method:     match[3],
		Path:       match[4],
		StatusCode: statusCode,
		Bytes:      bytesSent,
		UserAgent:  match[7],
		AgentID:    e.AgentID,
		Hostname:   e.Hostname,
	}, nil
}
