package handler

import (
	"ElifHW1/internal/command"
	"ElifHW1/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Send(c *gin.Context) {
	from := c.Param("sender")
	to := c.Param("receiver")

	var input command.Send
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	dialog := domain.Dialog{
		SendFrom: from,
		SendTo:   to,
		Text:     input.Text,
	}

	err := h.repos.Sender.Send(c, dialog)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "message successfully sent")
}
