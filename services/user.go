package services

import (
	"errors"
	mysqlDao "main/dao/mysql"
	"main/modal"
	"main/utils"
)

func Enroll(p *modal.ParamUser) error {
	//判断用户是否存在
	if mysqlDao.CheckUserExist(p.Username) {
		return errors.New("user has existed")
	}
	//生成user-id
	userID, err := utils.GetUserID()
	if err != nil {
		return errors.New("generate user failed")
	}
	//创建user实例
	User := new(modal.User)
	User.UserID = userID
	User.Username = p.Username
	//密码加密
	User.Password = utils.Md5(p.Password)
	//入库
	err = mysqlDao.AddUser(User)
	if err != nil {
		return errors.New("generate user failed")
	}
	return nil
}

func Login(p *modal.ParamUser) (string, error) {
	//判断用户是否存在
	if !mysqlDao.CheckUserExist(p.Username) {
		return "", errors.New("the user does not exist")
	}
	//判断密码是否正确
	user := new(modal.User)
	user.Username = p.Username
	err := mysqlDao.FindUser(user)
	if err != nil {
		return "", err
	}
	p.Password = utils.Md5(p.Password)
	if user.Password != p.Password {
		return "", errors.New("wrong password")
	}
	//发放凭证
	token, err := utils.GenToken(user.UserID)
	if err != nil {
		return "", errors.New("get token failed")
	}
	return token, nil
}

func RevisePwd(p *modal.ParamUser, uid int) error {
	//判断密码是否正确
	user := new(modal.User)
	user.UserID = uid
	err := mysqlDao.FindUser(user)
	if err != nil {
		return err
	}
	p.OriPassword = utils.Md5(p.OriPassword)
	if p.OriPassword != user.Password {
		return errors.New("wrong password")
	}
	//修改密码
	err = mysqlDao.RevisePwd(user, utils.Md5(p.Password))
	if err != nil {
		return err
	}
	return nil
}
