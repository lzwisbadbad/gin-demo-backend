package models

type RegisterReq struct {
	UserName     string `json:"userName"`
	UserPwd      string `json:"userPwd"`
	UserRole     string `json:"userRole"`
	UserNickName string `json:"userNickName"`
	UserPhoneNum string `json:"userPhoneNum"`
	UserEmail    string `json:"userEmail"`
}

type LoginReq struct {
	UserName string `json:"userName"`
	UserPwd  string `json:"userPwd"`
}

type AddDataReq struct {
	DataId      string `json:"dataId"`
	DataContent string `json:"dataContent"`
}

type GetDataReq struct {
	DataId string `json:"dataId"`
}
