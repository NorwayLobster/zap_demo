/*
 * @Date: 2022-02-24 11:53:00
 * @LastEditors: ChengWang
 * @LastEditTime: 2022-02-24 11:53:11
 * @FilePath: /zaplog/wrapper.go
 */

package zlog

import "go.uber.org/zap"

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

func AddContext(fields ...zap.Field) {
	logger = logger.With(fields...)
}

func Sync() {
	logger.Sync()
}
