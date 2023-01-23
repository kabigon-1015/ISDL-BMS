package controller

import (
	//HTTPクライアントとサーバーの実装

	"fmt"
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

	router.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.html", gin.H{
			"judge": true,
		})
	})

	router.POST("/login", func(c *gin.Context) {

		user_name := c.PostForm("name")
		password := c.PostForm("password")

		if repository.VerifybyName(user_name, password) {
			session := sessions.Default(c)
			session.Set("user_name", user_name)
			session.Save()

			c.Redirect(302, "/taskapp")
		} else {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"judge": false,
			})
		}
	})

	router.GET("/taskapp", func(c *gin.Context) {

		var user_list []string
		user_list = repository.GetUserList()
		checkSession(c)

		todo_task_list, doing_task_list, done_task_list := repository.GetTasks(login_user.Name)

		c.HTML(200, "index.html", gin.H{
			"user_name":       login_user.Name,
			"user_list":       user_list,
			"todo_task_list":  todo_task_list,
			"doing_task_list": doing_task_list,
			"done_task_list":  done_task_list,
		})
	})

	router.GET("taskapp/settings", func(c *gin.Context) {

		c.HTML(http.StatusOK, "settings.html", gin.H{
			"user_name": login_user.Name,
		})
	})
	router.GET("/changepass", func(c *gin.Context) {

		c.HTML(http.StatusOK, "changepass.html", gin.H{
			"user_name": login_user.Name,
			"judge":     true,
		})
	})

	router.POST("/changepass", func(c *gin.Context) {

		oldpass := c.PostForm("oldpass")
		newpass := c.PostForm("newpass")
		oldpass_judge := false
		changepass_successful := false

		if repository.Verify(login_user.Name, oldpass) {
			repository.ChangePass(login_user.Name, newpass)
			changepass_successful = true
		} else {
			oldpass_judge = true
		}

		c.HTML(http.StatusOK, "changepass.html", gin.H{
			"user_name":             login_user.Name,
			"oldpass_judge":         oldpass_judge,
			"changepass_successful": changepass_successful,
		})
	})

	router.POST("/create_a_task", func(c *gin.Context) {

		task_name := c.PostForm("task_name")
		description := c.PostForm("task_description")
		pic := c.PostForm("task_pic")
		deadline := c.PostForm("task_deadline")

		repository.CreateTask(task_name, description, pic, deadline)

		c.Redirect(302, "/taskapp")
	})

	router.POST("/update_todo", func(c *gin.Context) {

		taskName := c.PostForm("taskName")
		taskDescription := c.PostForm("taskDescription")

		repository.UpdateTaskStatus(login_user.Name, taskName, taskDescription, "todo")

		fmt.Println("\n[" + login_user.Name + "]")
		fmt.Println("Status of \"" + taskName + "-" + taskDescription + "\" is changed into \"todo\"\n")
	})

	router.POST("/update_doing", func(c *gin.Context) {

		taskName := c.PostForm("taskName")
		taskDescription := c.PostForm("taskDescription")

		repository.UpdateTaskStatus(login_user.Name, taskName, taskDescription, "doing")

		fmt.Println("\n[" + login_user.Name + "]")
		fmt.Println("Status of \"" + taskName + "-" + taskDescription + "\" is changed into \"doing\"\n")
	})

	router.POST("/update_done", func(c *gin.Context) {

		taskName := c.PostForm("taskName")
		taskDescription := c.PostForm("taskDescription")

		repository.UpdateTaskStatus(login_user.Name, taskName, taskDescription, "done")

		fmt.Println("\n[" + login_user.Name + "]")
		fmt.Println("Status of \"" + taskName + "-" + taskDescription + "\" is changed into \"done\"\n")
	})

	router.POST("/delete_task", func(c *gin.Context) {

		taskName := c.PostForm("taskName")
		taskDescription := c.PostForm("taskDescription")

		fmt.Println(taskDescription)

		repository.DeleteTask(login_user.Name, taskName, taskDescription)
	})
}
