package controller

import (
	//HTTPクライアントとサーバーの実装

	"fmt"
	"net/http"
	"strconv"
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
	// repository.CreateTask_tag()
	// repository.CreateTask2()
	// repository.InsertBooks()
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
		name, email := repository.VeryfyUser(id, password)
		fmt.Print(name, email)
		if len(name) == 0 {
			c.JSON(200, gin.H{})
		} else {
			c.JSON(200, gin.H{
				"name":  name,
				"email": email,
			})
		}

	})

	router.POST("/signup", func(c *gin.Context) {
		userid := c.PostForm("userid")
		password := c.PostForm("password")
		username := c.PostForm("username")
		emailadress := c.PostForm("emailadress")

		repository.SignUp(userid, password, username, emailadress)

		c.JSON(200, gin.H{
			"name": username,
			"id":   userid,
		})
	})

	router.POST("/filterbook", func(c *gin.Context) {
		var titlelist []string
		var authorlist []string
		var publisherlist []string
		var rentaluser_namelist []string
		var tagid []string

		_len := c.PostForm("taglength")
		len, _ := strconv.Atoi(_len)
		for i := 0; i < len; i++ {
			// tag := append(tags, c.PostForm("tag[" + strconv.Itoa(i) + "]"))
			tagid = append(tagid, repository.GetTagid(c.PostForm("tag["+strconv.Itoa(i)+"]")))
			// fmt.Print(c.PostForm("tag[" + strconv.Itoa(i) + "]"))
		}
		// fmt.Print(tagid)
		booklist_temp := repository.FilterBooks_ver2(tagid)

		for _, v := range booklist_temp {
			title := v[0]
			author := v[1]
			publisher := v[2]
			rentaluser_name := repository.GetRenterInfo(v[3])

			// fmt.Print(title, author, publisher, rentaluser_name, "\n")

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

	router.POST("/addbooktag", func(c *gin.Context) {
		var tagid []string

		isbn := c.PostForm("isbn")
		_len := c.PostForm("taglength")
		len, _ := strconv.Atoi(_len)
		for i := 0; i < len; i++ {
			// tag := append(tags, c.PostForm("tag[" + strconv.Itoa(i) + "]"))
			tagid = append(tagid, repository.GetTagid(c.PostForm("tag["+strconv.Itoa(i)+"]")))
		}

		repository.AddBookTag(tagid, isbn)
		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	router.POST("/userinfo", func(c *gin.Context) {
		var hist_titlelist []string
		var hist_authorlist []string
		var hist_publisherlist []string
		var rented_titlelist []string
		var rented_authorlist []string
		var rented_publisherlist []string

		userid := c.PostForm("user_id")

		renthist_books, rented_books := repository.GetUserInfo(userid)

		for _, v := range renthist_books {
			hist_title := v[0]
			hist_author := v[1]
			hist_publisher := v[2]

			// fmt.Print(hist_title, hist_author, hist_publisher, "\n")

			hist_titlelist = append(hist_titlelist, hist_title)
			hist_authorlist = append(hist_authorlist, hist_author)
			hist_publisherlist = append(hist_publisherlist, hist_publisher)
		}
		for _, v := range rented_books {
			rented_title := v[0]
			rented_author := v[1]
			rented_publisher := v[2]

			fmt.Print(rented_title, rented_author, rented_publisher, "\n")

			rented_titlelist = append(rented_titlelist, rented_title)
			rented_authorlist = append(rented_authorlist, rented_author)
			rented_publisherlist = append(rented_publisherlist, rented_publisher)
		}
		c.JSON(200, gin.H{
			"title":            hist_titlelist,
			"author":           hist_authorlist,
			"publisher":        hist_publisherlist,
			"rented_title":     rented_titlelist,
			"rented_author":    rented_authorlist,
			"rented_publisher": rented_publisherlist,
		})
	})

	router.POST("/rental_register", func(c *gin.Context) {
		userid := c.PostForm("user_id")
		isbns := c.PostForm("isbn")
		isbn := strings.Split(isbns, ",")
		// fmt.Print(userid)
		message := " "
		// fmt.Print(isbn)
		if userid == "null" {
			c.JSON(200, gin.H{
				"message": "ログインしてください",
			})
		} else {
			if len(isbn) != 0 {
				for _, v := range isbn {
					message = repository.Rentbook(userid, v)
				}
			}
			c.JSON(200, gin.H{
				"message": message,
			})
		}
	})

	router.POST("/return_register", func(c *gin.Context) {
		userid := c.PostForm("user_id")
		isbns := c.PostForm("isbn")
		isbn := strings.Split(isbns, ",")
		fmt.Print(userid)
		if userid == "null" {
			c.JSON(200, gin.H{
				"message": "ログインしてください",
			})
		} else {
			if len(isbn) != 0 {
				for _, v := range isbn {
					repository.Returnbook(userid, v)
				}
			}
			c.JSON(200, gin.H{
				"message": "返却完了",
			})
		}
	})

	router.POST("/book_list", func(c *gin.Context) {
		var titlelist []string
		var authorlist []string
		var publisherlist []string
		var rentaluser_namelist []string
		var isbnlist []string

		booklist_temp := repository.GetAllBookData()

		for _, v := range booklist_temp {
			title := v[0]
			isbn := v[1]
			author := v[2]
			publisher := v[3]
			rentaluser_name := repository.GetRenterInfo(v[4])

			// fmt.Print(title, isbn, author, publisher, rentaluser_name, "\n")

			titlelist = append(titlelist, title)
			isbnlist = append(isbnlist, isbn)
			authorlist = append(authorlist, author)
			publisherlist = append(publisherlist, publisher)
			rentaluser_namelist = append(rentaluser_namelist, rentaluser_name)
		}
		c.JSON(200, gin.H{
			"title":           titlelist,
			"isbn":            isbnlist,
			"author":          authorlist,
			"publisher":       publisherlist,
			"rentaluser_name": rentaluser_namelist,
		})
	})

	router.POST("/book_detail", func(c *gin.Context) {

		booktitle := c.PostForm("booktitle")
		// fmt.Print(booktitle)

		bookauthor, bookpublisher, bookitem_caption, bookimageurl := repository.GetBookDetail(booktitle)

		c.JSON(200, gin.H{
			"title":        booktitle,
			"author":       bookauthor,
			"publisher":    bookpublisher,
			"item_caption": bookitem_caption,
			"imageurl":     bookimageurl,
		})
	})
}
