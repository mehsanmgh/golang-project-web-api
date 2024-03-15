package middlewares

import (
	"net/http"

	"github.com/didip/tollbooth"
	"github.com/gin-gonic/gin"
	"github.com/mehsanmgh/golang-project-web-api/api/helper"
)

func LimitByRequest() gin.HandlerFunc {

	lmt := tollbooth.NewLimiter(1, nil)

	return func(c *gin.Context) {
		err := tollbooth.LimitByRequest(lmt, c.Writer, c.Request)
		//if err != nil {
		// 	c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
		// 		"error": err.Error(),
		// 	})
		// 	return
		// } else {
		// 	c.Next()
		// }

		if err != nil {
			c.AbortWithStatusJSON(http.StatusTooManyRequests,
				helper.GenerateBaseResponseWithError(nil, false, int(helper.LimiterError), err))
			return
		} else {
			c.Next()
		}
	}
}
