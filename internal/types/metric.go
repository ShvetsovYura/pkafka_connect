package types

type Metric struct {
	Type        string // gauge,
	Name        string // Alloc,
	Description string // Alloc is bytes of allocated heap objects.,
	Value       int    // 24293912
}
