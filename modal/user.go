package modal

import "time"

const CtxGetUID = "UserID"

type User struct {
	ID       int
	UserID   int `gorm:"column:user_id"`
	Username string
	Password string
	CreateAt *time.Time `gorm:"column:createAt;<-:false"`
	UpdateAt *time.Time `gorm:"column:updateAt;<-:false"`
}

type ParamUser struct {
	Username    string `form:"username"`
	Password    string `form:"password"`
	RePassword  string `form:"re-password"`
	OriPassword string `form:"ori-password"`
}
