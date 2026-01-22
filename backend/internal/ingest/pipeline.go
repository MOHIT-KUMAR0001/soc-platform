package ingest

import (
	"backend/internal/schema"
)

type Pipeline struct {
	Processors []Processor
}

func (p *Pipeline) Run(event *schema.Event) {
	for _, proc := range p.Processors {
		proc.Process(event)
	}
}
