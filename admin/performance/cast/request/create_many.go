package request

type CreateManyCastRequest struct {
	ID          uint `json:"id"`
	PersonID    uint `json:"personID" binding:"required"`
	CharacterID uint `json:"characterID" binding:"required"`
}
