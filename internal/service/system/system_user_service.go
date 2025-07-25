package system

import (
	"context"
	"errors"
	"github.com/rulessly/gin-base/internal/ent/systemuser"
	ststemReq "github.com/rulessly/gin-base/internal/request/system"
	"github.com/rulessly/gin-base/pkg/database"
	"github.com/rulessly/gin-base/pkg/utils/encrypt"
	"github.com/spf13/cast"
)

type UserService struct {
}

func NewSystemUserService() *UserService {
	return &UserService{}
}

// Register 用户注册
func (*UserService) Register(ctx context.Context, req ststemReq.RegisterRequest) error {
	_, err := database.DB.SystemUser.Query().Select(systemuser.FieldUsername).Where(systemuser.UsernameEQ(req.Username)).First(ctx)
	if err == nil {
		return errors.New("用户名已存在")
	}
	if _, err = database.DB.SystemUser.
		Create().
		SetUsername(req.Username).
		SetPassword(encrypt.BcryptHash(req.Password)).
		SetCreator("system").
		SetRole(cast.ToInt8(req.Role)).
		SetNickname(req.Nickname).
		Save(ctx); err != nil {
		return err
	}
	return nil
}
