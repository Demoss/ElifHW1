package command

type UpdateMessage struct {
	Text string `json:"text" binding:"required"`
}
