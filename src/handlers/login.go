package handlers

import (
	//"github.com/gin-backend/src/contract"
	"github.com/gin-backend/src/db"
	"github.com/gin-backend/src/models"
	"github.com/gin-backend/src/services"
	"time"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware(s *services.Server) gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.Request.Header.Get("token")
		if len(token) == 0 {
			ParamsMissingJSONResp("the token not found", c)
			c.Abort()
			return
		}

		claims, err := s.ParseToken(token)
		if err != nil {
			TokenErrorJSONResp(err.Error(), c)
			c.Abort()
			return
		}

		c.Set("token", claims)
		c.Next()
	}
}

func Login(s *services.Server) gin.HandlerFunc {
	return func(c *gin.Context) {

		var req models.LoginReq

		if err := c.ShouldBindJSON(&req); err != nil {
			ParamsTypeErrorJSONResp(err.Error(), c)
			return
		}

		err := isStringRequiredParamsEmpty(req.UserName, req.UserPwd)
		if err != nil {
			ParamsMissingJSONResp(err.Error(), c)
			return
		}
		user := new(db.User)
		err = s.QueryObjectByCondition(user, "user_name", req.UserName)
		if err != nil {
			NotExistJSONResp(err.Error(), c)
			return
		}

		if req.UserPwd != user.UserPwd {
			PwdErrorJSONResp("", c)
			return
		}

		tokenExpiresAt := time.Now().Add(2 * time.Hour).Unix()
		token, err := s.GenToken(user.Id,
			user.UserName, db.UserRoleTypeName[user.UserRole], tokenExpiresAt)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		loginLog := &db.LoginLog{
			UserName:  user.UserName,
			LoginIp:   c.ClientIP(),
			LoginTime: time.Now().Unix(),
		}

		err = s.InsertOneObjertToDB(loginLog)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		SuccessfulJSONResp(&models.LoginInfo{
			UserNickName: user.UserNickName,
			UserRole:     db.UserRoleTypeName[user.UserRole],
			Expires:      tokenExpiresAt,
			Token:        token,
		}, c)
	}
}

func Register(s *services.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.RegisterReq
		if err := c.ShouldBindJSON(&req); err != nil {
			ParamsTypeErrorJSONResp(err.Error(), c)
			return
		}

		err := isStringRequiredParamsEmpty(req.UserName, req.UserPwd,
			req.UserPhoneNum, req.UserNickName, req.UserEmail, req.UserRole)
		if err != nil {
			ParamsMissingJSONResp(err.Error(), c)
			return
		}

		err = checkTheKeyRule(req.UserName)
		if err != nil {
			ParamsFormatErrorJSONResp(err.Error(), c)
			return
		}

		err = checkThePhoneNum(req.UserPhoneNum)
		if err != nil {
			ParamsFormatErrorJSONResp(err.Error(), c)
			return
		}

		err = checkTheEmail(req.UserEmail)
		if err != nil {
			ParamsFormatErrorJSONResp(err.Error(), c)
			return
		}

		role, ok := db.UserRoleTypeValue[req.UserRole]
		if !ok {
			ParamsValueJSONResp("role type error", c)
			return
		}

		user := new(db.User)
		err = s.QueryObjectByCondition(user, "user_name", req.UserName)
		if err == nil {
			UniqueIndexJSONResp("用户已存在", c)
			return
		}

		err = s.InsertOneObjertToDB(&db.User{
			UserName:     req.UserName,
			UserRole:     role,
			UserPwd:      req.UserPwd,
			UserNickName: req.UserNickName,
			UserPhoneNum: req.UserPhoneNum,
			UserEmail:    req.UserEmail,
		})
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		SuccessfulJSONResp("", c)
	}
}

func GetUserInfo(s *services.Server) gin.HandlerFunc {
	return func(c *gin.Context) {

		token, ok1 := c.Get("token")
		claims, ok2 := token.(*services.MyClaims)

		if !ok1 || !ok2 {
			ServerErrorJSONResp("get the token from context failed", c)
			return
		}

		user := new(db.User)
		err := s.QueryObjectById(user, claims.Id)
		if err != nil {
			NotExistJSONResp(err.Error(), c)
			return
		}

		SuccessfulJSONResp(&models.UserInfo{
			Id:           user.Id,
			UserName:     user.UserName,
			UserRole:     db.UserRoleTypeName[user.UserRole],
			UserNickName: user.UserNickName,
			UserPhoneNum: user.UserPhoneNum,
			UserEmail:    user.UserEmail,
			Expires:      claims.StandardClaims.ExpiresAt,
		}, c)

	}
}
