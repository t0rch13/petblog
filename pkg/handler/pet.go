package handler

import (
	"net/http"
	"petblog"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createPet(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input petblog.Pet
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Pet.CreatePet(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) deletePet(c *gin.Context) {

}

type getAllPetsResponse struct {
	Data []petblog.Pet `json:"data"`
}

func (h *Handler) getAllPets(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	pets, err := h.services.Pet.GetAllPets(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllPetsResponse{
		Data: pets,
	})
}

func (h *Handler) getPetById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	pet, err := h.services.Pet.GetPetById(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, pet)
}

func (h *Handler) updatePet(c *gin.Context) {

}
