package petblog

type Post struct {
	ID          int    `json:"-"`
	UserID      int    `json:"-"` // foreign key
	Username    string `json:"username"`
	PetID       int    `json:"petid"` // foreign key
	PetName     string `json:"petname"`
	Title       string `json:"title"`
	Descritpion string `json:"description"`
}
