package middlewareratelimiter

import (
	"net/http"
	"sync"
	"time"
)

type RateLimiter struct {
	Requests      int
	Interval      time.Duration
	LastTimestamp time.Time
	Mutex         sync.Mutex
}

func NewRateLimiter(requests int, interval time.Duration) *RateLimiter {
	return &RateLimiter{
		Requests:      requests,
		Interval:      interval,
		LastTimestamp: time.Now(),
	}
}

func (rl *RateLimiter) Allow() bool {
	rl.Mutex.Lock()
	defer rl.Mutex.Unlock()

	now := time.Now()

	if now.Sub(rl.LastTimestamp) >= rl.Interval {
		rl.LastTimestamp = now
		return true
	}

	if rl.Requests > 0 {
		rl.Requests--
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
