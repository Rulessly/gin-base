package system

import "github.com/rulessly/gin-base/internal/model"

type User struct {
	model.CommonModel
	ID       int64  `gorm:"column:id" json:"id"`
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
	Role     int    `gorm:"column:role" json:"role"`
	Nickname string `gorm:"column:nickname" json:"nickname"`
	Creator  string `gorm:"column:creator" json:"creator"`
}
