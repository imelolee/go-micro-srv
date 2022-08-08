package controller

import (
	"context"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-micro-srv/web/proto/userOrder"
	"go-micro-srv/web/utils"
	"net/http"
)

type OrderStu struct {
	EndDate   string `json:"end_date"`
	HouseId   string `json:"house_id"`
	StartDate string `json:"start_date"`
}

//下订单
func PostOrders(ctx *gin.Context) {
	//获取数据
	var orderStu OrderStu
	err := ctx.Bind(&orderStu)

	//校验数据
	if err != nil {
		fmt.Println("获取数据错误", err)
		return
	}
	//获取用户名
	userName := sessions.Default(ctx).Get("username")

	//处理数据  服务端处理业务
	microClient := userOrder.NewUserOrderService("userorder", utils.GetMicroClient())
	//调用服务
	resp, err := microClient.CreateOrder(context.TODO(), &userOrder.Request{
		StartDate: orderStu.StartDate,
		EndDate:   orderStu.EndDate,
		HouseId:   orderStu.HouseId,
		UserName:  userName.(string),
	})

	if err != nil {
		fmt.Println("远程服务错误: ", err)
		return
	}

	//返回数据
	ctx.JSON(http.StatusOK, resp)
}

//获取订单信息
func GetUserOrder(ctx *gin.Context) {
	//获取get请求传参
	role := ctx.Query("role")
	//校验数据
	if role == "" {
		fmt.Println("获取数据失败")
		return
	}

	//处理数据  服务端
	microClient := userOrder.NewUserOrderService("userorder", utils.GetMicroClient())
	//调用远程服务
	resp, _ := microClient.GetOrderInfo(context.TODO(), &userOrder.GetRequest{
		Role:     role,
		UserName: sessions.Default(ctx).Get("username").(string),
	})

	//返回数据
	ctx.JSON(http.StatusOK, resp)
}

type StatusStu struct {
	Action string `json:"action"`
	Reason string `json:"reason"`
}

//更新订单状态
func PutOrders(ctx *gin.Context) {
	//获取数据
	id := ctx.Param("id")
	var statusStu StatusStu
	err := ctx.Bind(&statusStu)

	//校验数据
	if err != nil || id == "" {
		fmt.Println("获取数据错误", err)
		return
	}

	//处理数据   更新订单状态
	microClient := userOrder.NewUserOrderService("userorder", utils.GetMicroClient())
	//调用元和产能服务
	resp, _ := microClient.UpdateStatus(context.TODO(), &userOrder.UpdateRequest{
		Action: statusStu.Action,
		Reason: statusStu.Reason,
		Id:     id,
	})

	//返回数据
	ctx.JSON(http.StatusOK, resp)
}
