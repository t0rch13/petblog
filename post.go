package petblog

import "errors"

type Post struct {
	ID       int    `json:"-" db:"id"`
	UserID   int    `json:"-" db:"user_id"` // foreign key
	Username string `json:"username" db:"username"`
	PetID    int    `json:"-" db:"pet_id"` // foreign key
	PetName  string `json:"petname" db:"petname"`
	Title    string `json:"title" binding:"required" db:"title"`
	Content  string `json:"content" binding:"required" db:"content"`
}

type UpdatePostInput struct {
	Title   *string `json:"title"`
	Content *string `json:"content"`
}

func (i *UpdatePostInput) Validate() error {
	if i.Title == nil && i.Content == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
