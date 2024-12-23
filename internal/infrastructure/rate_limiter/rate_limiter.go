package rate_limiter

import (
	"golang.org/x/time/rate"
	"sync"
)

type RateLimiter struct {
	limiters map[int]*rate.Limiter
	mu       sync.Mutex
	rate     rate.Limit
	burst    int
}

func NewRateLimiter(r rate.Limit, b int) *RateLimiter {
	return &RateLimiter{
		limiters: make(map[int]*rate.Limiter),
		rate:     r,
		burst:    b,
	}

}

func (rl *RateLimiter) getLimiter(userID int) *rate.Limiter {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	if _, exists := rl.limiters[userID]; !exists {
		rl.limiters[userID] = rate.NewLimiter(rl.rate, rl.burst)
	}

	return rl.limiters[userID]
}

func (rl *RateLimiter) Allow(userID int) bool {
	limiter := rl.getLimiter(userID)
	return limiter.Allow()
}
