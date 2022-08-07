package controller

import (
	"context"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-micro-srv/web/proto/house"
	"go-micro-srv/web/utils"
	"net/http"
)

// 获取已发布房源信息  假数据
func GetUserHouses(ctx *gin.Context) {

	//获取用户名
	username := sessions.Default(ctx).Get("username")

	microClient := house.NewHouseService("house", utils.GetMicroClient())

	//调用远程服务
	resp, err := microClient.GetHouseInfo(context.TODO(), &house.GetRequest{UserName: username.(string)})
	if err != nil {
		fmt.Println("远程服务调用失败：", err)
		return
	}

	//返回数据
	ctx.JSON(http.StatusOK, resp)
}

type HouseStu struct {
	Acreage   string   `json:"acreage"`
	Address   string   `json:"address"`
	AreaId    string   `json:"area_id"`
	Beds      string   `json:"beds"`
	Capacity  string   `json:"capacity"`
	Deposit   string   `json:"deposit"`
	Facility  []string `json:"facility"`
	MaxDays   string   `json:"max_days"`
	MinDays   string   `json:"min_days"`
	Price     string   `json:"price"`
	RoomCount string   `json:"room_count"`
	Title     string   `json:"title"`
	Unit      string   `json:"unit"`
}

//发布房源
func PostHouses(ctx *gin.Context) {

	//获取数据   bind数据的时候不带自动转换   c.getInt()
	var houseStu HouseStu
	err := ctx.Bind(&houseStu)

	//校验数据
	if err != nil {
		fmt.Println("获取数据错误", err)
		return
	}

	//获取用户名
	userName := sessions.Default(ctx).Get("username")

	//处理数据  服务端处理
	microClient := house.NewHouseService("house", utils.GetMicroClient())
	//调用远程服务

	resp, err := microClient.PubHouse(context.TODO(), &house.Request{
		Acreage:   houseStu.Acreage,
		Address:   houseStu.Address,
		AreaId:    houseStu.AreaId,
		Beds:      houseStu.Beds,
		Capacity:  houseStu.Capacity,
		Deposit:   houseStu.Deposit,
		Facility:  houseStu.Facility,
		MaxDays:   houseStu.MaxDays,
		MinDays:   houseStu.MinDays,
		Price:     houseStu.Price,
		RoomCount: houseStu.RoomCount,
		Title:     houseStu.Title,
		Unit:      houseStu.Unit,
		UserName:  userName.(string),
	})

	if err != nil {
		fmt.Println("远程服务调用错误", err)
		return
	}

	//返回数据
	ctx.JSON(http.StatusOK, resp)
}

//上传房屋图片
func PostHousesImage(ctx *gin.Context) {
	//获取数据
	houseId := ctx.Param("id")
	file, fileHeader, err := ctx.Request.FormFile("house_image")
	//校验数据
	if houseId == "" || err != nil {
		fmt.Println("传入数据不完整", err)
		return
	}

	if fileHeader.Size > 50000000 {
		fmt.Println("文件过大,请重新选择")
		return
	}

	//获取文件字节切片
	buf := make([]byte, fileHeader.Size)
	file.Read(buf)

	//处理数据  服务中实现
	microClient := house.NewHouseService("house", utils.GetMicroClient())
	//调用服务
	resp, _ := microClient.UploadHouseImg(context.TODO(), &house.ImgRequest{
		HouseId:  houseId,
		ImgData:  buf,
		FileExt:  fileHeader.Filename,
		FileSize: fileHeader.Size,
	})

	//返回数据
	ctx.JSON(http.StatusOK, resp)
}

//获取房屋详情
func GetHouseInfo(ctx *gin.Context) {

	//获取数据
	houseId := ctx.Param("id")
	//校验数据
	if houseId == "" {
		fmt.Println("获取数据错误")
		return
	}
	userName := sessions.Default(ctx).Get("username")

	//处理数据
	microClient := house.NewHouseService("house", utils.GetMicroClient())
	//调用远程服务
	resp, _ := microClient.GetHouseDetail(context.TODO(), &house.DetailRequest{
		HouseId:  houseId,
		UserName: userName.(string),
	})

	//返回数据
	ctx.JSON(http.StatusOK, resp)
}

func GetIndex(ctx *gin.Context) {
	//处理数据
	microClient := house.NewHouseService("house", utils.GetMicroClient())
	//调用服务
	resp, _ := microClient.GetIndexHouse(context.TODO(), &house.IndexRequest{})

	ctx.JSON(http.StatusOK, resp)
}

//搜索房屋
func GetHouses(ctx *gin.Context) {
	//获取数据
	//areaId
	aid := ctx.Query("aid")
	//start day
	sd := ctx.Query("sd")
	//end day
	ed := ctx.Query("ed")
	//排序方式
	sk := ctx.Query("sk")
	//page  第几页
	//ctx.Query("p")
	//校验数据
	if aid == "" || sd == "" || ed == "" || sk == "" {
		fmt.Println("传入数据不完整")
		return
	}

	microClient := house.NewHouseService("house", utils.GetMicroClient())
	//调用远程服务
	resp, _ := microClient.SearchHouse(context.TODO(), &house.SearchRequest{
		Aid: aid,
		Sd:  sd,
		Ed:  ed,
		Sk:  sk,
	})

	//返回数据
	ctx.JSON(http.StatusOK, resp)

}
