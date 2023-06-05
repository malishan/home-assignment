package model

import "time"

var (
	AppVersion  string
	ServiceName string
)

type HomeConfig struct {
	ActivityCountInResponse int
	RateLimitCount          int
	RateLimitDuration       time.Duration
}

type SchedulerConfig struct {
}
