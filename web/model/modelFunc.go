package model

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"go-micro-srv/web/utils"
	"mime/multipart"
)

// 创建全局redis 连接池 句柄
var RedisPool redis.Pool

// 创建函数, 初始化Redis连接池
func InitRedis() {
	RedisPool = redis.Pool{
		MaxIdle:         20,
		MaxActive:       50,
		MaxConnLifetime: 60 * 5,
		IdleTimeout:     60,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}

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

func Login(mobile, pwd string) (string, error) {
	var user User

	// password加密处理
	m5 := md5.New()
	m5.Write([]byte(pwd))
	pwd_hash := hex.EncodeToString(m5.Sum(nil))

	err := GlobalConn.Select("name").
		Where("mobile = ?", mobile).Where("password_hash = ?", pwd_hash).Find(&user).Error

	return user.Name, err
}

// 获取用户信息
func GetUserInfo(username string) (User, error) {
	var user User
	err := GlobalConn.Where("name = ?", username).Find(&user).Error
	return user, err
}

// 更新用户名
func UpdateUsername(newName, oldName string) error {
	err := GlobalConn.Model(new(User)).Where("name = ?", oldName).Update("name", newName).Error
	return err
}

// 七牛云上传
func UpLoadFile(file multipart.File, fileName string, fileSize int64) (key string, err error) {

	key = utils.RandFileName(fileName)

	putPolicy := storage.PutPolicy{
		Scope: utils.Bucket,
	}
	mac := qbox.NewMac(utils.AccessKey, utils.SecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Zone:          &storage.ZoneHuadong,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}

	putExtra := storage.PutExtra{}

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	err = formUploader.Put(context.Background(), &ret, upToken, key, file, fileSize, &putExtra)

	return key, err
}

// 更新用户头像
func UpdateAvatar(username, avatar string) error {
	err := GlobalConn.Model(new(User)).Where("name = ?", username).Update("avatar_url", avatar).Error
	return err
}
