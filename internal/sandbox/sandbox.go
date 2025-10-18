package sandbox

import (
	"context"
	"time"

	"github.com/intraware/rodan-admin/internal/models"
)

type SandBox struct {
	UserID        uint
	TeamID        uint
	ChallengeMeta models.Challenge
	Container     *container
	CreatedAt     time.Time
	Active        bool
	Flag          string
	Context       context.Context
	CancelFunc    context.CancelFunc
}