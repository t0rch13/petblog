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
	}
}
