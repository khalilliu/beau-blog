package middleware

import (
	"beau-blog/model/request"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		 var userId uint

		 if claims, ok := c.Get("claims"); ok {
		 	waitUse := claims.(*request.ReqCustomClaims)
		 	userId = uint(waitUse.ID)
		 } else {
		 	id, err := strconv.ParseUint(c.Request.Header.Get("x-user-id"), 10, 64)
		 	if err != nil {
		 		userId = 0
			}
		 	userId = uint(id)
		 }

		 c.Set("userId", userId)
		 c.Next()
	}
}

