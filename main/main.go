package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Login struct {
	User string `form:"username" json:"user" uri:"user" binding:"required"`
	Pwd  string `form:"password" json:"pwd" uri:"pwd" binding:"required"`
}

//func login(c *gin.Context){
//	name := c.DefaultQuery("name","defaultName")
//	c.String(http.StatusOK,fmt.Sprintf("login success!\nhello %v~",name))
//	fmt.Println(name)
//}
//
//func submit1(c *gin.Context){
//	username := c.PostForm("username")
//	password := c.PostForm("password")
//	c.JSON(http.StatusOK, gin.H{
//		"username": username,
//		"password": password,
//	})
//}
//
//func submit2(c *gin.Context){
//	username := c.PostForm("username")
//	password := c.PostForm("password")
//	c.JSON(http.StatusOK, gin.H{
//		"username": username,
//		"password": password,
//	})
//}

func main() {

	//r := gin.Default()
	//
	//r.GET("/hello", func(c *gin.Context) {
	//	// c.JSON：返回JSON格式的数据
	//	c.JSON(200, gin.H{
	//		"message": "Hello world!",
	//		"data":    "gin go go ",
	//	})
	//})
	//
	//r.POST("/book", func(c *gin.Context) {
	//
	//	username := c.PostForm("username")
	//	password := c.PostForm("password")
	//	c.GetRawData()
	//
	//	c.JSON(200, gin.H{
	//		"message":  "this is POST method",
	//		"username": username,
	//		"password": password,
	//	})
	//})
	//
	//r.GET("/urlQuery", func(c *gin.Context) {
	//	name := c.DefaultQuery("name", "default name")
	//	c.JSON(200, gin.H{
	//		"name": name,
	//	})
	//})
	//
	//r.POST("/upload", func(c *gin.Context) {
	//	r.MaxMultipartMemory = 8 << 20
	//	file, err := c.FormFile("file")
	//	if err != nil {
	//		c.String(500,"upload failed~")
	//	}
	//	c.SaveUploadedFile(file,file.Filename)
	//	c.JSON(200, gin.H{
	//		"filename": file.Filename,
	//		"filesize": file.Size,
	//	})
	//
	//})
	//
	//r.POST("/upload/request", func(c *gin.Context) {
	//	_,headers,err := c.Request.FormFile("file")
	//	if err != nil{
	//		log.Printf("Error: %v",err)
	//		return
	//	}
	//	if headers.Size > 1024*1024*2{
	//		fmt.Println("filesize error!")
	//		return
	//	}
	//	if headers.Header.Get("Content-Type") != "image/png" {
	//		fmt.Println("need png type!")
	//		return
	//	}
	//	c.SaveUploadedFile(headers,"./img/"+headers.Filename)
	//	c.String(200,headers.Filename)
	//})
	//
	//
	//r.POST("/upload/multipart", func(c *gin.Context) {
	//	form, err := c.MultipartForm()
	//	if err!=nil {
	//		c.String(http.StatusBadRequest,fmt.Sprintf("get err: %v",err.Error()))
	//		return
	//	}
	//	files := form.File["files"]
	//	for _,file := range files {
	//		if err := c.SaveUploadedFile(file,file.Filename); err != nil {
	//			c.String(http.StatusBadRequest,fmt.Sprintf("upload err: %v",err.Error()))
	//			return
	//		}
	//	}
	//	c.String(http.StatusOK,fmt.Sprintf("files upload success!"))
	//
	//})
	//
	//g1 := r.Group("/g1")
	//{
	//	g1.GET("/login",login)
	//	g1.GET("/submit",submit1)
	//}
	//
	//g2 := r.Group("/g2")
	//{
	//	g2.POST("/login",login)
	//	g2.POST("/submit1",submit1)
	//	g2.GET("/submit2",submit2)
	//}
	//
	//
	//r.Run()

	//routers.Include(shop.LoadShopRouter,blog.LoadBlogRouter)
	//r := routers.Init()
	//if err := r.Run(); err != nil{
	//	fmt.Println("error to run !!!")
	//}

	//r := gin.Default()
	//r.POST("/loginJson", func(c *gin.Context) {
	//	var json Login
	//	if err := c.ShouldBindJSON(&json); err != nil{
	//		fmt.Println("error to parse json!")
	//	}
	//	if json.User != "lbh" || json.Pwd != "123456" {
	//		c.JSON(http.StatusBadRequest,gin.H{
	//			"message": "username or password wrong!",
	//		})
	//		return
	//	}
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "json tests successfully!",
	//	})
	//})
	//
	//r.POST("/loginForm", func(c *gin.Context) {
	//	var form Login
	//	if err := c.Bind(&form); err != nil{
	//		fmt.Println("error to parse form!")
	//	}
	//	if form.User != "lbh" || form.Pwd != "123456"{
	//		c.JSON(http.StatusBadRequest,gin.H{
	//			"message": "username or password wrong!",
	//		})
	//		return
	//	}
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "form tests successfully!",
	//	})
	//})
	//
	//r.GET("/loginUri/:user/:pwd", func(c *gin.Context) {
	//	var uri Login
	//	if err := c.ShouldBindUri(&uri); err != nil{
	//		fmt.Println("error to parse uri")
	//	}
	//	if uri.User != "lbh" || uri.Pwd != "123456"{
	//		c.JSON(http.StatusBadRequest,gin.H{
	//			"message": "username or password wrong!",
	//		})
	//	}
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "uri tests successfully!",
	//	})
	//})
	//
	//r.GET("/redirect", func(c *gin.Context) {
	//	c.Redirect(http.StatusMovedPermanently,"testRedirect")
	//})
	//
	//r.GET("/testRedirect", func(c *gin.Context) {
	//	c.String(http.StatusOK,"Redirect tests successfully!")
	//})

	r := gin.Default()
	r.Use(HandlerMiddleWare())
	{
		r.GET("/middleWare", func(c *gin.Context) {
			req, _ := c.Get("request")
			fmt.Println("request:", req)
			c.JSON(http.StatusOK, gin.H{
				"request": req,
			})
		})
	}

	r.Run()

}

func HandlerMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t1 := time.Now()
		fmt.Println("MiddleWare starts now")
		c.Set("request", "MiddleWare")
		status := c.Writer.Status()
		fmt.Println("status:", status)
		fmt.Println("MiddleWare ends now")
		t2 := time.Since(t1)
		fmt.Println("time:", t2)
	}
}
