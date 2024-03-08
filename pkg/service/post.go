package service

import (
	"petblog"
	"petblog/pkg/repository"
)

type PostService struct {
	repo repository.Post
}

func NewPostService(repo repository.Post) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) CreatePost(userId int, petId int, post petblog.Post) (int, error) {
	return s.repo.CreatePost(userId, petId, post)
}

func (s *PostService) GetAllPosts() ([]petblog.Post, error) {
	return s.repo.GetAllPosts()
}

func (s *PostService) GetPostById(postId int) (petblog.Post, error) {
	return s.repo.GetPostById(postId)
}

func (s *PostService) DeletePost(userId, postId int) error {
	return s.repo.DeletePost(userId, postId)
}

func (s *PostService) UpdatePost(userId, postId int, input petblog.UpdatePostInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.UpdatePost(userId, postId, input)
}

func (s *PostService) GetUserPosts(userId int) ([]petblog.Post, error) {
	return s.repo.GetUserPosts(userId)
}
