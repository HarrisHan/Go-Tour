package main

import (
	"github.com/gin-gonic/gin"
	
	"net/http"
)

type Login struct {
	User string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func main() {

	router := gin.Default()
	//
	//router.GET("user/:name/*action", func(c *gin.Context){
	//	name := c.Param("name")
	//	action := c.Param("action")
	//	message := name + " is " + action
	//	//c.String(http.StatusOK, message)
	//})

	//router.GET("/welcom", func(c *gin.Context){
	//	firstname := c.DefaultQuery("firstname", "Guest")
	//	lastname  := c.Query("lastname")
	//
	//	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	//})

	//router.POST("/form_post", func(c *gin.Context) {
	//	message := c.PostForm("message")
	//	nick := c.DefaultPostForm("nick", "anoymous")
	//
	//	c.JSON(200, gin.H{
	//		"status": "posted",
	//		"message": message,
	//		"nick": nick,
	//	})
	//})

	//router.POST("/post", func(c *gin.Context) {
	//
	//	id := c.Query("id")
	//	page := c.DefaultQuery("page", "0")
	//	name := c.PostForm("name")
	//	message := c.PostForm("message")
	//
	//	fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	//})

	//router.POST("/upload", func(c *gin.Context) {
	//	file, _ := c.FormFile("file")
	//	log.Println(file.Filename)
	//	// 1.upload 2.save
	//
	//	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	//})

	//router.POST("/upload", func(c *gin.Context) {
	//	// Multipart form
	//	form, _ := c.MultipartForm()
	//	files := form.File["upload[]"]
	//
	//	for _, file := range files {
	//		log.Println(file.Filename)
	//		// upload and save the file
	//	}
	//	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	//})

	//v1 := router.Group("/v1")
	//{
	//	1.POST("/login", loginEndpoint)
	//	v1.POST("/submit", submitEndpoint)
	//	v1.POST("/read", readEndpoint)
	//}
	//
	//v2 := router.Group("/v2")
	//{
	//	v2.POST("/login", loginEndpoint)
	//	v2.POST("/submit", submitEndpoint)
	//	v2.POST("/read", readEndpoint)
	//}

	//r := gin.New()
	//
	//r.Use(gin.Logger())
	//r.Use(gin.Recovery())
	//// Per route middleware, you can add as many as you desire.
	//r.GET("/benchmark", MyBenchLogger(), benchEndPoint)
	//
	//authorized := r.Group("/")
	//authorized.Use(AuthRequired())
	//{
	//	authorized.POST("/login", loginEndpoint)
	//	authorized.POST("/submit", submitEndpoint)
	//	authorized.POST("read", readEndpoint)
	//
	//	testing := authorized.Group("testing")
	//	testing.GET("/analytics", analyticsEndpoint)
	//}
    	//r.Run(":8080")

	// Disable console color you don't need console color when writing the logs to file
	//gin.DisableConsoleColor()
	//
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)

	// Use the following code if you need to write the logs to file and console at the same time.
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	//router := gin.Default()
	//router.GET("/ping", func(c *gin.Context) {
	//	c.String(200,"pong")
	//})
	//

	router.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if c.BindJSON(&json) == nil {
			if json.User == "manu" && json.Password == "123" {
				c.JSON(http.StatusOK, gin.H{"status": "unauthorized"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			}
		}
	})

	// bingind a HTML form (user=manu&password=123)
	router.POST("/loginForm", func(c *gin.Context) {
		var form Login
		if c.Bind(&form) == nil {
			if form.User == "manu" && form.Password == "123" {
				c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			}
		}
	})

	router.Run(":8081")

}

