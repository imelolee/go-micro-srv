package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/afocus/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-micro/plugins/v4/registry/consul"
	"github.com/gomodule/redigo/redis"
	"go-micro-srv/web/model"
	"go-micro-srv/web/proto/getCaptcha"
	"go-micro-srv/web/proto/user"
	"go-micro-srv/web/utils"
	"go-micro.dev/v4"
	"image/png"
	"net/http"
)

// 获取session信息
func GetSession(ctx *gin.Context) {
	// 初始化map
	resp := make(map[string]interface{})

	// 获取session数据
	s := sessions.Default(ctx)

	username := s.Get("username")

	if username == nil {
		// 用户没有登录
		fmt.Println("未登录.")
		resp["errno"] = utils.RECODE_SESSIONERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
	} else {
		// 用户已登录
		fmt.Println("已登录.")
		resp["errno"] = utils.RECODE_OK
		resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)

		var nameData struct {
			Name string `json:"name"`
		}
		nameData.Name = username.(string)
		resp["data"] = nameData
	}

	ctx.JSON(http.StatusOK, resp)
}

// 获取图片验证码
func GetImageCd(ctx *gin.Context) {
	// 获取图片验证码 uuid
	uuid := ctx.Param("uuid")

	// 指定 consul 服务发现
	consulReg := consul.NewRegistry()
	consulService := micro.NewService(
		micro.Registry(consulReg),
	)

	// 初始化客户端
	microClient := getCaptcha.NewGetCaptchaService("getcaptcha", consulService.Client())

	// 调用远程函数
	resp, err := microClient.Call(context.TODO(), &getCaptcha.Request{Uuid: uuid})
	if err != nil {
		fmt.Println("未找到远程服务...", err)
		return
	}

	// 将得到的数据,反序列化,得到图片数据
	var img captcha.Image
	json.Unmarshal(resp.Img, &img)

	// 将图片写出到 浏览器.
	png.Encode(ctx.Writer, img)

}

func GetSmscd(ctx *gin.Context) {

	phone := ctx.Param("phone")
	imgCode := ctx.Query("text")
	uuid := ctx.Query("id")

	consulReg := consul.NewRegistry()
	consulService := micro.NewService(
		micro.Registry(consulReg),
	)

	microClient := user.NewUserService("user", consulService.Client())
	resp, err := microClient.SendSms(context.TODO(), &user.CallRequest{Phone: phone, ImgCode: imgCode, Uuid: uuid})
	if err != nil {
		fmt.Println("未找到远程服务...", err)
		return
	}

	//fmt.Println(phone, imgCode, uuid)

	// 发送响应结果
	ctx.JSON(http.StatusOK, resp)
}

// 发送注册信息
func PostRet(ctx *gin.Context) {

	// 获取数据
	var regData struct {
		Mobile   string `json:"mobile"`
		Password string `json:"password"`
		SmsCode  string `json:"sms_code"`
	}

	ctx.Bind(&regData)

	// 初始化客户端
	microService := utils.InitMicro()
	microClient := user.NewUserService("user", microService.Client())

	rsp, err := microClient.Register(context.TODO(), &user.RegReq{
		Mobile:   regData.Mobile,
		SmsCode:  regData.SmsCode,
		Password: regData.Password,
	})

	if err != nil {
		fmt.Println("用户注册找不到远程服务，", err)
		return
	}

	// 写给浏览器
	ctx.JSON(http.StatusOK, rsp)
}

// 获取地域信息
func GetArea(ctx *gin.Context) {
	var areas []model.Area
	// 将数据写入redis
	conn := model.RedisPool.Get()
	areaData, _ := redis.Bytes(conn.Do("get", "areaData"))

	// redis中没有数据
	if len(areaData) == 0 {
		fmt.Println("从mysql获取数据...")
		model.GlobalConn.Find(&areas)
		areaBuf, _ := json.Marshal(areas)

		conn.Do("set", "areaData", areaBuf)
	} else {
		// redis中有数据
		fmt.Println("从redis获取数据...")
		json.Unmarshal(areaData, &areas)
	}

	resp := make(map[string]interface{})

	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	resp["data"] = areas

	ctx.JSON(http.StatusOK, resp)
}

