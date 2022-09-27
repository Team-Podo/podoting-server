package admin

type FindPlace struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
