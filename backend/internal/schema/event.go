package schema

type Signal struct {
	Type string
	Data map[string]any
}

type Event struct {
	EventType string `msgpack:"event_type" json:"event_type"`
	AgentID   string `msgpack:"agent_id" json:"agent_id"`
	Hostname  string `msgpack:"hostname" json:"hostname"`
	Source    string `msgpack:"source" json:"source"`
	Timestamp string `msgpack:"timestamp" json:"timestamp"`
	Raw       string `msgpack:"raw" json:"raw"`

	Fields   map[string]string
	Tags     map[string]struct{}
	ParseLog *LogEntry
	Signals  []Signal
}
