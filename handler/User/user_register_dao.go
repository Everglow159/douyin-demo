package User

import (
	"douyin.core/Model"
	"errors"
)

// 用户登录表
type UserLoginTable struct {
	Id       int64  `grom:"primary_key"`
	UserId   int64  `grom:"notnull"`
	Username string `grom:"notnull"`
	Password string `grom:"notnull"`
}

// 用户注册表数据操作结构体
type UserRigestDao struct {
}

// 新建用户
func NewUserLoginTable(username, password string) *UserLoginTable {
	return &UserLoginTable{Username: username, Password: password}
}

// 用户注册表数据操作结构体构造函数
func NewUserRigisterDao() *UserRigestDao {
	return &UserRigestDao{}
}

// 将用户信息持久化到数据库
func (u *UserRigestDao) RegistUsertoDb(userid int64, username, password string) error {
	user := UserLoginTable{
		UserId:   userid,
		Username: username,
		Password: password,
	}
	return Model.DB.Create(&user).Error
}

// 用户登录时检查用户的参数是否正确
func (u UserRigestDao) QueryUserLogin(username, password string, login *UserLoginTable) error {
	err := Model.DB.Where("username=?", username).First(&login).Error
	if err != nil {
		return err
	}
	if login.Password != password {
		err = errors.New("密码错误")
		return err
	}
	return nil
}
