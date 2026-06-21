package solutions

// ProcessedSet tracks handled message IDs for idempotency.
type ProcessedSet struct{ seen map[string]bool }

func NewProcessedSet() *ProcessedSet { return &ProcessedSet{seen: make(map[string]bool)} }

func (p *ProcessedSet) Handle(id string, fn func() error) error {
	if p.seen[id] { return nil }
	if err := fn(); err != nil { return err }
	p.seen[id] = true
	return nil
}