package main

import (
	//Ginの実装
	"github.com/gin-gonic/gin"
	
	"github.com/hoka-isdl/ISDL-BMS/backend/app/controller"
)

func main() {

	//Ginの変数
	router := gin.Default()

	//html/css/jsファイルのディレクトリをロード
	router.LoadHTMLGlob("../../frontend/dist/*.html")
	router.Static("/js", "../../frontend/dist/js")
	router.Static("/css", "../../frontend/dist/css")
	router.Static("/img", "../../frontend/dist/img")

	// //routerを渡す
	controller.Router(router)

	//ポートを指定して実行
	router.Run(":80")
}
