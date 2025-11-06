package app

import (
	"fmt"

	"github.com/lucastreille/crm-go/internal/notification"
	"github.com/lucastreille/crm-go/internal/storage"
)

type App struct {
	store     storage.Storage
	notifiers []notification.Notifier
}

func New(store storage.Storage, notifiers []notification.Notifier) *App {
	return &App{store: store, notifiers: notifiers}
}

func (a *App) AddContact(id int, name, email string) error {
	c := &storage.Contact{ID: id, Name: name, Email: email}
	if err := a.store.Add(c); err != nil {
		return err
	}

	msg := fmt.Sprintf("Nouveau contact ajouté : %s (%s)", c.Name, c.Email)
	notification.NotifyAll(a.notifiers, msg)
	return nil
}

func (a *App) ListContacts() ([]*storage.Contact, error) {
	return a.store.List()
}

func (a *App) DeleteContact(id int) error {
	if err := a.store.Delete(id); err != nil {
		return err
	}
	msg := fmt.Sprintf("Contact supprimé (ID: %d)", id)
	notification.NotifyAll(a.notifiers, msg)
	return nil
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

	if err := a.store.Update(c); err != nil {
		return err
	}

	msg := fmt.Sprintf("Contact mis à jour : %s (%s)", c.Name, c.Email)
	notification.NotifyAll(a.notifiers, msg)
	return nil
}
