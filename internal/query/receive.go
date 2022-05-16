package query

type Receive struct {
	To   string `json:"to" binding:"required"`
	From string `json:"from" binding:"required"`
}
