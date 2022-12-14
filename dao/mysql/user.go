package mysqlDao

import (
	"errors"
	"gorm.io/gorm"
	g "main/global"
	"main/modal"
)

func CheckUserExist(username string) bool {
	var u modal.User
	result := g.Mdb.Model(&u).Where("username = ?", username).First(&u)
	return !errors.Is(result.Error, gorm.ErrRecordNotFound) //如果用户存在，返回true，不存在返回false
}

func AddUser(user *modal.User) error {
	return g.Mdb.Create(user).Error
}

func FindUser(user *modal.User) error {
	return g.Mdb.Model(user).Where("username = ? or user_id=?", user.Username, user.UserID).Find(user).Error
}

func RevisePwd(user *modal.User, NewPassword string) error {
	user.Password = NewPassword
	return g.Mdb.Save(user).Error
}
