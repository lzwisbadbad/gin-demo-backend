package main

import (
	"fmt"
	"github.com/gin-backend/src/configs"
	db2 "github.com/gin-backend/src/db"
	"github.com/gin-backend/src/handlers"
	"github.com/gin-backend/src/loggers"
	"github.com/gin-backend/src/models"
	"github.com/gin-backend/src/routers"
	"github.com/gin-backend/src/services"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func main() {

	var err error

	conf, err := configs.InitConfig(configs.GetConfigEnv())
	if err != nil {
		panic(err)
	}

	sugaredLogger, logger := loggers.InitLogger(conf.LogConfig)

	fmt.Println(conf.DBConfig)
	db, err := db2.GormInit(conf.DBConfig, db2.TableSlice, sugaredLogger)
	if err != nil {
		panic(err)
	}

	server, err := services.NewServer(services.WithConfig(conf),
		services.WithGinEngin(),
		services.WithGormDb(db),
		services.WithLog(logger),
		services.WithSuLog(sugaredLogger),
	)
	if err != nil {
		panic(err)
	}

	err = Start(server)
	if err != nil {
		panic(err)
	}
}

func Start(s *services.Server) error {
	//loading middleware
	s.GetGinEngine().Use(handlers.Cors())
	s.GetGinEngine().Use(loggers.GinLogger(s.GetLogger()))
	s.GetGinEngine().Use(loggers.GinRecovery(s.GetLogger(), true))
	s.GetGinEngine().GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg":  models.RESP_CODE_SUCCESS,
			"data": "Hello,World!",
		})
	})

	//loading route

	routers.LoadNoTokenRouter(s)

	routers.LoadUserRouter(s)

	routers.LoadManageDataRouter(s)

	s.GetGinEngine().Use(handlers.JWTAuthMiddleware(s))

	err := s.GetGinEngine().Run(":" + s.GetSeverPort())
	if err != nil {
		return err
	}

	return nil
}
