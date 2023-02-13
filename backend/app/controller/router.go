package controller

import (
	//HTTPクライアントとサーバーの実装

	"fmt"
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"github.com/hoka-isdl/ISDL-BMS/backend/app/repository"
	"github.com/hoka-isdl/ISDL-BMS/backend/app/structure"
)

var session_user structure.Session
var login_user structure.Students

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

	router.POST("/bookinfo", func(c *gin.Context) {
		isbn := c.PostForm("isbn")
		fmt.Print(isbn)
		title, author, publisher := repository.GetBookinfo(isbn)
		fmt.Print(title, author, publisher)
		c.JSON(200, gin.H{
			"title":     title,
			"author":    author,
			"publisher": publisher,
		})
	})

	router.POST("/login", func(c *gin.Context) {
		id := c.PostForm("id")
		password := c.PostForm("password")
		fmt.Print(id)
		name, email := repository.GetUserinfo(id, password)
		fmt.Print(name, email)
		c.JSON(200, gin.H{
			"name":     name,
			"email":    email,
		})
	})

	router.POST("/filterbook", func(c *gin.Context) {
		var titlelist []string
		var authorlist []string
		var publisherlist []string
		var rentaluser_namelist []string
		tags := c.PostForm("tags")
		fmt.Print(tags)
		booklist_temp := repository.GetAllBookData()

		for _, v := range booklist_temp {
			title := v[0]
			author := v[1]
			publisher := v[2]
			rentaluser_name := repository.GetRenterInfo(v[3])

			fmt.Print(title, author, publisher, rentaluser_name, "\n")

			titlelist = append(titlelist, title)
			authorlist = append(authorlist, author)
			publisherlist = append(publisherlist, publisher)
			rentaluser_namelist = append(rentaluser_namelist, rentaluser_name)
		}
		c.JSON(200, gin.H{
			"title":           titlelist,
			"author":          authorlist,
			"publisher":       publisherlist,
			"rentaluser_name": rentaluser_namelist,
		})
	})

	router.POST("/return_register", func(c *gin.Context) {
		isbns := c.PostForm("isbn")
		isbn := strings.Split(isbns, ",")
		// fmt.Print(isbn)
		if len(isbn) != 0 {
			for _, v := range isbn {
				repository.Returnbook("1116190020", v)
			}
		}
		c.JSON(200, gin.H{
			"message": "返却完了",
		})
	})

	router.POST("/book_list", func(c *gin.Context) {
		var titlelist []string
		var authorlist []string
		var publisherlist []string
		var rentaluser_namelist []string

		booklist_temp := repository.GetAllBookData()

		for _, v := range booklist_temp {
			title := v[0]
			author := v[1]
			publisher := v[2]
			rentaluser_name := repository.GetRenterInfo(v[3])

			fmt.Print(title, author, publisher, rentaluser_name, "\n")

			titlelist = append(titlelist, title)
			authorlist = append(authorlist, author)
			publisherlist = append(publisherlist, publisher)
			rentaluser_namelist = append(rentaluser_namelist, rentaluser_name)
		}
		c.JSON(200, gin.H{
			"title":           titlelist,
			"author":          authorlist,
			"publisher":       publisherlist,
			"rentaluser_name": rentaluser_namelist,
		})
	})

	router.POST("/book_detail", func(c *gin.Context) {

		booktitle := c.PostForm("booktitle")
		fmt.Print(booktitle)

		bookauthor, bookpublisher, bookitem_caption, bookimageurl := repository.GetBookDetail(booktitle)

		c.JSON(200, gin.H{
			"title":        booktitle,
			"author":       bookauthor,
			"publisher":    bookpublisher,
			"item_caption": bookitem_caption,
			"imageyrl":     bookimageurl,
		})
	})
}
