package brass

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type Manager[T any] struct {
	mu       sync.Mutex
	value    T
	filePath string
}

// Create a new manager and optionally load from file
func NewManager[T any](filePath string) *Manager[T] {
	m := &Manager[T]{filePath: filePath}
	m.loadFromFile()
	return m
}

// Set replaces the value and saves it
func (m *Manager[T]) Set(newValue T) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.value = newValue
	m.saveToFile()
}

// Get returns a copy of the value
func (m *Manager[T]) Get() T {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.value
}

// Save value to file
func (m *Manager[T]) saveToFile() {
	data, err := json.MarshalIndent(m.value, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling:", err)
		return
	}

	err = os.WriteFile(m.filePath, data, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}
}

// Load value from file (if exists)
func (m *Manager[T]) loadFromFile() {
	data, err := os.ReadFile(m.filePath)
	if err != nil {
		return // file doesn't exist or unreadable
	}

	err = json.Unmarshal(data, &m.value)
	if err != nil {
		fmt.Println("Error unmarshaling:", err)
	}
}
