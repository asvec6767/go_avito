package domain

type User struct {
	ID       int
	Name     string
	IsActive bool
}

type UserRepository interface {
	GetById(id int) (*User, error)
	GetByName(name string) (*User, error)
	GetList(ids []int) ([]*User, error)
	Create(user *User) (int, error)
	Update(user *User) error
	Delete(id int) error
}
