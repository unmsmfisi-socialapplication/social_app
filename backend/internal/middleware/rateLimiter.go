package middleware

import (
	"net/http"
	"sync"
	"time"
)

type RateLimiter struct {
	MaxRequests   int
	Interval      time.Duration
	LastTimestamp time.Time
	Mutex         sync.Mutex
	requests      int
}

func NewRateLimiter(maxRequests int, interval time.Duration) *RateLimiter {
	return &RateLimiter{
		MaxRequests:   maxRequests,
		Interval:      interval,
		LastTimestamp: time.Now(),
		requests:      0,
	}
}

func (rl *RateLimiter) Allow() bool {
	rl.Mutex.Lock()
	defer rl.Mutex.Unlock()

	now := time.Now()

	if now.Sub(rl.LastTimestamp) >= rl.Interval {
		rl.LastTimestamp = now
		rl.requests = 0
		return true
	}

	if rl.requests < rl.MaxRequests {
		rl.requests++
		return true
	}

	return false
}

func (rl *RateLimiter) Handle(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !rl.Allow() {
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
