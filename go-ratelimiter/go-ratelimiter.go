package main

import (
	"encoding/json"
	"net/http"

	"golang.org/x/time/rate"
)

type RateLimiter struct {
	limiter *rate.Limiter
}

func NewRateLimiter(r rate.Limit, b int) *RateLimiter {
	return &RateLimiter{
		limiter: rate.NewLimiter(r, b),
	}
}

func (rl *RateLimiter) Limit(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !rl.limiter.Allow() {
			message := Message{
				Status: "Failed",
				Body:   "Rate limited request",
			}

			w.WriteHeader(http.StatusTooManyRequests)
			err := json.NewEncoder(w).Encode(&message)
			if err != nil {
				return
			}
			return
		}

		next(w, r)
	}
}
