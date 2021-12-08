package users_controller

import (
	users_services "github.com/Xin2050/go_course_assignments/s1/services"
	rest_errors "github.com/Xin2050/go_course_assignments/s1/utils/errors"
	"github.com/pkg/errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"strconv"
)

func GetUserId(userIdParam string) (int64, *rest_errors.RestError) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return -1, rest_errors.BadRequestError("user id should be a number")
	}
	return userId, nil
}
func GetUser(c *gin.Context) {
	userId, userErr := GetUserId(c.Param("user_id"))
	if userErr != nil {
		c.JSON(userErr.Status, &userErr)
		return
	}
	user, getErr := users_services.UsersService.GetUser(userId)
	if getErr != nil {
		err := errors.WithMessage(getErr, "user_controller:GetUser")
		//fmt.Printf("%+v\n", err)
		//c.Error(err)
		c.Set("Errors", err)
		return
	}
	c.JSON(http.StatusOK, &user)
}
