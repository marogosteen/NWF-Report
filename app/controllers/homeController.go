package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/nwf-report/app/models"
	"github.com/nwf-report/services"
)

func StartWebServer(router *gin.Engine, port string) error {
	router.Static("/scripts", "app/views/scripts")
	router.Static("/styles", "app/views/styles")
	router.LoadHTMLGlob("app/views/*.html")

	router.GET("/", homeHandler)
	router.GET("/detail", detailHandler)
	router.GET("/upload", uploadHandler)
	router.POST("/upload", uploadPostHandler)

	err := router.Run(":" + port)

	return err
}

func homeHandler(c *gin.Context) {
	var s services.ListService
	s.Fetch()
	var m models.ListModel
	m.ConvertModel(s)
	c.HTML(http.StatusOK, "index.html", m)
}

func detailHandler(c *gin.Context) {
	var s services.DetailService
	reportName := "2019"
	s.SearchBlob(reportName)
	var m models.DetailModel
	m.ConvertModel(s)
	c.HTML(http.StatusOK, "detail.html", m)
}

func uploadHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.html", gin.H{})
}

func uploadPostHandler(c *gin.Context) {
	// TODO formのnil対処 viewで
	var maxContentLenght int64 = 1024 * 1024
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
