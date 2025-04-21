package middlewares

import (
	"context"
	"fmt"
	"log"
	"vtuanjs/my-project/internal/utils/auth"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.Request.URL.Path
		log.Println(" url request: ", url)
		jwtToken, err := auth.ExtractBearerToken(c)
		if !err {
			c.AbortWithStatusJSON(401, gin.H{"code": 40001, "err": "Unauthorized"})

		}

		// validate
		// get user from cache/db
		fmt.Println(jwtToken)
		ctx := context.WithValue(c.Request.Context(), "subjectUUID", jwtToken)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
