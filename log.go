/*
 * @Date: 2022-02-15 16:30:28
 * @LastEditors: ChengWang
 * @LastEditTime: 2022-02-23 17:50:37
 * @FilePath: /zaplog/log.go
 */

package zaplog

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//Q1: logger.Info() 是否线程安全？ 答：应该是
//Q2: server一般不意外退出，无法执行logger.Sync(), 日志是否能及时刷出 答：能
//Q3:  模仿lugrus,把logger封装到包里面，每次get即可，且只能初始化一次
//git repo:  sandipb/zap-examples
var sugarLogger *zap.SugaredLogger
var logger *zap.Logger

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}
func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}
func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}
func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}
func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}

func Sync() {
	logger.Sync()
}

func InitLogger(level zapcore.Level, maxSize, maxBackups, maxAge int, isCompressed bool) {
	//TODO(chengwang):  init once
	encoder := getEncoder()
	writeSyncer := getLogWriter(maxSize, maxBackups, maxAge, isCompressed)
	core := zapcore.NewCore(encoder, writeSyncer, level)
	logger = zap.New(core, zap.AddCaller())
	// sugarLogger = logger.Sugar()
}

func InitLogger1() {
	// logger, _ = zap.NewProduction()
	// logger, _ = zap.NewDevelopment()
	// logger = zap.NewExample()
	// logger = zap.NewNop()
	// logger = zap.New(core, options)
}

//which one is text encoder, i.e. output log in text?
func getEncoder() zapcore.Encoder {
	//set encoder config
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // 修改时间编码器
	// encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	// enc.AppendString(t.Format("2006-01-02T15:04:05.000Z0700"))
	// }
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // 在日志文件中使用大写字母记录日志级别
	// return zapcore.NewConsoleEncoder(encoderConfig)         //
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getEncoder1() zapcore.Encoder { //
	// zapcore.NewJSONEncoder()
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

func getLogWriter(maxSize, maxBackups, maxAge int, isCompressed bool) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./zap_gin_rotation.log",
		MaxSize:    maxSize,      //1,
		MaxBackups: maxBackups,   //10,
		MaxAge:     maxAge,       //30,
		Compress:   isCompressed, // Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

/*
Lumberjack Logger采用以下属性作为输入:
Filename: 日志文件的位置
MaxSize：在进行切割之前，日志文件的最大大小（以MB为单位）
MaxBackups：保留旧文件的最大个数
MaxAges：保留旧文件的最大天数
Compress：是否压缩/归档旧文件
*/

func getLogWriter1() zapcore.WriteSyncer {
	file, _ := os.Create("./zap_test.log")
	return zapcore.AddSync(file)
}

/*
func zap_demo1() {
	file1, err := os.OpenFile("./access.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	file2, err := os.OpenFile("./error.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	var tops = []log.TeeOption{
		{
			W: file1,
			Lef: func(lvl log.Level) bool {
				return lvl <= log.InfoLevel
			},
		},
		{
			W: file2,
			Lef: func(lvl log.Level) bool {
				return lvl > log.InfoLevel
			},
		},
	}

	logger := log.NewTee(tops)
	log.ResetDefault(logger)

	log.Info("demo3:", log.String("app", "start ok"),
		log.Int("major version", 3))
	log.Error("demo3:", log.String("app", "crash"),
		log.Int("reason", -1))
}

*/
