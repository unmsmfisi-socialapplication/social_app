package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type MockTimeProvider struct {
	now time.Time
}

func (mtp *MockTimeProvider) Now() time.Time {
	return mtp.now
}

func (mtp *MockTimeProvider) Sleep(d time.Duration) {
	mtp.now = mtp.now.Add(d)
}

func newMockTimeProvider() *MockTimeProvider {
	return &MockTimeProvider{
		now: time.Now(),
	}
}

func TestNewRateLimiter(t *testing.T) {
	maxRequests := 10
	interval := time.Minute
	mockTimeProvider := newMockTimeProvider()

	rl := NewRateLimiter(maxRequests, interval, mockTimeProvider)

	if rl.MaxRequests != maxRequests {
		t.Errorf("Expected Requests to be %d, but got %d", maxRequests, rl.MaxRequests)
	}

	if rl.Interval != interval {
		t.Errorf("Expected Interval to be %v, but got %v", interval, rl.Interval)
	}

	mockTimeProvider.Sleep(1 * time.Minute)

	if time.Since(rl.LastTimestamp) > time.Second {
		t.Errorf("Expected LastTimestamp to be within the last second, but got %v", rl.LastTimestamp)
	}
}

func TestRateLimiter_Allow(t *testing.T) {
	waitTime := 2

	mockTimeProvider := newMockTimeProvider()

	rl := NewRateLimiter(waitTime, time.Second, mockTimeProvider)

	if !rl.Allow() {
		t.Error("Expected first request to be allowed, but it was denied")
	}

	if !rl.Allow() {
		t.Error("Expected second request to be allowed, but it was denied")
	}

	if rl.Allow() {
		t.Error("Expected third request to be denied, but it was allowed")
	}

	mockTimeProvider.Sleep(time.Second * time.Duration(waitTime))

	if !rl.Allow() {
		t.Error("Expected fourth request to be allowed, but it was denied")
	}
}

func TestRateLimiter_Handle(t *testing.T) {
	mockTimeProvider := newMockTimeProvider()

	waitTime := 2

	rl := NewRateLimiter(waitTime, time.Second, mockTimeProvider)

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

	mockTimeProvider.Sleep(time.Second * time.Duration(waitTime))

	res, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected fourth request to return status code %d, but got %d", http.StatusOK, res.StatusCode)
	}
}
