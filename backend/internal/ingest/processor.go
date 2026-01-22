package ingest

import "backend/internal/schema"

type Processor interface {
	Process(event *schema.Event)
}
