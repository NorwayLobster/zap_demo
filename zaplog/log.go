/*
 * @Date: 2022-02-24 11:53:49
 * @LastEditors: ChengWang
 * @LastEditTime: 2022-02-24 17:25:55
 * @FilePath: /zaplog/zaplog/log.go
 */
package zaplog

import (
	"fmt"
	"sync"

	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger
var logOnce sync.Once
var Out zapcore.WriteSyncer

// InitZapLog initializes Zap log setting once
func InitZapLog() {
	logOnce.Do(initZap)
	zap.Fields([]zap.Field{
		zap.String("game_name", "excalibur"),
		zap.String("app_id", "test appid"),
		zap.String("instance_id", "test INSTANCEID"),
		zap.String("app_name", "test APPNAME"),
	}...)
	// zap.Int("int", _tenInts[0]),
	// zap.Ints("ints", _tenInts),
	// zap.Strings("strings", _tenStrings),
	// zap.Time("time", _tenTimes[0]),
	// zap.Times("times", _tenTimes),
	// zap.Object("user1", _oneUser),
	// zap.Object("user2", _oneUser),
	// zap.Array("users", _tenUsers),
	// zap.Error(errExample),
}

// initZap initializes Zap
func initZap() {
	logLevel := viper.GetInt("Log.Level")
	logPath := viper.GetString("Log.LogPath")
	maxSize := viper.GetInt("Log.MaxSize")
	maxBackups := viper.GetInt("Log.MaxBackups")
	maxAge := viper.GetInt("Log.MaxAge")
	isCompressed := viper.GetBool("Log.IsCompressed")
	fmt.Println(logLevel)
	fmt.Println(logPath)
	fmt.Println(maxSize)
	fmt.Println(maxBackups)
	fmt.Println(maxAge)
	fmt.Println(isCompressed)
	initLogger(getLogLevel(logLevel), logPath, maxSize, maxBackups, maxAge, isCompressed)
	// InitLogger(zapcore.DebugLevel, "./zap_gin_rotation.log", 1, 10, 30, false)
}

// initLogger initializes log settings
func initLogger(level zapcore.Level, logPath string, maxSize, maxBackups, maxAge int, isCompressed bool) {
	encoder := getEncoder()
	Out = getLogWriter(logPath, maxSize, maxBackups, maxAge, isCompressed)
	core := zapcore.NewCore(encoder, Out, level)
	logger = zap.New(core, zap.AddCaller())
}

// getEncoder sets and gets encoder.
// which one is text encoder, i.e. output log in text?
func getEncoder() zapcore.Encoder {
	//set encoder config
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // ?????????????????????
	// encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) { enc.AppendString(t.Format("2006-01-02T15:04:05.000Z0700")) }
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // ??????????????????????????????????????????????????????
	// return zapcore.NewConsoleEncoder(encoderConfig)         //
	return zapcore.NewJSONEncoder(encoderConfig)
}

//getLogWriter sets and gets log writer with rotation
func getLogWriter(logPath string, maxSize, maxBackups, maxAge int, isCompressed bool) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logPath,      //"./zap_gin_rotation.log",
		MaxSize:    maxSize,      //1,
		MaxBackups: maxBackups,   //10,
		MaxAge:     maxAge,       //30,
		Compress:   isCompressed, // Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// getLogLevel gets log level
func getLogLevel(logLevel int) zapcore.Level {
	switch logLevel {
	case -1:
		return zapcore.DebugLevel
	case 1:
		return zapcore.WarnLevel
	case 2:
		return zapcore.ErrorLevel
	case 3:
		return zapcore.DPanicLevel
	case 4:
		return zapcore.PanicLevel
	case 5:
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}
