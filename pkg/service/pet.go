package service

import (
	"petblog"
	"petblog/pkg/repository"
)

type PetService struct {
	repo repository.Pet
}

func NewPetService(repo repository.Pet) *PetService {
	return &PetService{repo: repo}
}

func (s *PetService) CreatePet(userId int, pet petblog.Pet) (int, error) {
	return s.repo.CreatePet(userId, pet)
}

func (s *PetService) GetAllPets(userId int) ([]petblog.Pet, error) {
	return s.repo.GetAllPets(userId)
}

func (s *PetService) GetPetById(userId, petId int) (petblog.Pet, error) {
	return s.repo.GetPetById(userId, petId)
}
