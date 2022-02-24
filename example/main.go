/*
 * @Date: 2022-02-10 20:06:55
 * @LastEditors: ChengWang
 * @LastEditTime: 2022-02-24 15:39:53
 * @FilePath: /zaplog/example/main.go
 */
package main

import (
	"fmt"

	"github.com/NorwayLobster/zap_demo/zaplog"

	// "github.com/NorwayLobster/gomodone"
	// "github.com/NorwayLobster/moduletest"
	// cron "github.com/robfig/cron/v3"
	// moduletestV2 "github.com/NorwayLobster/moduletest/v2"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var configFile = pflag.String("configfile", "./server.toml", "config file")

// SetConfig 设置配置
func SetConfig(configPath string) {
	viper.SetConfigFile(configPath)
	viper.SetConfigType("toml")
	if err := viper.ReadInConfig(); err != nil {
		panic("加载配置文件失败: " + configPath + ", err: " + err.Error())
	}
}

func main() {
	fmt.Printf("Hello world\n")
	fmt.Printf("Hello\n")
	pflag.Parse()
	SetConfig(*configFile)
	zaplog.InitZapLog()
	// InitLogger(zapcore.DebugLevel)
	// zaplog.InitLogger(zapcore.DebugLevel, 1, 10, 30, false)
	defer zaplog.Sync()

	startHTTPServer()
}
