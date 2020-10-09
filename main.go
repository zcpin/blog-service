package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zcpin/blog-service/global"
	"github.com/zcpin/blog-service/internal/model"
	"github.com/zcpin/blog-service/internal/routers"
	"github.com/zcpin/blog-service/pkg/logger"
	"github.com/zcpin/blog-service/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

// @title 博客系统
// @version 1.0
// @description Go 博客系统
// @termsOfService https://github.com/zcpin/blog-service
func main() {
	gin.SetMode(global.ServerSetting.RunModel)
	router := routers.NewRouter()
	serve := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	err := serve.ListenAndServe()
	if err != nil {
		log.Fatalf("服务启动失败:%v\n", err)
	}
}

func init() {

	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err %v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err %v", err)
	}
}

func setupSetting() error {
	settings, err := setting.NewSetting()

	if err != nil {
		return err
	}
	err = settings.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = settings.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = settings.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	err = settings.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	global.JWTSetting.Expire *= time.Second

	err = settings.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	global.DBEngine.LogMode(true)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}
