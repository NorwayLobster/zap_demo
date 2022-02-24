package main

import "github.com/gin-gonic/gin"

func startHTTPServer() {
	r := gin.Default()
	r.GET("/ping", PongHandler)
	r.GET("/rankinglist/:name/:num", RankingListHandler)
	r.Run(":7080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
