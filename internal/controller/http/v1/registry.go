package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type signInInput struct {
	ManagerName string `json:"managername" binding:"required" example:"Manager"`
	Password    string `json:"password" binding:"required" example:"qwerty"`
}

func (h *Handler) addClient(c *gin.Context) {

	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err := h.services.RegisterClient()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Handler addClient successfully"})
}
