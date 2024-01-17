package mw

import (
	"strconv"

	siogo "gitea.slauson.io/slausonio/siogo"
	"github.com/gin-gonic/gin"
)

// IntQueryMiddleware function  
// Converts integer query params to ints
// and stores them in the gin context.
func IntQueryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		intParams := make(map[string]int)

		for k, v := range c.Request.URL.Query() {
			if i, err := strconv.Atoi(v[0]); err == nil {
				intParams[k] = i
			} else {
				_ = c.AbortWithError(400, siogo.NewBadRequestError(siogo.INVALID_ID))
			}
		}

		c.Set("intParams", intParams)
		c.Next()
	}
}
