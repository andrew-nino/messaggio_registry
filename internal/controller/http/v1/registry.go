package v1

import (
	"net/http"
	"strconv"

	"github.com/andrew-nino/messaggio/internal/domain/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) addClient(c *gin.Context) {

	client := models.Client{}

	if err := c.BindJSON(&client); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.RegisterClient(client)
	if err != nil {
		h.log.Printf("error RegisterClient: %v", err)
		newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Client added successfully",
		"id":      id,
	})
}

func (h *Handler) updateClient(c *gin.Context) {

	client := models.Client{}

	if err := c.ShouldBindJSON(&client); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err := h.services.UpdateClient(client)
	if err != nil {
		h.log.Printf("error UpdateClient: %v", err)
		newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Client updated successfully",
	})
}

func (h *Handler) deleteClient(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid client ID")
		return
	}

	err = h.services.DeleteClient(id)
	if err != nil {
		h.log.Printf("error DeleteClient: %v", err)
		newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Client deleted successfully",
	})
}

func (h *Handler) getStatistic(c *gin.Context) {

	statistic, err := h.services.GetStatistic()
	if err != nil {
		h.log.Printf("error getStatistic: %v", err)
		newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, statistic)
}
