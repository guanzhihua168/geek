package main

import (
	"context"
	"errors"
	"flag"
	"github.com/DeanThompson/ginpprof"
	"github.com/SkyAPM/go2sky"
	httpPlugin "github.com/SkyAPM/go2sky/plugins/http"
	"github.com/SkyAPM/go2sky/reporter"
	"parent-api-go/pkg/redis"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"os"
	"os/signal"
	"parent-api-go/global"
	"parent-api-go/internal/model"
	"parent-api-go/internal/routers"
	"parent-api-go/pkg/linkerd"
	"parent-api-go/pkg/logger"
	"parent-api-go/pkg/setting"
	"parent-api-go/pkg/tracer"
	"parent-api-go/pkg/validator"
	"strings"
	"syscall"
	"time"
)

var (
	port      string
	runMode   string
	config    string
	isVersion bool
)

//初始化各项组件
func init() {
	err := setupFlag()
	if err != nil {
		log.Fatalf("init.setupFlag err %v", err)
	}
	err = setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err %v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	err = setupRedisPoolEngine()
	if err != nil {
		log.Fatalf("init.setupRedisPoolEngine err: %v", err)
	}

	err = setupValidator()
	if err != nil {
		log.Fatalf("init.setupValidator err: %v", err)
	}
	err = setupSkywalk()
	if err != nil {
		log.Fatalf("init.setupSkywalk err %v", err)
	}
	//err = setupTracer()
	//if err != nil {
	//	log.Fatalf("init.setupTracer err %v", err)
	//}

}

func main() {

	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	ginpprof.Wrap(router)
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("s.ListenAndServe err: %v", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shuting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")
}

//配置项目cli启动参数
func setupFlag() error {
	flag.StringVar(&port, "port", "", "启动端口")
	flag.StringVar(&runMode, "mode", "", "启动模式")
	flag.StringVar(&config, "config", "configs/", "指定使用配置文件的路径")
	flag.BoolVar(&isVersion, "version", false, "编译信息")
	flag.Parse()

	return nil
}

//启动项目配置初始化
func setupSetting() error {
	s, err := setting.NewSetting(strings.Split(config, ",")...)
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Databases.oldBoss.master", &global.DatabaseOldBossMasterSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Databases.oldBoss.slave", &global.DatabaseOldBossSlaveSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("Databases.parent.master", &global.DatabaseParentMasterSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("Databases.parent.slave", &global.DatabaseParentSlaveSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("Databases.mop.master", &global.DatabaseMopMasterSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("Databases.mop.slave", &global.DatabaseMopSlaveSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("Linkerd", &global.LinkerdSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("skywalking", &global.SkywalkingSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("rabbitmq", &global.RabbitmqSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Redis.Default", &global.RedisDefaultSetting)
	if err != nil {
		return err
	}

	global.AppSetting.DefaultContextTimeout *= time.Second
	global.JWTSetting.Expire *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	if port != "" {
		global.ServerSetting.HttpPort = port
	}

	if runMode != "" {
		global.ServerSetting.RunMode = runMode
	}

	return nil
}

//日志初始化
func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName +
		global.AppSetting.LogFileExt

	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   500,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}

//初始化数据库engine
func setupDBEngine() error {
	var err error

	global.OldBossMasterDBEngine, err = model.NewDBEngine(global.DatabaseOldBossMasterSetting)
	if err != nil {
		return err
	}
	global.OldBossSlaveDBEngine, err = model.NewDBEngine(global.DatabaseOldBossSlaveSetting)
	if err != nil {
		return err
	}

	global.ParentMasterDBEngine, err = model.NewDBEngine(global.DatabaseParentMasterSetting)
	if err != nil {
		return err
	}
	global.ParentSlaveDBEngine, err = model.NewDBEngine(global.DatabaseParentSlaveSetting)
	if err != nil {
		return err
	}
	global.MopMasterDBEngine, err = model.NewDBEngine(global.DatabaseMopMasterSetting)
	if err != nil {
		return err
	}
	global.MopSlaveDBEngine, err = model.NewDBEngine(global.DatabaseMopSlaveSetting)
	if err != nil {
		return err
	}

	return nil
}

//初始化验证组件
func setupValidator() error {
	global.Validator = validator.NewCustomValidator()
	global.Validator.Engine()
	binding.Validator = global.Validator

	return nil
}

//初始化jaeger组件
func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer("parent-api-go", "127.0.0.1:6831")
	if err != nil {
		return err
	}

	global.Tracer = jaegerTracer
	return nil
}

func setupSkywalk() error {
	linkerd.LinkerdRequestClient = &http.Client{}
	if global.SkywalkingSetting.ServerHost == "" {
		return errors.New("ServerHost is empty")
	}

	// skewalking 初始化
	logs := logrus.WithFields(logrus.Fields{
		"sky address": global.SkywalkingSetting.ServerHost,
		"server name": global.SkywalkingSetting.ServerName,
	})
	if sk, e := reporter.NewGRPCReporter(global.SkywalkingSetting.ServerHost); e == nil {
		global.SKWalking = sk
		global.SKWTrace, e = go2sky.NewTracer(global.SkywalkingSetting.ServerName, go2sky.WithReporter(global.SKWalking))
		if e != nil {
			logs.Error("new trace fail:" + e.Error())
			return e
		} else {
			logs.Info("skewalking info init success")
		}
	} else {
		logs.Error("connection fail:" + e.Error())
		return e
	}

	if global.SKWTrace != nil {
		linkerd.LinkerdRequestClient, _ = httpPlugin.NewClient(global.SKWTrace)
		logs.Info("linkerd request client init use sky")
	}
	return nil
}

func setupRedisPoolEngine() error {
	var err error
	global.RedisPoolDefaultEngine, err = model.NewRedisEngine(global.RedisDefaultSetting)
	if err != nil {
		return err
	}
	redis.RedisPrefix = global.AppSetting.Name

	return nil
}
