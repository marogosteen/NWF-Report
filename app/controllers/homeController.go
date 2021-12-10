package controllers

import (
	"log"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"

	"github.com/nwf-report/app/models"
	"github.com/nwf-report/services"
)

type sample struct {
	title string
}

func StartWebServer(router *gin.Engine, port string) error {
	router.Static("/scripts", "app/views/scripts")
	router.Static("/styles", "app/views/styles")
	router.LoadHTMLGlob("app/views/*.html")
	router.GET("/", homeHandler)
	router.GET("/upload", uploadHandler)
	router.POST("/upload", uploadPostHandler)
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
	// TODO formのnil対処 viewで
	var maxContentLenght int64 = 1024 * 1024
	if c.Request.ContentLength > maxContentLenght {
		// TODO Viewに返す
		log.Fatalln("ContextLengthがmaxContentLenghtより大きい。")
	}

	m := models.NewUploadModel(c)
	var s services.UploadService
	m.ConvertToService(&s)
	s.Upload()
	c.Redirect(http.StatusFound, "/")
}
