package core

// Severity определяет важность проверки.
type Severity uint8

const (
	SeverityInfo Severity = iota
	SeverityWarning
	SeverityCritical
)

// Descriptor содержит метаданные проверки.
type Descriptor struct {
	ID          string
	Name        string
	Category    string
	Description string
	Severity    Severity
}
