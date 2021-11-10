package challenge

type Question struct {
	ID       uint   `json:"id"`
	Text     string `json:"text"`
	Level    string `json:"level"`
	Hint     string `json:"hint"`
	Duration string `json:"duration"`
	Type     string `json:"type"`
}
