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

	err := h.services.Sender.Send(c, dialog)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "message successfully sent")
}

func (h *Handler) DeleteAllMessages(c *gin.Context) {
	from := c.Param("sender")
	to := c.Param("receiver")

	request := command.DeleteAllMessages{
		From: from,
		To:   to,
	}
	err := h.services.Sender.DeleteAllMessages(c, request)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "messages deleted")
}

func (h *Handler) UpdateMessage(c *gin.Context) {
	from := c.Param("sender")
	to := c.Param("receiver")
	message := c.Param("message")

	var input command.UpdateMessage

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	dialog := domain.Dialog{
		SendFrom: from,
		SendTo:   to,
		Text:     message,
	}
	err := h.services.Sender.UpdateMessage(c, dialog, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "message updated")
}
