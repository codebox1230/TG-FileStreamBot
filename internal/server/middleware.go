package server

import (
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type rateLimiterEntry struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

type ipRateLimiter struct {
	mu       sync.RWMutex
	visitors map[string]*rateLimiterEntry
	rate     rate.Limit
	burst    int
}

func newIPRateLimiter(r rate.Limit, burst int) *ipRateLimiter {
	l := &ipRateLimiter{
		visitors: make(map[string]*rateLimiterEntry),
		rate:     r,
		burst:    burst,
	}
	go l.cleanup()
	return l
}

func (l *ipRateLimiter) cleanup() {
	ticker := time.NewTicker(10 * time.Minute)
	for range ticker.C {
		l.mu.Lock()
		cutoff := time.Now().Add(-30 * time.Minute)
		for ip, entry := range l.visitors {
			if entry.lastSeen.Before(cutoff) {
				delete(l.visitors, ip)
			}
		}
		l.mu.Unlock()
	}
}

func (l *ipRateLimiter) getVisitor(ip string) *rate.Limiter {
	l.mu.Lock()
	defer l.mu.Unlock()
	entry, exists := l.visitors[ip]
	if !exists {
		entry = &rateLimiterEntry{
			limiter: rate.NewLimiter(l.rate, l.burst),
		}
		l.visitors[ip] = entry
	}
	entry.lastSeen = time.Now()
	return entry.limiter
}

var defaultLimiter = newIPRateLimiter(10, 20)

func RateLimit() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ip := ctx.ClientIP()
		if parsed := net.ParseIP(ip); parsed != nil && parsed.IsLoopback() {
			ctx.Next()
			return
		}
		limiter := defaultLimiter.getVisitor(ip)
		if !limiter.Allow() {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"message": "rate limit exceeded",
				"ok":      false,
			})
			return
		}
		ctx.Next()
	}
}

func SecurityHeaders() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("X-Content-Type-Options", "nosniff")
		ctx.Header("X-Frame-Options", "DENY")
		ctx.Header("X-XSS-Protection", "1; mode=block")
		ctx.Header("Referrer-Policy", "strict-origin-when-cross-origin")
		ctx.Next()
	}
}
