package petblog

type Pet struct {
	ID      int    `json:"-" db:"id"`
	UserID  int    `json:"-"` // foreign key
	PetName string `json:"petname" db:"petname" binding:"required"`
	Species string `json:"species" db:"species" binding:"required"`
	Breed   string `json:"breed" db:"breed"`
	Age     int    `json:"age" db:"age"`
}
