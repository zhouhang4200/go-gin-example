package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhouhang4200/go-gin-example/models"
	"github.com/zhouhang4200/go-gin-example/pkg/gredis"
	"github.com/zhouhang4200/go-gin-example/pkg/logging"
	"github.com/zhouhang4200/go-gin-example/pkg/setting"
	"github.com/zhouhang4200/go-gin-example/routers"
	"log"
	"net/http"
)

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	_ = gredis.Setup()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	_ = server.ListenAndServe()
}
