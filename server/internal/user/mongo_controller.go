package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/db/mongo/mongodb"
	"server/db/scylla/scylladb"
	"server/internal/websocket"
	"server/utils"
	"time"
)

type mongoController struct {
	service Service[mongodb.User]
	hub     *websocket.Hub
}

func GetMongoController(service Service[mongodb.User], hub *websocket.Hub) Controller {
	return &mongoController{service: service, hub: hub}
}
func (c *mongoController) newLogin(con *gin.Context, res *Response) {
	logMsg := &scylladb.NewReportsParams{
		UserId:  res.ID,
		Type:    fmt.Sprintf("New user's login at %v", time.Now()),
		Message: "Login",
	}
	c.hub.Log <- logMsg
	con.SetCookie("jwt", res.AccessToken, 60*60, "/", "localhost", false, true)
	con.JSON(201, res)
}

func (c *mongoController) Register(con *gin.Context) {
	var payload RegisterReq
	err := con.ShouldBindJSON(&payload)
	if err != nil || payload.Email == "" || payload.Password == "" || payload.Username == "" {
		msg := fmt.Errorf("couldn't bind json: %s", err)
		utils.RespondError(con, 400, msg)
		return
	}
	res, err := c.service.RegisterUser(con.Request.Context(), payload)
	if err != nil {
		utils.RespondError(con, 400, err)
		return
	}
	c.newLogin(con, res)
}
func (c *mongoController) Login(con *gin.Context) {
	var payload LoginReq
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
	c.newLogin(con, res)
}
func (c *mongoController) Logout(con *gin.Context) {
	con.SetCookie("jwt", "", -1, "", "", false, true)
	user := con.MustGet("user").(mongodb.User)
	err := c.service.Logout(con.Request.Context(), user)

	if err != nil {
		utils.RespondError(con, 400, err)
		return
	}
	logMsg := &scylladb.NewReportsParams{
		UserId:  user.ID.Hex(),
		Type:    fmt.Sprintf("New user's logout at %v", time.Now()),
		Message: "Logout",
	}
	c.hub.Log <- logMsg
	con.JSON(200, gin.H{"Message": "Logout successful"})
}
func (c *mongoController) Detail(con *gin.Context) {
	user := con.MustGet("user").(mongodb.User)

	con.JSON(200, MapMongoUser(user, user.AccessToken))
}
