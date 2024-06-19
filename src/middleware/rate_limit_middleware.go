package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"koivu/gateway/config"

	"github.com/redis/go-redis/v9"
)

type rateLimiter struct {
	client      *redis.Client
	limit       int
	resetPeriod time.Duration
}

func newRateLimiter(redisAddr string, redisPassword string, limit int, resetPeriod time.Duration) *rateLimiter {
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       0, // use default DB
	})

	return &rateLimiter{
		client:      client,
		limit:       limit,
		resetPeriod: resetPeriod,
	}
}

func (rl *rateLimiter) allowVisit(ctx context.Context, ip string) (bool, error) {
	key := fmt.Sprintf("rate_limit:%s", ip)

	// Increment the visit count
	visits, err := rl.client.Incr(ctx, key).Result()
	if err != nil {
		return false, err
	}

	// Set the expiration if this is the first visit
	if visits == 1 {
		_, err := rl.client.Expire(ctx, key, rl.resetPeriod).Result()
		if err != nil {
			return false, err
		}
	}

	if visits > int64(rl.limit) {
		return false, nil
	}

	return true, nil
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

func RateLimitMiddleware(redisAddr string, redisPassword string, rateLimitConfig *config.RateLimitConfiguration, next http.Handler) http.Handler {
	if rateLimitConfig == nil {
		return next
	}

	rl := newRateLimiter(redisAddr, redisPassword, int(rateLimitConfig.Requests), getResetPeriod(rateLimitConfig.Timeframe))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("in here")
		ctx := r.Context()
		ip := r.RemoteAddr

		allowed, err := rl.allowVisit(ctx, ip)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		if !allowed {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
