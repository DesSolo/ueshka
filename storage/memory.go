package storage

// Memory ...
type Memory struct {
	data map[string]bool
}

// NewMemory ...
func NewMemory() *Memory {
	return &Memory{
		data: make(map[string]bool),
	}
}

// Add ...
func (m *Memory) Add(k string) {
	m.data[k] = true
}

// IsExist ...
func (m *Memory) IsExist(k string) bool {
	_, ok := m.data[k]
	return ok
}
