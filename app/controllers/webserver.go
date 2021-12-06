package controllers

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"

	"example.com/osunnwf.report/app/models"
	// "example.com/osunnwf.report/services"
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
	router.GET("/", homeViewHandler)
	router.GET("/admin", adminViewHandler)
	router.GET("/add", addViewHandler)
	router.POST("/add", addFormReadHandler)
	// TODO configを扱ったPortのセット
	err := router.Run(":" + port)
	return err
}

func homeViewHandler(c *gin.Context) {
	model := sample{title: "OK!"}
	elem := reflect.ValueOf(&model).Elem()
	size := elem.NumField()
	viewInterface := gin.H{}
	for i := 0; i < size; i++ {
		field := elem.Type().Field(i).Name
		value := elem.Field(i)
		viewInterface[field] = value
	}
	c.HTML(http.StatusOK, "index.html", viewInterface)
}

func adminViewHandler(c *gin.Context) {
	// TODO CaseのListView
}

func addViewHandler(c *gin.Context) {
	model := models.ReportAddModel{}
	elem := reflect.ValueOf(&model).Elem()
	size := elem.NumField()
	viewInterface := gin.H{}
	for i := 0; i < size; i++ {
		field := elem.Type().Field(i).Name
		value := elem.Field(i)
		viewInterface[field] = value
	}
	c.HTML(http.StatusOK, "add.html", viewInterface)
}

func addFormReadHandler(c *gin.Context) {
	// TODO formのnil対処
	model := models.NewReportAddModel(c)
	service := model.ConvertToService()
	service.InsertReport()
	// TODO ListViewにRedirect
}
