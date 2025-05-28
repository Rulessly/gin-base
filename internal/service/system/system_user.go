package system

import (
	"context"
	"github.com/rulessly/gin-base/internal/model/system"
	"github.com/rulessly/gin-base/internal/repository/systemRepo"
	"github.com/rulessly/gin-base/pkg/utils/encrypt"
)

type UserService struct {
	systemUserRepo systemRepo.BaseRepo
}

func (u *UserService) Register(ctx context.Context, req *system.User) error {
	_, err := u.systemUserRepo.GetOneByUsername(req.Username)
	if err == nil {
		return err
	}

	//TODO 可加入验证码的校验

	user := &system.User{
		Username: req.Username,
		Password: encrypt.BcryptHash(req.Password),
		Role:     req.Role,
		Creator:  "system",
	}

	err = u.systemUserRepo.Create(user)
	if err != nil {
		return err
	}
	return nil
}
