package repository

import (
	"fmt"
	"petblog"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
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
	query := fmt.Sprintf("SELECT id, petname, species FROM %s WHERE user_id=$1", petsTable)
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

func (r *PetPostgres) DeletePet(userId, petId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id=$1 AND id=$2", petsTable)
	_, err := r.db.Exec(query, userId, petId)
	return err
}

func (r *PetPostgres) UpdatePet(userId, petId int, input petblog.UpdatePetInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.PetName != nil {
		setValues = append(setValues, fmt.Sprintf("petname=$%d", argId))
		args = append(args, *input.PetName)
		argId++
	}
	if input.Species != nil {
		setValues = append(setValues, fmt.Sprintf("species=$%d", argId))
		args = append(args, *input.Species)
		argId++
	}
	if input.Breed != nil {
		setValues = append(setValues, fmt.Sprintf("breed=$%d", argId))
		args = append(args, *input.Breed)
		argId++
	}
	if input.Age != nil {
		setValues = append(setValues, fmt.Sprintf("age=$%d", argId))
		args = append(args, *input.Age)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE user_id=$%d AND id=$%d", petsTable, setQuery, argId, argId+1)

	args = append(args, userId, petId)

	logrus.Debugf("query: %s", query)
	logrus.Debugf("args: %v", args)

	_, err := r.db.Exec(query, args...)
	return err
}
