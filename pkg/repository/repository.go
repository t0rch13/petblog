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
	CreatePost(userId int, petId int, post petblog.Post) (int, error)
	GetAllPosts() ([]petblog.Post, error)
	GetPostById(postId int) (petblog.Post, error)
	DeletePost(userId, postId int) error
	UpdatePost(userId, postId int, input petblog.UpdatePostInput) error
	GetUserPosts(userId int) ([]petblog.Post, error)
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
		Post:         NewPostPostgres(db),
	}
}
