package controllers

import (
	"MATHB/app/apiException"
	authservice "MATHB/app/services/authService"
	userservice "MATHB/app/services/userService"
	"MATHB/app/utils"
	"MATHB/app/utils/env"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	Password string `json:"password"`
	Account  string `json:"account"`
}

func LoginByIDHandler(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithError(400, apiException.ParamError)
		return
	}

	account, err := strconv.ParseUint(req.Account, 0, 0)
	if err != nil {
		c.AbortWithError(400, apiException.ParamError)
		return
	}

	user, err := userservice.GetUserByIDAndPass(uint(account), req.Password)
	if err != nil {
		c.AbortWithError(400, apiException.NoThatUserOrPasswordWrong)
		return
	}

	token, err := authservice.CreateJWT(user.ID, time.Duration(env.TokenDuration))

	if err != nil {
		c.AbortWithError(500, apiException.ServerError)
		return
	}

	utils.JsonSuccessResponse(c, gin.H{
		"user":  user,
		"token": token,
	})
}
