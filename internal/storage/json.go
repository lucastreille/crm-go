package storage

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"sync"
)

type JSON struct {
	path     string
	mu       sync.RWMutex
	contacts map[int]*Contact
}

func NewJSON(path string) (*JSON, error) {

	j := &JSON{
		path:     path,
		contacts: make(map[int]*Contact),
	}

	if dir := filepath.Dir(path); dir != "." && dir != "" {
		_ = os.MkdirAll(dir, 0o755)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return j, nil
		}
		return nil, err
	}

	var list []*Contact
	if err := json.Unmarshal(data, &list); err != nil {
		return nil, err
	}
	for _, c := range list {
		j.contacts[c.ID] = c
	}
	return j, nil

}

func (j *JSON) saveLocked() error {

	list := make([]*Contact, 0, len(j.contacts))
	for _, c := range j.contacts {
		cp := *c
		list = append(list, &cp)
	}
	data, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(j.path, data, 0o644)

}

func (j *JSON) Add(c *Contact) error {

	j.mu.Lock()
	defer j.mu.Unlock()
	if _, exists := j.contacts[c.ID]; exists {
		return errors.New("contact déjà existant")
	}
	cp := *c
	j.contacts[c.ID] = &cp
	return j.saveLocked()

}

func (j *JSON) List() ([]*Contact, error) {

	j.mu.RLock()
	defer j.mu.RUnlock()
	out := make([]*Contact, 0, len(j.contacts))
	for _, c := range j.contacts {
		cp := *c
		out = append(out, &cp)
	}
	return out, nil

}

func (j *JSON) Delete(id int) error {

	j.mu.Lock()
	defer j.mu.Unlock()
	if _, ok := j.contacts[id]; !ok {
		return errors.New("contact introuvable")
	}
	delete(j.contacts, id)
	return j.saveLocked()

}

func (j *JSON) Update(c *Contact) error {

	j.mu.Lock()
	defer j.mu.Unlock()
	if _, ok := j.contacts[c.ID]; !ok {
		return errors.New("contact introuvable")
	}
	cp := *c
	j.contacts[c.ID] = &cp
	return j.saveLocked()

}

func (j *JSON) Get(id int) (*Contact, error) {

	j.mu.RLock()
	defer j.mu.RUnlock()
	c, ok := j.contacts[id]
	if !ok {
		return nil, errors.New("contact introuvable")
	}
	cp := *c
	return &cp, nil

}
