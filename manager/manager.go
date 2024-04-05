package manager

import (
	"log/slog"
	"sync"
)

type Manager struct {
	Clients map[string][]chan string
	mu      sync.Mutex
	logger  *slog.Logger
}

func New() *Manager {
	clients := make(map[string][]chan string)
	return &Manager{
		Clients: clients,
	}
}

func (m *Manager) SendMessage(watcherID string, msg string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if chs, ok := m.Clients[watcherID]; ok {
		for _, ch := range chs {
			ch <- msg
		}
	}
}

func (m *Manager) AddClient(watcherID string, clientChan chan string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Clients[watcherID] = append(m.Clients[watcherID], clientChan)
}

func (m *Manager) RemoveClient(watcherID string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if len(m.Clients[watcherID]) > 1 {
		m.Clients[watcherID] = m.Clients[watcherID][:len(m.Clients[watcherID])-1]
	} else {
		delete(m.Clients, watcherID)
	}
}
