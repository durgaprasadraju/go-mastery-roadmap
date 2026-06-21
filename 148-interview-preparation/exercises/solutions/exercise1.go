package solutions

// ModuleProgress tracks learning progress.
type ModuleProgress struct {
	Module  string
	Percent int
}

// Completed reports whether module is done.
func (p ModuleProgress) Completed() bool { return p.Percent >= 100 }