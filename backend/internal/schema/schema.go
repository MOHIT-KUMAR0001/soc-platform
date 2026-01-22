package schema

// my agent data schema
type MsgPack struct {
	StreamID    string  `msgpack:"stream_id" json:"stream_id"`
	WorkerName  string  `msgpack:"worker_name" json:"worker_name"`
	StreamCount int     `msgpack:"stream_count" json:"stream_count"`
	StreamSize  int     `msgpack:"stream_size" json:"stream_size"`
	SendAt      string  `msgpack:"send_at" json:"send_at"`
	Events      []Event `msgpack:"events" json:"events"`
}

// raw strems converter schems
type LogEntry struct {
	Timestamp  string `json:"timestamp"`
	SourceIP   string `json:"source_ip"`
	Method     string `json:"method"`
	Path       string `json:"path"`
	StatusCode int    `json:"status_code"`
	Bytes      int    `json:"bytes"`
	UserAgent  string `json:"userAgent"`
	AgentID    string `json:"agent_id"`
	Hostname   string `json:"hostname"`
}

type Element struct {
	Count     int    `json:"count"`
	FirstSeen string `json:"firstseen"`
	LastSeen  string `json:"lastseen"`
}

type ScanIP struct {
	Source Element `json:"source"`
}
