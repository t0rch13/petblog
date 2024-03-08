package repository

import (
	"fmt"
	"petblog"
	"strings"

	"github.com/jmoiron/sqlx"
)

type PostPostgres struct {
	db *sqlx.DB
}

func NewPostPostgres(db *sqlx.DB) *PostPostgres {
	return &PostPostgres{db: db}
}

func (r *PostPostgres) CreatePost(userId int, petId int, post petblog.Post) (int, error) {
	var id int
	var username string
	var petname string

	userQuery := `SELECT username FROM users WHERE id = $1`
	err := r.db.QueryRow(userQuery, userId).Scan(&username)
	if err != nil {
		return 0, err
	}

	petQuery := `SELECT petname FROM pets WHERE id = $1`
	err = r.db.QueryRow(petQuery, petId).Scan(&petname)
	if err != nil {
		return 0, err
	}

	query := fmt.Sprintf("INSERT INTO %s (user_id, username, pet_id, petname, title, content) values ($1, $2, $3, $4, $5, $6) RETURNING id", postsTable)
	row := r.db.QueryRow(query, userId, username, petId, petname, post.Title, post.Content)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *PostPostgres) GetAllPosts() ([]petblog.Post, error) {
	var posts []petblog.Post
	query := fmt.Sprintf("SELECT id, user_id, username, pet_id, petname, title, content FROM %s", postsTable)
	err := r.db.Select(&posts, query)
	return posts, err
}

func (r *PostPostgres) GetPostById(userId, postId int) (petblog.Post, error) {
	var post petblog.Post
	query := fmt.Sprintf("SELECT id, user_id, username, pet_id, petname, title, content FROM %s WHERE id=$1 and user_id=$2", postsTable)
	err := r.db.Get(&post, query, postId, userId)
	return post, err
}

func (r *PostPostgres) DeletePost(userId, postId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1 AND user_id=$2", postsTable)
	_, err := r.db.Exec(query, postId, userId)
	return err
}

func (r *PostPostgres) UpdatePost(userId, postId int, input petblog.UpdatePostInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Content != nil {
		setValues = append(setValues, fmt.Sprintf("content=$%d", argId))
		args = append(args, *input.Content)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d AND user_id=$%d", postsTable, setQuery, argId, argId+1)
	args = append(args, postId, userId)
	_, err := r.db.Exec(query, args...)
	return err
}

func (r *PostPostgres) GetUserPosts(userId int) ([]petblog.Post, error) {
	var posts []petblog.Post
	query := fmt.Sprintf("SELECT id, user_id, username, pet_id, petname, title, content FROM %s WHERE user_id=$1", postsTable)
	err := r.db.Select(&posts, query, userId)
	return posts, err
}
