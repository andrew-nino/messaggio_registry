package v1

import (
	"net/http"

	"github.com/andrew-nino/messaggio/internal/domain/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) updateStatus(c *gin.Context) {

	answer := models.Answer{}

	if err := c.ShouldBindJSON(&answer); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err := h.approval.Approve(answer)
	if err != nil {
		h.log.Infof("Failed to update status: %v", err)
		newErrorResponse(c, http.StatusInternalServerError, "failed to approve client")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
