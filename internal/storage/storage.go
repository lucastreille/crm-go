package storage

type Contact struct {
	ID    int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Storage interface {
	Add(c *Contact) error
	List() ([]*Contact, error)
	Delete(id int) error
	Update(c *Contact) error
	Get(id int) (*Contact, error)
}
