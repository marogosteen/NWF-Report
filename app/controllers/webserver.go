package controllers

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"

	"github.com/nwf-report/app/models"
	"github.com/nwf-report/services"
)

var (
	port   = "8080"
	router = gin.Default()
)

type sample struct {
	title string
}

func StartWebServer() error {
	router.Static("/scripts", "app/views/scripts")
	router.Static("/styles", "app/views/styles")
	router.LoadHTMLGlob("app/views/*.html")
	router.GET("/", homeHandler)
	router.GET("/admin", adminHandler)
	router.GET("/upload", uploadHandler)
	router.POST("/upload", uploadPostHandler)
	// TODO configを扱ったPortのセット
	err := router.Run(":" + port)
	return err
}

func homeHandler(c *gin.Context) {
	m := sample{title: "OK!"}
	elem := reflect.ValueOf(&m).Elem()
	size := elem.NumField()
	viewInterface := gin.H{}
	for i := 0; i < size; i++ {
		field := elem.Type().Field(i).Name
		value := elem.Field(i)
		viewInterface[field] = value
	}
	c.HTML(http.StatusOK, "index.html", viewInterface)
}

func adminHandler(c *gin.Context) {
	// TODO CaseのListView
}

func uploadHandler(c *gin.Context) {
	m := models.UploadModel{}
	elem := reflect.ValueOf(&m).Elem()
	size := elem.NumField()
	viewInterface := gin.H{}
	for i := 0; i < size; i++ {
		field := elem.Type().Field(i).Name
		value := elem.Field(i)
		viewInterface[field] = value
	}
	c.HTML(http.StatusOK, "upload.html", viewInterface)
}

func uploadPostHandler(c *gin.Context) {
	// TODO formのnil対処 csvのFileSizeの規制
	m := models.NewUploadModel(c)
	var s services.UploadService
	m.ConvertToService(&s)
	s.Upload()
	// TODO ListViewにRedirect
}
