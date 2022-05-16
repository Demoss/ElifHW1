package command

type DeleteAllMessages struct {
	From string `json:"from" binding:"required"`
	To   string `json:"to" binding:"required"`
}
