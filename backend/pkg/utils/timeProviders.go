package utils

import "time"

type TimeProvider interface {
	Now() time.Time
	Sleep(time.Duration)
}

type RealTimeProvider struct{}

func (rtp *RealTimeProvider) Now() time.Time {
	return time.Now()
}

func (rtp *RealTimeProvider) Sleep(d time.Duration) {
	time.Sleep(d)
}

func NewRealTimeProvider() *RealTimeProvider {
	return &RealTimeProvider{}
}
