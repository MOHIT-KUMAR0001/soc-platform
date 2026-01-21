package ingest

import (
	"backend/internal/schema"
	"fmt"
	"regexp"
	"strconv"
)

func ParseRawLog(raw, agentID, hostname string) (schema.LogEntry, error) {
	fmt.Println(raw)
	// Regex breakdown:
	// ^(\S+)           -> Source IP
	// .*?\[(.*?)\]     -> Timestamp inside brackets
	// "(.*?) (.*?) .*?" -> Method and Path inside quotes
	// (\d{3})          -> Status Code
	// (\d+|-)          -> Bytes (handles '-' if no bytes sent)
	// ".*?" "(.*?)"    -> Referrer (ignored) and User Agent
	re := regexp.MustCompile(`^(\S+)\s+\S+\s+\S+\s+\[(.*?)\]\s+"(\S+)\s+(\S+)\s+.*?"\s+(\d{3})\s+(\d+|-)\s+".*?"\s+"(.*?)"`)

	match := re.FindStringSubmatch(raw)
	if len(match) < 8 {
		return schema.LogEntry{}, fmt.Errorf("raw log did not match pattern")
	}

	statusCode, _ := strconv.Atoi(match[5])
	bytesSent, _ := strconv.Atoi(match[6]) // Note: will be 0 if match is "-"

	return schema.LogEntry{
		Timestamp:  match[2], // You may need time.Parse to change format
		SourceIP:   match[1],
		Method:     match[3],
		Path:       match[4],
		StatusCode: statusCode,
		Bytes:      bytesSent,
		UserAgent:  match[7],
		AgentID:    agentID,
		Hostname:   hostname,
	}, nil
}
