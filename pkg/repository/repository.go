package repository

import (
	"petblog"

	"github.com/jmoiron/sqlx"
)

type Autorization interface {
	CreateUser(user petblog.User) (int, error)
	GetUser(username, password string) (petblog.User, error)
}

type Pet interface {
	CreatePet(userId int, pet petblog.Pet) (int, error)
	GetAllPets(userId int) ([]petblog.Pet, error)
	GetPetById(userId, petId int) (petblog.Pet, error)
	DeletePet(userId, petId int) error
	UpdatePet(userId, petId int, input petblog.UpdatePetInput) error
}

type Post interface {
}

type Repository struct {
	Autorization
	Pet
	Post
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Autorization: NewAuthPostgres(db),
		Pet:          NewPetPostgres(db),
	}
}
