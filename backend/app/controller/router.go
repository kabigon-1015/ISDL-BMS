package controller

import (
	//HTTPクライアントとサーバーの実装

	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"github.com/hoka-isdl/ISDL-BMS/backend/app/repository"
	"github.com/hoka-isdl/ISDL-BMS/backend/app/structure"
)

var session_user structure.Session
var login_user structure.User

func checkSession(c *gin.Context) {

	session := sessions.Default(c)
	session_user.Name = session.Get("user_name")

	if session_user.Name == nil {
		c.Redirect(302, "/login")
		c.Abort()
	} else {
		c.Set("user_name", session_user.Name)
		login_user.Name = session_user.Name.(string)
		c.Next()
	}
}

func Router(router *gin.Engine) {

	//session
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	repository.CreateTask()
	router.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.html", gin.H{
			"judge": true,
		})
	})

	router.POST("/research", func(c *gin.Context) {
		isbn := c.PostForm("isbn")
		result := repository.Researchbook(isbn)
		c.JSON(200, result)
	})
}
