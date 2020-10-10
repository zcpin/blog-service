package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zcpin/blog-service/global"
	"github.com/zcpin/blog-service/internal/model"
	"github.com/zcpin/blog-service/internal/routers"
	"github.com/zcpin/blog-service/pkg/logger"
	"github.com/zcpin/blog-service/pkg/setting"
	"github.com/zcpin/blog-service/pkg/tracer"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"strings"
	"time"
)

// @title 博客系统
// @version 1.0
// @description Go 博客系统
// @termsOfService https://github.com/zcpin/blog-service
func main() {
	if isVersion {
		fmt.Printf("build_time: %s\n", buildTime)
		fmt.Printf("build_version: %s\n", buildVersion)
		fmt.Printf("git_commit_id: %s\n", gitCommitID)
		return
	}
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

var (
	port    string
	runMode string
	config  string

	isVersion    bool
	buildTime    string
	buildVersion string
	gitCommitID  string
)

func init() {
	err := setupFlag()
	if err != nil {
		log.Fatalf("init.setupFlag err: %v", err)
	}

	err = setupSetting()
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

	err = setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err: %v", err)
	}
}

func setupSetting() error {
	settings, err := setting.NewSetting(strings.Split(config, ",")...)

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

	if port != "" {
		global.ServerSetting.HttpPort = port
	}
	if runMode != "" {
		global.ServerSetting.RunModel = runMode
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

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer("blog-service", "127.0.0.1:6831")
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}

func setupFlag() error {
	flag.StringVar(&port, "port", "", "启动端口")
	flag.StringVar(&runMode, "mode", "", "启动模式")
	flag.StringVar(&config, "config", "configs/", "指定要使用的配置文件路径")
	flag.BoolVar(&isVersion, "version", false, "编译信息")
	flag.Parse()
	return nil
}
