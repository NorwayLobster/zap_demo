/*
 * @Date: 2022-02-10 20:06:55
 * @LastEditors: ChengWang
 * @LastEditTime: 2022-02-21 11:53:31
 * @FilePath: /zaplog/example/main.go
 */
package main

import (
	"fmt"

	"github.com/NorwayLobster/zaplog"
	"go.uber.org/zap/zapcore"
	// "github.com/NorwayLobster/gomodone"
	// "github.com/NorwayLobster/moduletest"
	// cron "github.com/robfig/cron/v3"
	// moduletestV2 "github.com/NorwayLobster/moduletest/v2"
)

func main() {
	fmt.Printf("Hello world\n")
	fmt.Printf("Hello\n")
	// InitLogger(zapcore.DebugLevel)
	zaplog.InitLogger(zapcore.DebugLevel, 1, 10, 30, false)
	defer zaplog.Sync()
	startHTTPServer()
}
