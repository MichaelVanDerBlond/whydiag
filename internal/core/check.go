package core

// Check — базовый интерфейс любой проверки.
type Check interface {
	Descriptor() Descriptor
	Run(*Context) Result
}
