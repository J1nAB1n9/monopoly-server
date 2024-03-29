package logger

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"monopoly-server/utils"
	"time"
)

var lineEnding string

var zapLogger *zap.Logger
var sugarLogger *zap.SugaredLogger

func getWriter(filename string) io.Writer {
	// 生成rotatelogs的Logger 实际生成的文件名 demo.log.YYmmddHH
	// demo.log是指向最新日志的链接
	// 保存7天内的日志，每1小时(整点)分割一次日志
	hook, err := rotatelogs.New(
		filename+"_%Y%m%d%H.log", // 没有使用go风格反人类的format格式
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour),
	)

	if err != nil {
		panic(err)
	}
	return hook
}

func NewLogger() {
	lineEnding = utils.GetSystemLineEnding(utils.GetCurrentOS())

	// 设置一些基本日志格式 具体含义还比较好理解，直接看zap源码也不难懂
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder,
		TimeKey:     "ts",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		CallerKey:    "file",
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})
	zap.AddCaller()

	// 实现两个判断日志等级的interface (其实 zapcore.*Level 自身就是 interface)
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.WarnLevel
	})

	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel
	})

	// 获取 info、warn日志文件的io.Writer 抽象 getWriter() 在下方实现
	infoWriter := getWriter("./log/serverLog")
	warnWriter := getWriter("./log/serverError")

	// 最后创建具体的Logger
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel,),
		zapcore.NewCore(encoder, zapcore.AddSync(warnWriter), warnLevel),
	)

	zapLogger = zap.New(core, zap.AddCaller(),zap.AddCallerSkip(1))  // 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数
	sugarLogger = zapLogger.Sugar()
}

func GetSugarLogger() *zap.SugaredLogger {
	return sugarLogger
}

func Sync() {
	sugarLogger.Sync()
	zapLogger.Sync()
}

func Debug(template string, args ...interface{}) {
	sugarLogger.Debugf(template,args...)
}

func Info(template string, args ...interface{}) {
	sugarLogger.Infof(template,args...)
}

func  Warm(template string, args ...interface{}){
	sugarLogger.Warnf(template,args...)
}

func  Error(template string, args ...interface{}){
	sugarLogger.Errorf(template,args...)
}

func DPanic(template string, args ...interface{}) {
	sugarLogger.DPanicf(template,args...)
}

func Panic(template string, args ...interface{}){
	sugarLogger.Panicf(template,args...)
}

func Fatal(template string, args ...interface{}){
	sugarLogger.Fatalf(template,args...)
}