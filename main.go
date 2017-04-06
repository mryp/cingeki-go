package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/mryp/cingeki-go/config"
)

func main() {
	//環境設定読み込み
	if !config.LoadConfig() {
		log.Println("設定ファイル読み込み失敗（デフォルト値動作）")
	}

	//ECHO初期設定
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORS()) //CORS対応（他ドメインからAJAX通信可能にする）
	if config.GetConfig().Log.Output == "stream" {
	}
	switch config.GetConfig().Log.Output {
	case "stream":
		e.Use(middleware.Logger())
	case "file":
		//未実装
	}

	//ルーティング
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "cingeki-go")
	})
	apiGroup := e.Group("/api")
	apiGroup.POST("/regist", RegistHandler)
	apiGroup.POST("/regist/matome", RegistMatomeHandler)
	apiGroup.GET("/story/:number", StoryHandler)
	apiGroup.GET("/image/:number", ImageHandler)

	//開始
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(config.GetConfig().Server.PortNum)))
}
