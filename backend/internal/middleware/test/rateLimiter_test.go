package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	middlewareratelimiter "github.com/unmsmfisi-socialapplication/social_app/internal/middleware"
)

func TestNewRateLimiter(t *testing.T) {
	requests := 10
	interval := time.Minute
	rl := middlewareratelimiter.NewRateLimiter(requests, interval)

	if rl.Requests != requests {
		t.Errorf("Expected Requests to be %d, but got %d", requests, rl.Requests)
	}

	if rl.Interval != interval {
		t.Errorf("Expected Interval to be %v, but got %v", interval, rl.Interval)
	}

	if time.Since(rl.LastTimestamp) > time.Second {
		t.Errorf("Expected LastTimestamp to be within the last second, but got %v", rl.LastTimestamp)
	}
}

func TestRateLimiter_Allow(t *testing.T) {
	rl := middlewareratelimiter.NewRateLimiter(2, time.Second)

	// First request should be allowed
	if !rl.Allow() {
		t.Error("Expected first request to be allowed, but it was denied")
	}

	// Second request should be allowed
	if !rl.Allow() {
		t.Error("Expected second request to be allowed, but it was denied")
	}

	// Third request should be denied
	if rl.Allow() {
		t.Error("Expected third request to be denied, but it was allowed")
	}

	// Wait for the interval to pass
	time.Sleep(time.Second * 2)

	// Fourth request should be allowed
	if !rl.Allow() {
		t.Error("Expected fourth request to be allowed, but it was denied")
	}
}

func TestRateLimiter_Handle(t *testing.T) {
	// Create a new rate limiter with a limit of 2 requests per second
	rl := middlewareratelimiter.NewRateLimiter(2, time.Second)

	// Create a new test server with the rate limiter middleware
	ts := httptest.NewServer(rl.Handle(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})))
	defer ts.Close()

	// Create a new request to the test server
	req, err := http.NewRequest(http.MethodGet, ts.URL, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Send the first request, should be allowed
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected first request to return status code %d, but got %d", http.StatusOK, res.StatusCode)
	}

	// Send the second request, should be allowed
	res, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected second request to return status code %d, but got %d", http.StatusOK, res.StatusCode)
	}

	// Send the third request, should be denied
	res, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusTooManyRequests {
		t.Errorf("Expected third request to return status code %d, but got %d", http.StatusTooManyRequests, res.StatusCode)
	}

	// Wait for the interval to pass
	time.Sleep(time.Second * 2)

	// Send the fourth request, should be allowed
	res, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected fourth request to return status code %d, but got %d", http.StatusOK, res.StatusCode)
	}
}
