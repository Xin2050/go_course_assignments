package middleware

import (
	"github.com/Xin2050/go_course_assignments/s1/logger"
	rest_errors "github.com/Xin2050/go_course_assignments/s1/utils/errors"
	"github.com/gin-gonic/gin"
)

func GinLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
		err := c.Errors.Last()
		if err != nil {
			// write log for debugger
			logger.Errors(err)
			// get json response for client
			restError := rest_errors.ParseError(err)
			c.JSON(restError.Status, &restError)
		}

	}
}
