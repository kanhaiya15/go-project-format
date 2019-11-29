package middlewares

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kanhaiya15/gopf/constants"
	"github.com/kanhaiya15/gopf/lib/logging/gopflogrus"
)

var logger = gopflogrus.NewLogger()

//Authenticate to accessKey
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		if accessKey := c.Request.Header.Get("Authorization"); accessKey != "" {
			c.Next()
		} else {
			c.Header("Status", constants.Status)
			c.AbortWithError(http.StatusUnauthorized, errors.New("401 Unauthorized"))
			c.Writer.Write([]byte(constants.AuthFailed))
		}
	}
}
