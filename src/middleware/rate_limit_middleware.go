package middleware

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"koivu/gateway/config"
)

type rateLimiter struct {
	visits      map[string]int
	mu          sync.Mutex
	limit       int
	resetPeriod time.Duration
}

func newRateLimiter(limit int, resetPeriod time.Duration) *rateLimiter {
	rl := &rateLimiter{
		visits:      make(map[string]int),
		limit:       limit,
		resetPeriod: resetPeriod,
	}

	go rl.resetVisits()
	return rl
}

func (rl *rateLimiter) resetVisits() {
	for {
		time.Sleep(rl.resetPeriod)
		rl.mu.Lock()
		rl.visits = make(map[string]int)
		rl.mu.Unlock()
	}
}

func (rl *rateLimiter) allowVisit(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	visits := rl.visits[ip]
	if visits >= rl.limit {
		return false
	}

	rl.visits[ip]++
	return true
}

func getResetPeriod(timeframe config.RateLimitTimeFrame) time.Duration {
	switch timeframe {
	case config.Second:
		return time.Second
	case config.Minute:
		return time.Minute
	case config.Hour:
		return time.Hour
	case config.Day:
		return 24 * time.Hour
	default:
		return time.Minute // default to a minute if unknown
	}
}

func RateLimitMiddleware(rateLimitConfig *config.RateLimitConfiguration, next http.Handler) http.Handler {
	fmt.Println("ratelimitconfig", rateLimitConfig)

	// If rateLimitConfig is nil, provide a default configuration
	if rateLimitConfig == nil {
		rateLimitConfig = &config.RateLimitConfiguration{
			Requests:  100, // Default request limit
			Timeframe: config.Minute,
			Type:      config.IP,
		}
	}

	rl := newRateLimiter(int(rateLimitConfig.Requests), getResetPeriod(rateLimitConfig.Timeframe))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		if !rl.allowVisit(ip) {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
