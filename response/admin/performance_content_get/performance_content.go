package performance_content_get

import "github.com/Team-Podo/podoting-server/repository"

type PerformanceContent struct {
	ID            uint   `json:"id"`
	ManagingTitle string `json:"managingTitle"`
	Content       string `json:"content"`
	Visible       bool   `json:"visible"`
	CreatedAt     string `json:"createdAt"`
	UpdatedAt     string `json:"updatedAt"`
}

func ParseResponse(contents []repository.PerformanceContent) []PerformanceContent {
	var response []PerformanceContent

	for _, content := range contents {
		response = append(response, PerformanceContent{
			ID:            content.ID,
			ManagingTitle: content.ManagingTitle,
			Content:       content.Content,
			Visible:       content.Visible,
			CreatedAt:     content.CreatedAt.String(),
			UpdatedAt:     content.UpdatedAt.String(),
		})
	}

	return response
}
