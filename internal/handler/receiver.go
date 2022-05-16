package handler

import (
	"ElifHW1/internal/domain"
	"ElifHW1/internal/query"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) ReceiveMessages(c *gin.Context) {
	from := c.Param("sender")
	to := c.Param("receiver")

	res := query.Receive{
		To:   to,
		From: from,
	}
	messages, err := h.repos.Receiver.ReceiveLast(c, res)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	text := getTextMessages(messages)
	c.JSON(http.StatusOK, map[string]interface{}{
		"messages: ": text,
		"you have ":  len(messages),
	})
}

func getTextMessages(mess []domain.Dialog) []string {
	var res []string

	for _, v := range mess {
		res = append(res, v.Text)
	}
	return res
}
