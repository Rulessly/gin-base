package systemRepo

import (
	"github.com/rulessly/gin-base/internal/model/system"
	"github.com/rulessly/gin-base/pkg/database"
)

type BaseRepo struct{}

func (b *BaseRepo) GetOneByID(id int64) (*system.User, error) {
	user := new(system.User)
	err := database.Mysql.Model(&system.User{}).Where("id = ?", id).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (b *BaseRepo) GetOneByUsername(username string) (*system.User, error) {
	user := new(system.User)
	err := database.Mysql.Model(user).Where("username = ?", username).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (b *BaseRepo) Create(user *system.User) error {
	err := database.Mysql.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}
