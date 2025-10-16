package challenge

import "github.com/intraware/rodan-admin/internal/models"

// swagger:model
type ChallengeResponse struct {
	ID            uint                  `json:"id"`
	Name          string                `json:"name"`
	Author        string                `json:"author"`
	Desc          string                `json:"desc"`
	Category      int8                  `json:"category"`
	PointsMin     int                   `json:"points_min"`
	PointsMax     int                   `json:"points_max"`
	Difficulty    int8                  `json:"difficulty"`
	IsStatic      bool                  `json:"is_static"`
	IsVisible     bool                  `json:"is_visible"`
	StaticConfig  *models.StaticConfig  `json:"static_config,omitempty"`
	DynamicConfig *models.DynamicConfig `json:"dynamic_config,omitempty"`
	Hints         []models.Hint         `json:"hints,omitempty"`
}