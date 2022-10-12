package performance_content_find

import "github.com/Team-Podo/podoting-server/repository"

type PerformanceContent struct {
	ID            uint   `json:"id"`
	ManagingTitle string `json:"managingTitle"`
	Content       string `json:"content"`
	Visible       bool   `json:"visible"`
	CreatedAt     string `json:"createdAt"`
	UpdatedAt     string `json:"updatedAt"`
}

func ParseResponse(content *repository.PerformanceContent) PerformanceContent {
	return PerformanceContent{
		ID:            content.ID,
		ManagingTitle: content.ManagingTitle,
		Content:       content.Content,
		Visible:       content.Visible,
		CreatedAt:     content.CreatedAt.String(),
		UpdatedAt:     content.UpdatedAt.String(),
	}
}
