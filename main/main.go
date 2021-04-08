package main

//type Login struct {
//	User string `form:"username" json:"user" uri:"user" binding:"required"`
//	Pwd  string `form:"password" json:"pwd" uri:"pwd" binding:"required"`
//}

//type MyClaims struct {
//	jwt.StandardClaims
//	Username string `json:"username"`
//}
//
//type User struct {
//	Username string `form:"username" json:"username" binding:"required"`
//	Password string `form:"password" json:"password" binding:"required"`
//}
//
//const TokenExpireDuration = 120 * time.Second
//
//var Secret = []byte("handsome")
//
////获取token
//func GetToken(username string) (string,error){
//	//声明字段
//	c := &MyClaims{
//		jwt.StandardClaims{
//			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
//			Issuer: "project-hi",
//		},
//		username,
//	}
//	//获取token并将字段存入token中
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256,c)
//	return token.SignedString(Secret)
//
//}
//
////解析token
//func ParseToken(tokenString string) (*MyClaims,error){
//	token, err := jwt.ParseWithClaims(tokenString,&MyClaims{}, func(token *jwt.Token) (interface{}, error) {
//		return Secret,nil
//	})
//	if err != nil{
//		return nil, err
//	}
//	if claims,ok := token.Claims.(*MyClaims); ok && token.Valid{
//		return claims,nil
//	}
//	return nil,errors.New("invalid token1")
//}
//
//func authHandler(c *gin.Context){
//	var user User
//	if err := c.ShouldBind(&user); err != nil{
//		c.String(2000,fmt.Sprintf("Error when bind user:%s",err))
//		return
//	}
//
//	if user.Username != "lbh" || user.Password != "123456" {
//		c.JSON(2001,gin.H{
//			"status": "fail",
//			"message": "鉴权失败",
//		})
//	}
//	tokenString,_ := GetToken(user.Username)
//	c.JSON(http.StatusOK,gin.H{
//		"status": "success",
//		"message": gin.H{"token": tokenString},
//	})
//
//}
//
//func JWTVerify() func(c *gin.Context){
//	return func(c *gin.Context) {
//		//获取请求头中的"Authorization"
//		authHeader := c.GetHeader("Authorization")
//		if authHeader == ""{
//			c.JSON(http.StatusOK,gin.H{
//				"message": "no auth!",
//			})
//			c.Abort()
//			return
//		}
//
//		claims,err := ParseToken(authHeader)
//		if err != nil {
//			c.JSON(http.StatusOK,gin.H{
//				"message": "invalid token2",
//				"err": err,
//			})
//			c.Abort()
//			return
//		}
//		c.Set("username",claims.Username)
//		c.Next()
//
//	}
//}
//
//func HomeHandler(c *gin.Context){
//	username := c.MustGet("username").(string)
//	c.JSON(http.StatusOK,gin.H{
//		"status": "success",
//		"data": username,
//	})
//}
//
//func main()  {
//	r := gin.Default()
//	r.POST("/send",authHandler)
//	r.GET("/login",JWTVerify(),HomeHandler)
//	r.Run()
//}

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

//func main() {

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

//r := gin.Default()
//r.Use(HandlerMiddleWare())
//{
//	r.GET("/middleWare", func(c *gin.Context) {
//		req, _ := c.Get("request")
//		fmt.Println("request:", req)
//		c.JSON(http.StatusOK, gin.H{
//			"request": req,
//		})
//	})
//}

//r := gin.Default()
//temp := ""
//r.Use(func(c *gin.Context) {
//	temp += "A"
//	c.Next()
//	fmt.Println("MiddleWare1:",temp)
//})
//
//r.Use(func(c *gin.Context) {
//	temp += "B"
//	fmt.Println("MiddleWare2:",temp)
//})
//
//r.GET("/testNext", func(c *gin.Context) {
//	temp += "C"
//	fmt.Println("MiddleWare3:",temp)
//})

//r := gin.Default()
//r.Use(MiddleWare)
//shopGroup := r.Group("/shop")
//{
//	shopGroup.GET("/index",ShopIndexHandler)
//	shopGroup.GET("/comment",ShopCommentHandler)
//}

//r := gin.Default()
//r.GET("cookie", func(c *gin.Context) {
//	cookie, err := c.Cookie("key_cookie")
//	if err != nil {
//		cookie = "NoSet"
//		c.SetCookie("key_cookie","value_cookie",5,"/","localhost",false,true)
//	}
//	fmt.Printf("cookie is:%s\n",cookie)
//})

//	r := gin.Default()
//	cookieGroup := r.Group("/cookie")
//	{
//		cookieGroup.GET("/login",login)
//		cookieGroup.GET("/home",AuthMiddleWare(),home)
//	}
//
//	r.Run()
//
//}

//func HandlerMiddleWare() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		t1 := time.Now()
//		fmt.Println("MiddleWare starts now")
//		c.Set("request", "MiddleWare")
//		c.Next()
//		status := c.Writer.Status()
//		fmt.Println("status:", status)
//		fmt.Println("MiddleWare ends now")
//		t2 := time.Since(t1)
//		fmt.Println("time:", t2)
//	}
//}

//func ShopIndexHandler(c *gin.Context){
//	time.Sleep(2*time.Second)
//}
//func ShopCommentHandler(c *gin.Context){
//	time.Sleep(3*time.Second)
//}
//
//func MiddleWare(c *gin.Context) {
//	t1 := time.Now()
//	c.Next()
//	t2 := time.Since(t1)
//	fmt.Println("time:",t2)
//}

//func login(c *gin.Context){
//	_, err := c.Cookie("key_cookie")
//	if err != nil{
//		c.SetCookie("key_cookie","pass",10,"/","localhost",false,true)
//	}
//	c.String(http.StatusOK,"Success to login and you have a cookie now!")
//}
//
//func home(c *gin.Context) {
//	c.JSON(http.StatusOK,gin.H{
//		"data": "home",
//	})
//}
//
//func AuthMiddleWare() gin.HandlerFunc{
//	return func(c *gin.Context) {
//		cookie, err := c.Cookie("key_cookie")
//		if cookie == "pass"{
//			c.Next()
//			return
//		}
//		c.String(http.StatusBadRequest,fmt.Sprintf("Error: %s",err))
//		c.Abort()
//	}
//}
