package utils

import (
	"BE_Hospital_Management/constant"
	"BE_Hospital_Management/pkg"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAuthUserId(c *gin.Context) *int64 {
	rawAuthUserId, exists := c.Get("authUserId")
	if !exists {
		return nil
	}
	authUserIdStr := fmt.Sprint(rawAuthUserId)
	authUserId, err := strconv.ParseInt(authUserIdStr, 10, 64)
	if err != nil {
		log.Error("Happened error when getting AuthUserId. Error: ", err)
		pkg.PanicExeption(constant.UnknownError, "Happened error when validating request.")
	}
	return &authUserId
}

func GetAuthUserRole(c *gin.Context) *string {
	rawAuthUserRole, exists := c.Get("authUserRole")
	if !exists {
		return nil
	}
	authUserRole := fmt.Sprint(rawAuthUserRole)
	return &authUserRole
}
