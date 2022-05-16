package domain

type Dialog struct {
	SendFrom string `json:"send_from"`
	SendTo   string `json:"send_to"`
	Text     string `json:"text"`
}
