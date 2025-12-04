package storage

import (
	"errors"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type GORMStore struct {
	db *gorm.DB
}

func NewGORMStore(dbPath string) (*GORMStore, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto Migrate the schema
	if err := db.AutoMigrate(&Contact{}); err != nil {
		return nil, err
	}

	return &GORMStore{db: db}, nil
}

func (s *GORMStore) Add(c *Contact) error {
	result := s.db.Create(c)
	return result.Error
}

func (s *GORMStore) List() ([]*Contact, error) {
	var contacts []*Contact
	result := s.db.Find(&contacts)
	return contacts, result.Error
}

func (s *GORMStore) Delete(id int) error {
	result := s.db.Delete(&Contact{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("contact introuvable")
	}
	return nil
}

func (s *GORMStore) Update(c *Contact) error {
	result := s.db.Model(&Contact{ID: c.ID}).Updates(c)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("contact introuvable")
	}
	return nil
}

func (s *GORMStore) Get(id int) (*Contact, error) {
	var contact Contact
	result := s.db.First(&contact, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("contact introuvable")
		}
		return nil, result.Error
	}
	return &contact, nil
}
