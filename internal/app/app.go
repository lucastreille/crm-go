package app

import "github.com/lucastreille/crm-go/internal/storage"

type App struct {
	store storage.Storage
}

func New(store storage.Storage) *App {
	return &App{store: store}
}

func (a *App) AddContact(id int, name, email string) error {
	c := &storage.Contact{ID: id, Name: name, Email: email}
	return a.store.Add(c)
}

func (a *App) ListContacts() ([]*storage.Contact, error) {
	return a.store.List()
}

func (a *App) DeleteContact(id int) error {
	return a.store.Delete(id)
}

func (a *App) UpdateContact(id int, name, email string) error {
	c, err := a.store.Get(id)
	if err != nil {
		return err
	}
	if name != "" {
		c.Name = name
	}
	if email != "" {
		c.Email = email
	}
	return a.store.Update(c)
}
