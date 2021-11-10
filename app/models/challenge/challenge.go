package challenge

type Challenge struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	StartsAt string `json:"starts_at"`
	EndsAt   string `json:"ends_at"`
	IsActive string `json:"is_active"`
	Type     string `json:"type"`
}

func (Challenge) TableName() string {
	return "challenge_challenges"
}

func (Challenge) Filterable() []string {
	return []string{
		"id",
		"is_active",
		"type",
	}
}
