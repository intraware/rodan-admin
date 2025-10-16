package user

// swagger:model
type UserResponse struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatar_url"`
	Active    bool   `json:"active"`
	Ban       bool   `json:"ban"`
	Blacklist bool   `json:"blacklist"`
	TeamID    *uint  `json:"team_id,omitempty"`
}