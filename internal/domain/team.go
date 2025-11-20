package domain

type Team struct {
	ID   int
	Name string
}

type TeamRepository interface {
	GetById(id int) (*Team, error)
	GetByName(name string) (*Team, error)
	Create(team *Team) (int, error)
	Update(team *Team) error
	Delete(id int) error
}
