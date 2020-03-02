package logging

import (
    "time"
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
    "github.com/lestrrat-go/file-rotatelogs"
    "github.com/rifflock/lfshook"
    "vblog/common/setting"
    "fmt"
)

type Level int

var (
    logger = logrus.New()
)

// 日志记录到文件
func LogToFile() gin.HandlerFunc {

    //处理日志写入文件
    src, err := openLogFile()
    if err != nil {
        fmt.Println("err", err)
    }

    //设置输出
    logger.Out = src

    //设置日志级别
    logger.SetLevel(setting.LogLevel)

    // 设置 rotatelogs
    logWriter, err := rotatelogs.New(
        // 分割后的文件名称
        fullLogFilePath + ".%Y%m%d.txt",

        // 生成软链，指向最新日志文件
        rotatelogs.WithLinkName(fullLogFilePath),

        // 设置最大保存时间(7天)
        rotatelogs.WithMaxAge(7*24*time.Hour),

        // 设置日志切割时间间隔(1天)
        rotatelogs.WithRotationTime(24*time.Hour),
    )

    writeMap := lfshook.WriterMap{
        logrus.InfoLevel:  logWriter,
        logrus.FatalLevel: logWriter,
        logrus.DebugLevel: logWriter,
        logrus.WarnLevel:  logWriter,
        logrus.ErrorLevel: logWriter,
        logrus.PanicLevel: logWriter,
    }
    lfHook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
        TimestampFormat:"2006-01-02 15:04:05",
    })

    // 新增 Hook
    logger.AddHook(lfHook)

    //设置日志格式
    logger.SetFormatter(&logrus.TextFormatter{
        TimestampFormat:"2006-01-02 15:04:05",
    })

    return func(c *gin.Context) {
        // 开始时间
        startTime := time.Now()

        // 处理请求
        c.Next()

        // 结束时间
        endTime := time.Now()

        // 执行时间
        latencyTime := endTime.Sub(startTime)

        // 请求方式
        reqMethod := c.Request.Method

        // 请求路由
        reqUri := c.Request.RequestURI

        // 状态码
        statusCode := c.Writer.Status()

        // 请求IP
        clientIP := c.ClientIP()

        // 日志格式
        logger.WithFields(logrus.Fields{
            "code"  : statusCode,
            "latency" : latencyTime,
            "client"    : clientIP,
            "method"   : reqMethod,
            "uri"      : reqUri,
        }).Info()
    }
}

func Debug(v ...interface{}) {
    logger.Debug(v)
}

func Info(v ...interface{}) {
    logger.Info(v)
}

func Warn(v ...interface{}) {
    logger.Warn(v)
}

func Error(v ...interface{}) {
    logger.Error(v)
}

func Fatal(v ...interface{}) {
    logger.Fatal(v)
}
