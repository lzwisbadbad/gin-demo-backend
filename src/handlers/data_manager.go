package handlers

import (
	"github.com/gin-backend/src/db"
	"github.com/gin-backend/src/models"
	"github.com/gin-backend/src/services"
	"github.com/gin-gonic/gin"
	"time"
)

func AddData(s *services.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		//if err := checkTheAccessPermission(c, db.ADMIN, db.USER); err != nil {
		//	WithoutPermissionJSONResp(err.Error(), c)
		//	return
		//}

		// 1. 读取参数
		var req models.AddDataReq
		if err := c.ShouldBindJSON(&req); err != nil {
			ParamsTypeErrorJSONResp(err.Error(), c)
			return
		}

		// 2. 参数校验
		err := isStringRequiredParamsEmpty(req.DataId, req.DataContent)
		if err != nil {
			ParamsMissingJSONResp(err.Error(), c)
			return
		}

		// 3. 校验数据
		err = checkTheKeyRule(req.DataId)
		if err != nil {
			ParamsFormatErrorJSONResp(err.Error(), c)
			return
		}

		// 4. 入库
		data := &db.DataManager{
			DataId:      req.DataId,
			DataContent: req.DataContent,
			AddDataTime: time.Now().Unix(),
		}

		err = s.InsertOneObjertToDB(data)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		SuccessfulJSONResp("", c)
	}
}

func DeleteData(s *services.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		SuccessfulJSONResp("", c)
	}
}

func UpdateData(s *services.Server) gin.HandlerFunc {
	return func(c *gin.Context) {

		SuccessfulJSONResp("", c)
	}
}

func GetData(s *services.Server) gin.HandlerFunc {
	return func(c *gin.Context) {

		// 1. 读取参数
		var req models.GetDataReq
		if err := c.ShouldBindJSON(&req); err != nil {
			ParamsTypeErrorJSONResp(err.Error(), c)
			return
		}

		// 2. 参数校验
		err := isStringRequiredParamsEmpty(req.DataId)
		if err != nil {
			ParamsMissingJSONResp(err.Error(), c)
			return
		}

		// 3. 校验数据
		err = checkTheKeyRule(req.DataId)
		if err != nil {
			ParamsFormatErrorJSONResp(err.Error(), c)
			return
		}

		data := new(db.DataManager)
		err = s.QueryObjectByCondition(data, "data_id", req.DataId)
		if err != nil {
			NotExistJSONResp(err.Error(), c)
			return
		}

		SuccessfulJSONResp(&models.DataInfo{
			Id:          data.Id,
			UserName:    "lzw",
			DataId:      req.DataId,
			DataContent: data.DataContent,
			AddDataTime: data.AddDataTime,
		}, c)

	}
}
