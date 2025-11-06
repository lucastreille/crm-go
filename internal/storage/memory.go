package storage

import (
	"errors"
	"sync"
)

type Memory struct {
	mu       sync.RWMutex
	contacts map[int]*Contact
}

func NewMemory() *Memory {

	return &Memory{
		contacts: make(map[int]*Contact),
	}

}

func (m *Memory) Add(c *Contact) error {

	m.mu.Lock()
	defer m.mu.Unlock()
	if _, exists := m.contacts[c.ID]; exists {
		return errors.New("contact déjà existant")
	}
	cp := *c
	m.contacts[c.ID] = &cp
	return nil

}

func (m *Memory) List() ([]*Contact, error) {

	m.mu.RLock()
	defer m.mu.RUnlock()
	out := make([]*Contact, 0, len(m.contacts))
	for _, c := range m.contacts {
		cp := *c
		out = append(out, &cp)
	}
	return out, nil

}

func (m *Memory) Delete(id int) error {

	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.contacts[id]; !ok {
		return errors.New("contact introuvable")
	}
	delete(m.contacts, id)
	return nil

}

func (m *Memory) Update(c *Contact) error {

	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.contacts[c.ID]; !ok {
		return errors.New("contact introuvable")
	}
	cp := *c
	m.contacts[c.ID] = &cp
	return nil

}

func (m *Memory) Get(id int) (*Contact, error) {

	m.mu.RLock()
	defer m.mu.RUnlock()
	c, ok := m.contacts[id]
	if !ok {
		return nil, errors.New("contact introuvable")
	}
	cp := *c
	return &cp, nil

}
