package _1_non_idempotent

import (
	"time"
)

func TimeRemaining(when time.Time) time.Duration {
	return when.Sub(time.Now())
}
