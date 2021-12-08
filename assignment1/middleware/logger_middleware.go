package middleware

import (
	"github.com/Xin2050/go_course_assignments/s1/logger"
	rest_errors "github.com/Xin2050/go_course_assignments/s1/utils/errors"
	"github.com/gin-gonic/gin"
)

func GinLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
		// if use c.Error(err) set error to context, then use err := c.Errors.Last() get error back,
		// then error gonna to change to a different object, not showing the correct stack trace info
		// so i use c.Get and c.Set
		var err error
		data, ok := c.Get("Errors")
		if !ok {
			return
		}
		err = data.(error)
		if err != nil {
			// write log for debugger
			logger.Errors(err)
			// get json response for client
			restError := rest_errors.ParseError(err)
			c.JSON(restError.Status, &restError)
		}

	}
}
