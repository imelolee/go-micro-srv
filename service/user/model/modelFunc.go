package model

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

// 校验图片验证码
func CheckImgCode(uuid, imgCode string) bool {
	// 链接 redis --- 从链接池中获取链接
	/*	conn, err := redis.Dial("tcp", "192.168.6.108:6379")
		if err != nil {
			fmt.Println("redis.Dial err:", err)
			return false
		}*/
	conn := RedisPool.Get()
	defer conn.Close()

	// 查询 redis 数据
	code, err := redis.String(conn.Do("get", uuid))

	if err != nil {
		fmt.Println("查询错误 err:", err)
		return false
	}

	// 返回校验结果
	return code == imgCode
}

// 存储短信验证码
func SaveSmsCode(phone, code string) error {
	// 链接 Redis --- 从链接池中获取一条链接
	conn := RedisPool.Get()
	defer conn.Close()

	// 存储短信验证码到 redis 中
	_, err := conn.Do("setex", phone+"_code", 60*3, code)

	return err
}

// 校验短信验证码
func CheckSmsCode(phone, code string) error {
	// 连接redis
	conn := RedisPool.Get()

	// 根据key 获取value
	sms_code, err := redis.String(conn.Do("get", phone+"_code"))
	if err != nil {
		fmt.Println("Redis get err: ", err)
		return err
	}

	if sms_code != code {
		return errors.New("验证码错误.")
	}

	return nil
}

// 注册用户写入数据库
func RegisterUser(mobile, pwd string) error {
	var user User
	user.Name = mobile
	user.Mobile = mobile
	// 密码加密
	m5 := md5.New()
	m5.Write([]byte(pwd))
	pwd_hash := hex.EncodeToString(m5.Sum(nil))

	user.Password_hash = pwd_hash

	// 插入数据库
	return GlobalConn.Create(&user).Error
}

func Login(mobile, password string) error {
	var user User

	// password加密处理
	m5 := md5.New()
	m5.Write([]byte(password))
	pwd_hash := hex.EncodeToString(m5.Sum(nil))

	err := GlobalConn.Select("name").
		Where("mobile = ?", mobile).Where("password_hash = ?", pwd_hash).Find(&user).Error

	return err
}

func GetUserInfo(userName string) (User, error) {
	//连接数据库
	var user User
	err := GlobalConn.Where("name = ?", userName).Find(&user).Error
	return user, err
}

//存储用户真实姓名
func SaveRealName(userName, realName, idCard string) error {
	return GlobalConn.Model(new(User)).Where("name = ?", userName).
		Updates(map[string]interface{}{"real_name": realName, "id_card": idCard}).Error
}

//更新用户名
func UpdateUserName(oldName, newName string) error {
	//更新  链式调用
	return GlobalConn.Model(new(User)).
		Where("name = ?", oldName).
		Update("name", newName).Error
}
