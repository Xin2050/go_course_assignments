package app

import (
	users_controller "github.com/Xin2050/go_course_assignments/s1/controllers/users"
	"github.com/Xin2050/go_course_assignments/s1/logger"
	"github.com/Xin2050/go_course_assignments/s1/middleware"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {

	router.Use(middleware.GinLoggerMiddleware())
	logger.Info("application is trying to run on port 3000.")
	router.SetTrustedProxies([]string{"localhost"})
	// only for error
	router.GET("/user/:user_id", users_controller.GetUser)
	router.POST("/user", users_controller.SaveUser)

	err := router.Run(":3000")
	if err != nil {
		logger.Error("application was failed to start", err)
	}

}
