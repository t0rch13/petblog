package petblog

import "errors"

type Pet struct {
	ID      int    `json:"id" db:"id"`
	UserID  int    `json:"-"` // foreign key
	PetName string `json:"petname" db:"petname" binding:"required"`
	Species string `json:"species" db:"species" binding:"required"`
	Breed   string `json:"breed" db:"breed"`
	Age     int    `json:"age" db:"age"`
}

type UpdatePetInput struct {
	PetName *string `json:"petname"`
	Species *string `json:"species"`
	Breed   *string `json:"breed"`
	Age     *int    `json:"age"`
}

func (i *UpdatePetInput) Validate() error {
	if i.PetName == nil && i.Species == nil && i.Breed == nil && i.Age == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
