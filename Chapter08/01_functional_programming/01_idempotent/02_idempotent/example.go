package _1_non_idempotent

import (
	"time"
)

func TimeRemaining(when, now time.Time) time.Duration {
	return when.Sub(now)
}
