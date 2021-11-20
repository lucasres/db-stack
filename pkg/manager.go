package pkg

import (
	"errors"

	"github.com/lucasres/db-stack/internal"
)

type Manager struct {
	keyValues map[string]*internal.KeyValue
}

func NewManager() *Manager {
	return &Manager{
		keyValues: make(map[string]*internal.KeyValue),
	}
}

func (m *Manager) Set(k, v string) {
	keyValue := internal.NewKeyValue(k)

	keyValue.Push(v)

	m.keyValues[k] = keyValue
}

func (m *Manager) Push(k, v string) error {
	if kv, ok := m.keyValues[k]; ok {
		kv.Push(v)
	}

	return errors.New("try push in undefined key")
}

func (m *Manager) Get(k string) (string, error) {
	if kv, ok := m.keyValues[k]; ok {
		return kv.Get(), nil
	}

	return "", errors.New("try get in undefined key")
}

func (m *Manager) Pop(k string) error {
	if kv, ok := m.keyValues[k]; ok {
		kv.Pop()
		return nil
	}

	return errors.New("try pop in undefined key")
}

func (m *Manager) Flush() {
	m.keyValues = make(map[string]*internal.KeyValue, 0)
}
