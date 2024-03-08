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

	// c.JSON(http.StatusOK, map[string]interface{}{
	// 	"id": id,
	// })
	c.HTML(http.StatusOK, "create.html", gin.H{
		"id": id,
	})
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

	// c.JSON(http.StatusOK, getAllPetsResponse{
	// 	Data: pets,
	// })
	c.HTML(http.StatusOK, "pets.html", gin.H{
		"pets": pets,
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

	// c.JSON(http.StatusOK, pet)
	c.HTML(http.StatusOK, "pet.html", gin.H{
		"pet": pet,
	})
}

func (h *Handler) updatePet(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input petblog.UpdatePetInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Pet.UpdatePet(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})

}

func (h *Handler) deletePet(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Pet.DeletePet(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})

}
