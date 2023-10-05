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

	if !rl.Allow() {
		t.Error("Expected first request to be allowed, but it was denied")
	}

	if !rl.Allow() {
		t.Error("Expected second request to be allowed, but it was denied")
	}

	if rl.Allow() {
		t.Error("Expected third request to be denied, but it was allowed")
	}

	time.Sleep(time.Second * 2)

	if !rl.Allow() {
		t.Error("Expected fourth request to be allowed, but it was denied")
	}
}

func TestRateLimiter_Handle(t *testing.T) {
	rl := middlewareratelimiter.NewRateLimiter(2, time.Second)

	ts := httptest.NewServer(rl.Handle(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})))
	defer ts.Close()

	req, err := http.NewRequest(http.MethodGet, ts.URL, nil)
	if err != nil {
		t.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected first request to return status code %d, but got %d", http.StatusOK, res.StatusCode)
	}

	res, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected second request to return status code %d, but got %d", http.StatusOK, res.StatusCode)
	}

	res, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusTooManyRequests {
		t.Errorf("Expected third request to return status code %d, but got %d", http.StatusTooManyRequests, res.StatusCode)
	}

	time.Sleep(time.Second * 2)

	res, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected fourth request to return status code %d, but got %d", http.StatusOK, res.StatusCode)
	}
}
