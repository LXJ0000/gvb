package flags

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/utils"
)

// CreateUser 创建用户
func CreateUser(permission string) {
	//	username nickname password RePassword email
	var (
		userName   string
		nickName   string
		password   string
		rePassword string
		email      string
	)
	fmt.Print("请输入用户名：")
	_, _ = fmt.Scan(&userName)
	fmt.Print("请输入昵称：")
	_, _ = fmt.Scan(&nickName)
	fmt.Print("请输入密码：")
	_, _ = fmt.Scan(&password)
	fmt.Print("请确认密码：")
	_, _ = fmt.Scan(&rePassword)
	fmt.Print("请输入邮箱：")
	_, _ = fmt.Scanln(&email)

	//判断用户名是否存在
	var user models.UserModel
	if err := global.DB.Where("user_name = ?", userName).First(&user).Error; err == nil {
		global.Log.Error("用户名已存在，请重新输入")
		return
	}
	//密码验证
	if password != rePassword {
		global.Log.Error("密码不一致，请重新输入")
		return
	}
	//密码hash
	hashPwd := utils.SHA1(password)

	//身份
	role := ctype.PermissionUser
	if permission == "admin" {
		role = ctype.PermissionAdmin
	}

	// 头像
	//1 默认
	//2 随机
	avatar := "/static/avatar/1.jpg"

	//入库
	user = models.UserModel{
		UserName:   userName,
		NickName:   nickName,
		Password:   hashPwd,
		Email:      email,
		Role:       role,
		SignStatus: ctype.SignEmail,
		IP:         "127.0.0.1",
		Addr:       "local",
		Avatar:     avatar,
	}
	if err := global.DB.Create(&user).Error; err != nil {
		global.Log.Error(err)
		return
	}

	global.Log.Infof("用户%s创建成功", userName)
}
