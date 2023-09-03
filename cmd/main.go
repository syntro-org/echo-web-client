package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/syntro-org/echo-web-client/internal/handler"
	"github.com/syntro-org/echo-web-client/pkg/util"
)

func main() {
	gin.SetMode(gin.DebugMode)

	var p util.AppParams
	err := util.ReadParams("configs/config.yml", &p)
	if err != nil {
		logrus.Fatalf("config file read error: %s", err.Error())
	}

	handlers := handler.NewHandler(p.BaseAPIUrl, p.CookieSecrets)

	var server http.Server
	{
		server.Addr = util.MakeServerAddr(p.AppMode, p.Port)
		server.Handler = handlers.Init()
	}

	err = server.ListenAndServe()
	if err != nil {
		logrus.Fatalf("server startup error: %s", err.Error())
	}
}
