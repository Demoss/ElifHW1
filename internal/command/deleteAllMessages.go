package command

type DeleteAllMessages struct {
	From string `json:"from"`
	To   string `json:"to"`
}
