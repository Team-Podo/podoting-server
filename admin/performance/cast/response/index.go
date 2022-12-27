package response

import "github.com/Team-Podo/podoting-server/repository"

type IndexPerformanceResponse struct {
	Title string `json:"title"`
}

type IndexCastResponse struct {
	ID           uint    `json:"id"`
	ProfileImage *string `json:"profileImage"`
	CharacterID  *uint   `json:"characterID"`
	PersonID     *uint   `json:"personID"`
	CreatedAt    string  `json:"createdAt"`
	UpdatedAt    string  `json:"updatedAt"`
}

type IndexResponse struct {
	Performance IndexPerformanceResponse `json:"performance"`
	Casts       []IndexCastResponse      `json:"casts"`
}

func GetIndexResponse(casts []repository.Cast, performance *repository.Performance) IndexResponse {
	return IndexResponse{
		Casts:       getIndexCasts(casts),
		Performance: getIndexPerformance(performance),
	}
}

func getIndexPerformance(performance *repository.Performance) IndexPerformanceResponse {
	return IndexPerformanceResponse{
		Title: performance.Title,
	}
}

func getIndexCasts(casts []repository.Cast) []IndexCastResponse {
	castResponses := make([]IndexCastResponse, len(casts))
	for i, cast := range casts {
		castResponses[i] = IndexCastResponse{
			ID:           cast.ID,
			ProfileImage: cast.GetProfileImageUrl(),
			CharacterID:  &cast.CharacterID,
			PersonID:     &cast.PersonID,
			CreatedAt:    cast.CreatedAt.String(),
			UpdatedAt:    cast.UpdatedAt.String(),
		}
	}
	return castResponses
}
