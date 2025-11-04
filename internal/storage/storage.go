package storage

type Contact struct {
	ID    int
	Name  string
	Email string
}

type Storage interface {
	Add(c *Contact) error
	List() ([]*Contact, error)
	Delete(id int) error
	Update(c *Contact) error
	Get(id int) (*Contact, error)
}
