package handler

import (
	"context"
	"fmt"

	"user/model"
	"user/utils"

	"math/rand"
	"time"
	pb "user/proto"
)

type User struct{}

// SendSms 发送短信验证码
func (e *User) SendSms(ctx context.Context, req *pb.SmsRequest, rsp *pb.RegResponse) error {
	result := model.CheckImgCode(req.Uuid, req.ImgCode)

	if result {
		// 发送短信
		rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
		smsCode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
		templateParam := `{"code":"` + smsCode + `"}`

		fmt.Println(templateParam)

		// 验证码存入redis
		err := model.SaveSmsCode(req.Phone, smsCode)
		if err != nil {
			// 存储验证码失败
			rsp.Errno = utils.RECODE_DBERR
			rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		}

		rsp.Errno = utils.RECODE_OK
		rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)
	} else {
		// 发送错误信息
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
	}
	return nil
}

// Register 用户注册
func (e *User) Register(ctx context.Context, req *pb.RegRequest, rsp *pb.RegResponse) error {
	// 校验验证码是否正确
	err := model.CheckSmsCode(req.Mobile, req.SmsCode)
	if err == nil {
		// 注册用户写入数据库
		err := model.RegisterUser(req.Mobile, req.Password)
		if err != nil {
			rsp.Errno = utils.RECODE_DBERR
			rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		} else {
			rsp.Errno = utils.RECODE_OK
			rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)
		}
	} else {
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
	}

	return nil
}

// Login 用户登录
func (e *User) Login(ctx context.Context, req *pb.RegRequest, rsp *pb.RegResponse) error {

	// 获取数据库数据
	err := model.Login(req.Mobile, req.Password)
	if err == nil {
		// 登录成功
		rsp.Errno = utils.RECODE_OK
		rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)

	} else {
		// 登录失败
		rsp.Errno = utils.RECODE_LOGINERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_LOGINERR)
	}

	return nil
}

func (e *User) AuthUpdate(ctx context.Context, req *pb.AuthRequest, rsp *pb.AuthResponse) error {
	//调用借口校验realName和idcard是否匹配

	//存储真实姓名和真是身份证号  数据库
	err := model.SaveRealName(req.UserName, req.RealName, req.IdCard)
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		return nil
	}

	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)

	return nil
}

func (e *User) MicroGetUser(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	//根据用户名获取用户信息 在mysql数据库中
	myUser, err := model.GetUserInfo(req.Name)
	if err != nil {
		rsp.Errno = utils.RECODE_USERERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_USERERR)
		return err
	}
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)

	//获取一个结构体对象
	var userInfo pb.UserInfo
	userInfo.UserId = int32(myUser.ID)
	userInfo.Name = myUser.Name
	userInfo.Mobile = myUser.Mobile
	userInfo.RealName = myUser.Real_name
	userInfo.IdCard = myUser.Id_card
	userInfo.AvatarUrl = myUser.Avatar_url

	rsp.Data = &userInfo

	return nil
}

func (e *User) UpdateUserName(ctx context.Context, req *pb.UpdateRequest, resp *pb.UpdateResponse) error {
	//根据传递过来的用户名更新数据中新的用户名
	err := model.UpdateUserName(req.OldName, req.NewName)
	if err != nil {
		fmt.Println("更新失败", err)
		resp.Errno = utils.RECODE_DATAERR
		resp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
		//micro规定如果有错误,服务端只给客户端返回错误信息,不返回resp,如果没有错误,就返回resp
		return nil
	}

	resp.Errno = utils.RECODE_OK
	resp.Errmsg = utils.RecodeText(utils.RECODE_OK)
	var nameData pb.NameData
	nameData.Name = req.NewName

	resp.Data = &nameData

	return nil
}

func (e *User) UploadAvatar(ctx context.Context, req *pb.UploadRequest, resp *pb.UploadResponse) error {
	// 上传头像到云存储
	formFile := req.Avatar
	fileSize := req.FileSize
	fileName := req.FileExt
	fmt.Println("Filename：", fileName)
	// 将文件上传到七牛云
	key, err := utils.UpLoadQiniu(formFile, fileName, fileSize)

	if err != nil {
		fmt.Println("存储用户头像错误", err)
		resp.Errno = utils.RECODE_DBERR
		resp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		return nil
	}

	resp.Errno = utils.RECODE_OK
	resp.Errmsg = utils.RecodeText(utils.RECODE_OK)

	var uploadData pb.UploadData
	uploadData.AvatarUrl = utils.Domain + key
	resp.Data = &uploadData
	return nil
}
