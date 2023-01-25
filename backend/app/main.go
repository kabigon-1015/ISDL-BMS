package main

import (
	//Ginの実装
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	// "github.com/hoka-isdl/ISDL-BMS/backend/app/controller"
)

func main() {

	//Ginの変数
	router := gin.Default()

	//html/css/jsファイルのディレクトリをロード
	router.LoadHTMLGlob("../../frontend/dist/*.html")
	router.Static("/js", "../../frontend/dist/js")
	router.Static("/css", "../../frontend/dist/css")

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	router.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.html", gin.H{
			"judge": true,
		})
	})

	// //routerを渡す
	// controller.Router(router)

	//ポートを指定して実行
	router.Run(":80")
}
