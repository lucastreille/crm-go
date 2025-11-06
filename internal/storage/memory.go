package storage

import "errors"

type Memory struct {
	contacts map[int]*Contact
}

func NewMemory() *Memory {
	return &Memory{contacts: make(map[int]*Contact)}
}

func (m *Memory) Add(c *Contact) error {
	if _, exists := m.contacts[c.ID]; exists {
		return errors.New("contact déjà existant")
	}
	m.contacts[c.ID] = c
	return nil
}

func (m *Memory) List() ([]*Contact, error) {
	var result []*Contact
	for _, c := range m.contacts {
		result = append(result, c)
	}
	return result, nil
}

func (m *Memory) Delete(id int) error {
	if _, ok := m.contacts[id]; !ok {
		return errors.New("contact introuvable")
	}
	delete(m.contacts, id)
	return nil
}

func (m *Memory) Update(c *Contact) error {
	if _, ok := m.contacts[c.ID]; !ok {
		return errors.New("contact introuvable")
	}
	m.contacts[c.ID] = c
	return nil
}

func (m *Memory) Get(id int) (*Contact, error) {
	c, ok := m.contacts[id]
	if !ok {
		return nil, errors.New("contact introuvable")
	}
	return c, nil
}
