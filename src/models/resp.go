package models

type StandardResp struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type StandardRespWithPage struct {
	StandardResp
	Total int64 `json:"total"`
}

const (
	RESP_CODE_SUCCESS = 200

	RESP_CODE_SERVER_ERROR = 500

	RESP_CODE_UNIQUE_INDEX = 501

	RESP_CODE_NOT_EXIST = 502

	RESP_CODE_PWD_ERROR = 503

	RESP_CODE_PARAMS_TYPE_ERROR = 101

	RESP_CODE_PARAMS_MISSING = 102

	RESP_CODE_PARAMS_FORMAT_ERROR = 103

	RESP_CODE_PARAMS_VALUE_ERROR = 104

	RESP_CODE_WITHOUT_PERMISSION = 301

	RESP_CODE_TOKEN_ERROR = 303
)

const (
	RESP_MSG_SERVER_ERROR = "服务出错"

	RESP_MSG_PARAMS_TYPE_ERROR = "错误的参数类型"

	RESP_MSG_PARAMS_MISSING = "缺少必要参数"

	RESP_MSG_PARAMS_FORMAT_ERROR = "参数不符合规范"

	RESP_MSG_PARAMS_VALUE_ERROR = "参数值域不符合预期"

	RESP_MSG_WITHOUT_PERMISSION = "无权限访问"

	RESP_MSG_UNIQUE_INDEX = "唯一索引已经存在"

	RESP_MSG_NOT_EXIST = "唯一索引不存在"

	RESP_MSG_PWD_ERROR = "密码错误"

	RESP_MSG_TOKEN_ERROR = "请重新登录"
)

type UserInfo struct {
	Id           int32  `json:"id"`
	UserName     string `json:"userName"`
	UserNickName string `json:"userNickName"`
	UserRole     string `json:"userRole"`
	UserPhoneNum string `json:"userPhoneNum"`
	UserEmail    string `json:"userEmail"`
	Expires      int64  `json:"expires"`
}

type LoginInfo struct {
	UserNickName string `json:"userNickName"`
	UserRole     string `json:"userRole"`
	Expires      int64  `json:"expires"`
	Token        string `json:"token"`
}

type BaseRespInfo struct {
	Id       int32 `json:"id"`
	LastTime int64 `json:"lastTime"`
}

type DataInfo struct {
	Id          int32  `json:"id"`
	UserName    string `json:"userName"`
	DataId      string `json:"dataId"`
	DataContent string `json:"dataContent"`
	AddDataTime int64  `json:"addDataTime"`
}

type LoginLogInfo struct {
	BaseRespInfo
	UserName  string `json:"userName"`
	LoginTime int64  `json:"loginTime"`
	LoginIp   string `json:"loginIp"`
}
