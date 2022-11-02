package controllers

import (
	"gin-boilerplate/helpers"
	"gin-boilerplate/infra/logger"
	"gin-boilerplate/models"
	"gin-boilerplate/repository"
	"gin-boilerplate/util/constants"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) (interface{}, error) {
	var loginVals models.Login
	if err := c.ShouldBind(&loginVals); err != nil {
		return "", jwt.ErrMissingLoginValues
	}

	Email := loginVals.Email
	password := loginVals.Password

	user := models.User{
		Email: Email,
	}

	repository.GetOne(&user)
	repository.Update(&user, models.User{IsLogin: 1})

	if user.Id != 0 {
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
			logger.Errorf("bcrypt.CompareHashAndPassword() error : %v", err)
			return nil, jwt.ErrFailedAuthentication
		}

		return &models.User{
			Email: user.Email,
			Name:  user.Name,
		}, nil
	}

	return nil, jwt.ErrFailedAuthentication
}

func Authorizator(data interface{}, c *gin.Context) bool {
	if v, ok := data.(*models.User); ok {
		user := models.User{
			Email: v.Email,
		}

		repository.GetOne(&user)

		return user.IsLogin == 1
	}

	return false
}

func (base *Controller) Logout(ctx *gin.Context) {
	response := helpers.Response{}

	user := new(models.User)

	calimUser, _ := ctx.Get(constants.IdentityKey)
	user.Email = calimUser.(*models.User).Email

	repository.GetOne(&user)
	repository.Update(&user, models.User{IsLogin: 2})

	user.Password = ""
	response.Data = &user

	ctx.JSON(http.StatusOK, response.Success())
}

func (base *Controller) UserProfile(ctx *gin.Context) {
	response := helpers.Response{}

	user := new(models.User)

	calimUser, _ := ctx.Get(constants.IdentityKey)
	user.Email = calimUser.(*models.User).Email

	repository.GetOne(&user)

	user.Password = ""
	response.Data = &user

	ctx.JSON(http.StatusOK, response.Success())
}
