package handler

import (
	"net/http"
	"petblog"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createPost(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	petId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input petblog.Post
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Post.CreatePost(userId, petId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

func (h *Handler) getAllPosts(c *gin.Context) {

	posts, err := h.services.Post.GetAllPosts()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"posts": posts,
	})
}

func (h *Handler) getUserPosts(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	posts, err := h.services.Post.GetUserPosts(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"posts": posts,
	})
}

func (h *Handler) deletePost(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Post.DeletePost(userId, postId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) getPostById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	post, err := h.services.Post.GetPostById(userId, postId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"post": post,
	})
}

func (h *Handler) updatePost(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input petblog.UpdatePostInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Post.UpdatePost(userId, postId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
