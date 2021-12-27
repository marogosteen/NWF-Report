package controllers

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/nwf-report/app/models"
	"github.com/nwf-report/services"
)

const (
	maxContentLenght int64 = 4 * 1024 * 1024
)

func StartWebServer() error {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.Static("/scripts", "app/views/scripts")
	router.Static("/styles", "app/views/styles")
	router.LoadHTMLGlob("app/views/*.html")

	router.GET("/", homeHandler)
	router.GET("/detail/:reportname", detailHandler)
	router.GET("/upload", uploadHandler)
	router.POST("/upload", uploadPostHandler)
	router.POST("/delete/:reportname", deletePostHandler)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	err := router.Run(":" + port)

	return err
}

func homeHandler(c *gin.Context) {
	var s services.ListService
	// TODO Pager
	s.Fetch()
	var m models.ListModel
	m.ConvertModel(s)
	c.HTML(http.StatusOK, "index.html", m)
}

func detailHandler(c *gin.Context) {
	fileName := c.Param("reportname") + ".json"
	var s services.DetailService
	s.SearchBlob(fileName)
	var m models.DetailModel
	m.ConvertModel(s)
	c.HTML(http.StatusOK, "detail.html", m)
}

func uploadHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.html", gin.H{})
}

func uploadPostHandler(c *gin.Context) {
	// TODO formのnil対処 viewで
	if c.Request.ContentLength > maxContentLenght {
		// TODO Viewに返す
		log.Fatalln("ContextLengthがmaxContentLenghtより大きい。")
	}

	m := models.NewUploadModel(c)
	var s services.UploadService
	m.ConvertService(&s)
	s.Upload()

	c.Redirect(http.StatusFound, "/upload")
}

func deletePostHandler(c *gin.Context) {
	fileName := c.Param("casename") + ".json"
	var s services.DeleteService
	s.Delete(fileName)
	
	c.Redirect(http.StatusFound, "/")
}