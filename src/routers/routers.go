package routers

import (
	"github.com/gin-backend/src/handlers"
	"github.com/gin-backend/src/services"
)

const ROUTERS_HEADER = "/gin-backend"

const ROUTERS_USER = ROUTERS_HEADER + "/user"

const ROUTERS_MANAGE_DATA = ROUTERS_HEADER + "/manage_data"

func LoadNoTokenRouter(s *services.Server) {
	routerGroup := s.GetGinEngine().Group(ROUTERS_USER)
	{
		routerGroup.POST("/register", handlers.Register(s))
		routerGroup.POST("/login", handlers.Login(s))
		routerGroup.POST("/test2", handlers.Login(s))
	}
}

func LoadUserRouter(s *services.Server) {
	routerGroup := s.GetGinEngine().Group(ROUTERS_USER)
	{
		routerGroup.GET("/getuserinfo", handlers.GetUserInfo(s))
	}
}

func LoadManageDataRouter(s *services.Server) {
	routerGroup := s.GetGinEngine().Group(ROUTERS_MANAGE_DATA)
	{
		routerGroup.POST("/addData", handlers.AddData(s))
		routerGroup.POST("/updateData", handlers.UpdateData(s))
		routerGroup.POST("/deleteData", handlers.DeleteData(s))
		routerGroup.POST("/getData", handlers.GetData(s))
	}
}