// 处理登录业务
func PostLogin(ctx *gin.Context) {
	// 获取前端数据
	var loginData struct {
		Mobile   string `json:"mobile"`
		Password string `json:"password"`
	}
	ctx.Bind(&loginData)

	resp := make(map[string]interface{})

	// 获取数据库数据
	username, err := model.Login(loginData.Mobile, loginData.Password)
	if err == nil {
		// 登录成功
		resp["errno"] = utils.RECODE_OK
		resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)

		// 将登录状态保存到session
		session := sessions.Default(ctx)
		session.Set("username", username)
		session.Save()

	} else {
		// 登录失败
		resp["errno"] = utils.RECODE_LOGINERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_LOGINERR)
	}

	ctx.JSON(http.StatusOK, resp)
}

func DeleteSession(ctx *gin.Context) {
	fmt.Println("-----")
	resp := make(map[string]interface{})
	// 初始化session对象
	s := sessions.Default(ctx)
	s.Delete("username")
	err := s.Save()
	if err != nil {

		resp["errno"] = utils.RECODE_IOERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_IOERR)
	} else {
		resp["errno"] = utils.RECODE_OK
		resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	}
	ctx.JSON(http.StatusOK, resp)
}

func GetUserInfo(ctx *gin.Context) {
	resp := make(map[string]interface{})
	defer ctx.JSON(http.StatusOK, resp)

	// 获取 Session, 得到 当前 用户信息
	s := sessions.Default(ctx)
	username := s.Get("username")
	// 判断用户名是否存在.
	if username == nil { // 用户没登录, 但进入该页面, 恶意进入.
		resp["errno"] = utils.RECODE_SESSIONERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
		return // 如果出错, 报错, 退出
	}

	// 根据用户名, 获取 用户信息  ---- 查 MySQL 数据库  user 表.
	user, err := model.GetUserInfo(username.(string))
	if err != nil {
		resp["errno"] = utils.RECODE_DBERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DBERR)
		return // 如果出错, 报错, 退出
	}

	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)

	temp := make(map[string]interface{})
	temp["user_id"] = user.ID
	temp["name"] = user.Name
	temp["mobile"] = user.Mobile
	temp["real_name"] = user.Real_name
	temp["id_card"] = user.Id_card
	temp["avatar_url"] = user.Avatar_url

	resp["data"] = temp

}

// 更新用户信息
func PutUserInfo(ctx *gin.Context) {
	// 当前用户名
	s := sessions.Default(ctx)
	username := s.Get("username")

	// 新用户名
	var nameData struct {
		Name string `json:"name"`
	}
	ctx.Bind(&nameData)

	resp := make(map[string]interface{})
	defer ctx.JSON(http.StatusOK, resp)

	// 更新用户名
	err := model.UpdateUsername(nameData.Name, username.(string))
	if err != nil {
		resp["errno"] = utils.RECODE_DBERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DBERR)
		return
	}

	s.Set("username", nameData.Name)
	err = s.Save()
	if err != nil {
		resp["errno"] = utils.RECODE_SESSIONERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
		return
	}
	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	resp["data"] = nameData
}

// 上传用户头像
func PostAvatar(ctx *gin.Context) {
	// 获取静态文件对象
	formFile, fileHeader, _ := ctx.Request.FormFile("avatar")

	/*// 上传文件到项目中
	err := ctx.SaveUploadedFile(formFile, "test/"+formFile.Filename)
	if err != nil {
		fmt.Println(err)
	}*/

	username := sessions.Default(ctx).Get("username")

	// 上传头像到云存储
	fileSize := fileHeader.Size
	fileName := fileHeader.Filename
	key, err := model.UpLoadFile(formFile, fileName, fileSize)
	if err != nil {
		fmt.Println("Upload err: ", err)
		return
	}

	resp := make(map[string]interface{})
	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)

	temp := make(map[string]interface{})
	avatar_url := utils.Domain + key
	temp["avatar_url"] = avatar_url
	resp["data"] = temp

	err = model.UpdateAvatar(username.(string), avatar_url)
	if err != nil {
		fmt.Println("UpdateAvatar err: ", err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
