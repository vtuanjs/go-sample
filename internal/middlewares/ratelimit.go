package middlewares

import (
	"fmt"
	"log"
	"time"
	"vtuanjs/my-project/global"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	redisStore "github.com/ulule/limiter/v3/drivers/store/redis"
)

type RateLimiter struct {
	globalRateLimiter         *limiter.Limiter
	publicAPIRateLimiter      *limiter.Limiter
	userPrivateAPIRateLimiter *limiter.Limiter
}

func NewRateLimiter() *RateLimiter {
	rateLimit := &RateLimiter{
		globalRateLimiter:         rateLimiter("100-S"),
		publicAPIRateLimiter:      rateLimiter("80-S"),
		userPrivateAPIRateLimiter: rateLimiter("50-S"),
	}
	return rateLimit
}

func rateLimiter(interval string) *limiter.Limiter {
	store, err := redisStore.NewStoreWithOptions(global.RDB, limiter.StoreOptions{
		Prefix:          "rate-limiter",
		MaxRetry:        3,
		CleanUpInterval: time.Hour,
	})
	if err != nil {
		return nil
	}
	rate, err := limiter.NewRateFromFormatted(interval)
	if err != nil {
		panic(err)
	}
	instance := limiter.New(store, rate)
	return instance
}

// GLOBAL
func (rl *RateLimiter) GlobalRateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := "global"
		log.Printf("key: Global --->")
		limitContext, err := rl.globalRateLimiter.Get(c, key)
		if err != nil {
			fmt.Println("Failed to check rate limit GLOBAL", err)
			c.Next()
			return
		}
		if limitContext.Reached {
			log.Printf("Rate limit reached for %s", key)
			c.AbortWithStatusJSON(429, gin.H{"error": "Too many requests"})
			return
		}
		c.Next()
	}
}

func (rl *RateLimiter) PublicAPIRateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("key: Public --->")
		urlPath := c.Request.URL.Path
		ratelimitPath := rl.filterLimitUrlPath(urlPath)
		if ratelimitPath == nil {
			log.Println("Client IP: ", c.ClientIP())

			key := fmt.Sprintf("%s:%s", c.ClientIP(), urlPath)
			limitContext, err := ratelimitPath.Get(c, key)
			if err != nil {
				fmt.Println("Failed to check rate limit", err)
				c.Next()
				return
			}
			if limitContext.Reached {
				log.Printf("Rate limit reached for %s", key)
				c.AbortWithStatusJSON(429, gin.H{"error": "Too many requests"})
				return
			}
		}

		c.Next()
	}
}

func (rl *RateLimiter) UserAndPrivateAPIRateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("key: User And Private --->")
		urlPath := c.Request.URL.Path
		ratelimitPath := rl.filterLimitUrlPath(urlPath)
		if ratelimitPath == nil {
			userId := "11111111"
			key := fmt.Sprintf("%s:%s", userId, urlPath)
			limitContext, err := ratelimitPath.Get(c, key)
			if err != nil {
				fmt.Println("Failed to check rate limit", err)
				c.Next()
				return
			}
			if limitContext.Reached {
				log.Printf("Rate limit reached for %s", key)
				c.AbortWithStatusJSON(429, gin.H{"error": "Too many requests"})
				return
			}
		}

		c.Next()
	}
}

func (rl *RateLimiter) filterLimitUrlPath(urlPath string) *limiter.Limiter {
	if urlPath == "/api/v1/users/login" {
		return rl.publicAPIRateLimiter
	} else if urlPath == "/api/v1/users/test-private" {
		return rl.userPrivateAPIRateLimiter
	}
	return rl.globalRateLimiter
}
