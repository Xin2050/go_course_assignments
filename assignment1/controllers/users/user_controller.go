package users_controller

import (
	users_services "github.com/Xin2050/go_course_assignments/s1/services"
	"github.com/Xin2050/go_course_assignments/s1/utils/errors"
	"github.com/gin-gonic/gin"

	"net/http"
	"strconv"
)

func GetUserId(userIdParam string) (int64, *errors.RestError) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return -1, errors.BadRequestError("user id should be a number")
	}
	return userId, nil
}
func GetUser(c *gin.Context) {
	userId, userErr := GetUserId(c.Param("user_id"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}
	user, getErr := users_services.UsersService.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, &user)
}
