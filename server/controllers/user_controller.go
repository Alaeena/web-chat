package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/models"
	"server/services"
	"server/utils"
)

type UserController struct {
	service services.UserService
}

func GetUser(service services.UserService) UserController {
	return UserController{service: service}
}

func (c *UserController) Register(con *gin.Context) {
	var payload models.RegisterReq
	err := con.ShouldBindJSON(&payload)
	if err != nil {
		msg := fmt.Errorf("couldn't bind json: %s", err)
		utils.RespondError(con, 400, msg)
		return
	}
	res, err := c.service.RegisterUser(con.Request.Context(), payload)
	if err != nil {
		utils.RespondError(con, 400, err)
		return
	}
	con.SetCookie("jwt", res.AccessToken, 60*60, "/", "localhost", false, true)
	con.JSON(201, res)
}
func (c *UserController) Login(con *gin.Context) {
	var payload models.LoginReq
	err := con.ShouldBindJSON(&payload)
	if err != nil {
		msg := fmt.Errorf("couldn't bind json: %s", err)
		utils.RespondError(con, 400, msg)
		return
	}
	res, err := c.service.LoginUser(con.Request.Context(), payload)
	if err != nil {
		utils.RespondError(con, 400, err)
		return
	}
	con.SetCookie("jwt", res.AccessToken, 60*60, "/", "localhost", false, true)
	con.JSON(200, res)
}
func (c *UserController) Logout(con *gin.Context) {
	con.SetCookie("jwt", "", -1, "", "", false, true)
	con.JSON(200, gin.H{"message": "Logout successful"})
}
