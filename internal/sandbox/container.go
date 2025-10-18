package sandbox

import (
	"context"
	"time"
)

type container struct {
	Context     context.Context
	ContainerID string
	ImageName   string
	ChallengeID uint
	TTL         time.Duration
	StartedAt   time.Time
}