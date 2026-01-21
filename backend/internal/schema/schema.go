package schema

type appEvent struct {
	EventType string `msgpack:"event_type" json:"event_type"`
	AgentID   string `msgpack:"agent_id" json:"agent_id"`
	Hostname  string `msgpack:"hostname" json:"hostname"`
	Source    string `msgpack:"source" json:"source"`
	Timestamp string `msgpack:"timestamp" json:"timestamp"`
	Raw       string `msgpack:"raw" json:"raw"`
}

// my agent data schema
type MsgPack struct {
	StreamID    string     `msgpack:"stream_id" json:"stream_id"`
	WorkerName  string     `msgpack:"worker_name" json:"worker_name"`
	StreamCount int        `msgpack:"stream_count" json:"stream_count"`
	StreamSize  int        `msgpack:"stream_size" json:"stream_size"`
	SendAt      string     `msgpack:"send_at" json:"send_at"`
	Events      []appEvent `msgpack:"events" json:"events"`
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
