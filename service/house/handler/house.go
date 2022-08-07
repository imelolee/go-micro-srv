package handler

import (
	"context"
	"fmt"

	"house/model"
	house "house/proto"
	"house/utils"
	"strconv"
)

type House struct{}

func (e *House) PubHouse(ctx context.Context, req *house.Request, rsp *house.Response) error {
	//上传房屋业务  把获取到的房屋数据插入数据库
	fmt.Println("-----------", req)
	houseId, err := model.AddHouse(req)
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		return nil
	}

	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)

	var h house.HouseData
	h.HouseId = strconv.Itoa(houseId)
	rsp.Data = &h

	return nil
}

func (e *House) UploadHouseImg(ctx context.Context, req *house.ImgRequest, resp *house.ImgResponse) error {

	key, err := utils.UpLoadQiniu(req.ImgData, req.FileExt, req.FileSize)

	ImgPath := utils.Domain + key
	//把凭证存储到数据库中
	err = model.SaveHouseImg(req.HouseId, ImgPath)
	if err != nil {
		resp.Errno = utils.RECODE_DBERR
		resp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		return nil
	}

	resp.Errno = utils.RECODE_OK
	resp.Errmsg = utils.RecodeText(utils.RECODE_OK)

	var img house.ImgData
	img.Url = ImgPath

	resp.Data = &img

	return nil
}

func (e *House) GetHouseInfo(ctx context.Context, req *house.GetRequest, resp *house.GetResponse) error {
	//根据用户名获取所有的房屋数据

	houseInfos, err := model.GetUserHouse(req.UserName)
	if err != nil {
		resp.Errno = utils.RECODE_DBERR
		resp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		return nil
	}

	resp.Errno = utils.RECODE_OK
	resp.Errmsg = utils.RecodeText(utils.RECODE_OK)

	var getData house.GetData
	getData.Houses = houseInfos

	resp.Data = &getData

	return nil
}

func (e *House) GetHouseDetail(ctx context.Context, req *house.DetailRequest, resp *house.DetailResponse) error {
	//根据houseId获取所有的返回数据
	respData, err := model.GetHouseDetail(req.HouseId, req.UserName)
	if err != nil {
		resp.Errno = utils.RECODE_DATAERR
		resp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
		return nil
	}

	resp.Errno = utils.RECODE_OK
	resp.Errmsg = utils.RecodeText(utils.RECODE_OK)
	resp.Data = &respData

	return nil
}

func (e *House) GetIndexHouse(ctx context.Context, req *house.IndexRequest, resp *house.GetResponse) error {
	//获取房屋信息
	houseResp, err := model.GetIndexHouse()
	if err != nil {
		resp.Errno = utils.RECODE_DBERR
		resp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		return nil
	}

	resp.Errno = utils.RECODE_OK
	resp.Errmsg = utils.RecodeText(utils.RECODE_OK)

	resp.Data = &house.GetData{Houses: houseResp}

	return nil
}

func (e *House) SearchHouse(ctx context.Context, req *house.SearchRequest, resp *house.GetResponse) error {
	//根据传入的参数,查询符合条件的房屋信息
	houseResp, err := model.SearchHouse(req.Aid, req.Sd, req.Ed, req.Sk)
	if err != nil {
		resp.Errno = utils.RECODE_DATAERR
		resp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
		return nil
	}

	resp.Errno = utils.RECODE_OK
	resp.Errmsg = utils.RecodeText(utils.RECODE_OK)

	resp.Data = &house.GetData{Houses: houseResp}
	return nil
}
