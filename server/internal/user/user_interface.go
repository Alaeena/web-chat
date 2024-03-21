package user

import (
	"context"
	"github.com/gin-gonic/gin"
)

type Service[T User] interface {
	RegisterUser(ctx context.Context, payload RegisterReq) (*Response, error)
	LoginUser(ctx context.Context, payload LoginReq) (*Response, error)
	Logout(ctx context.Context, user T) error
}

type Controller interface {
	Detail(con *gin.Context)
	Register(con *gin.Context)
	Login(con *gin.Context)
	Logout(con *gin.Context)
}
