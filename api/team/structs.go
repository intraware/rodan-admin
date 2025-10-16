package team

// swagger:model
type TeamResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	Ban       bool   `json:"ban"`
	Blacklist bool   `json:"blacklist"`
	LeaderID  uint   `json:"leader"`
}