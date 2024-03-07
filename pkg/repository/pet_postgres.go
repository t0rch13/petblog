package repository

import (
	"fmt"
	"petblog"

	"github.com/jmoiron/sqlx"
)

type PetPostgres struct {
	db *sqlx.DB
}

func NewPetPostgres(db *sqlx.DB) *PetPostgres {
	return &PetPostgres{db: db}
}

func (r *PetPostgres) CreatePet(userId int, pet petblog.Pet) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (user_id, petname, species, breed, age) values ($1, $2, $3, $4, $5) RETURNING id", petsTable)
	row := r.db.QueryRow(query, userId, pet.PetName, pet.Species, pet.Breed, pet.Age)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *PetPostgres) GetAllPets(userId int) ([]petblog.Pet, error) {
	var pets []petblog.Pet
	query := fmt.Sprintf("SELECT id, petname FROM %s WHERE user_id=$1", petsTable)
	if err := r.db.Select(&pets, query, userId); err != nil {
		return nil, err
	}

	return pets, nil
}

func (r *PetPostgres) GetPetById(userId, petId int) (petblog.Pet, error) {
	var pet petblog.Pet
	query := fmt.Sprintf("SELECT id, petname, species, breed, age FROM %s WHERE user_id=$1 AND id=$2", petsTable)
	if err := r.db.Get(&pet, query, userId, petId); err != nil {
		return pet, err
	}

	return pet, nil
}
