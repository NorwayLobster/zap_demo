/*
 * @Date: 2022-02-24 15:28:58
 * @LastEditors: ChengWang
 * @LastEditTime: 2022-02-24 15:29:00
 * @FilePath: /zaplog/example/gin_demo.go
 */
package main

import "github.com/gin-gonic/gin"

func startHTTPServer() {
	r := gin.Default()
	r.GET("/ping", PongHandler)
	r.GET("/rankinglist/:name/:num", RankingListHandler)
	r.Run(":7080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
