package ratelimit

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

// GetRateLimit creates a rate limitter
func GetRateLimit(formatted string) gin.HandlerFunc {
	rate, err := limiter.NewRateFromFormatted(formatted)
	if err != nil {
		log.Fatal(err.Error())
	}
	store := memory.NewStore()
	return mgin.NewMiddleware(limiter.New(store, rate))
}
