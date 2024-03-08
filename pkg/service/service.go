package service

import (
	"petblog"
	"petblog/pkg/repository"
)

type Authorization interface {
	CreateUser(user petblog.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
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

type Service struct {
	Authorization
	Pet
	Post
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Autorization),
		Pet:           NewPetService(repos.Pet),
		Post:          NewPostService(repos.Post),
	}
}
