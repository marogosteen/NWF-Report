package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/nwf-report/app/models"
	"github.com/nwf-report/config"
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

	config := config.Config
	err := router.Run(":" + config.Port)

	return err
}

func homeHandler(c *gin.Context) {
	var s services.ListService
	// TODO Pager
	s.Fetch()
	var m models.ListModel
	m.ConvertModel(s)
	c.HTML(200, "index.html", m)
}

func detailHandler(c *gin.Context) {
	fileName := c.Param("reportname") + ".json"
	var s services.DetailService
	s.SearchBlob(fileName)
	var m models.DetailModel
	m.ConvertModel(s)
	c.HTML(200, "detail.html", m)
}

func uploadHandler(c *gin.Context) {
	m := models.UploadModel{}
	c.HTML(200, "upload.html", m)
}

func uploadPostHandler(c *gin.Context) {
	if c.Request.ContentLength > maxContentLenght {
		model := models.UploadModel{
			Error: "ファイルサイズは4MBまでです。",
		}
		c.HTML(400, "upload.html", model)
		c.Abort()
		return
	}

	m, err := models.NewUploadModel(c)
	if err != nil {
		model := models.UploadModel{
			Error: err.Error(),
		}
		c.HTML(400, "upload.html", model)
		c.Abort()
		return
	}

	var s services.UploadService
	err = m.ConvertService(&s)
	if err != nil {
		model := models.UploadModel{
			Error: err.Error(),
		}
		c.HTML(400, "upload.html", model)
		c.Abort()
		return
	}
	s.Upload()

	c.Redirect(302, "/")
}

func deletePostHandler(c *gin.Context) {
	fileName := c.Param("reportname") + ".json"
	var s services.DeleteService
	s.Delete(fileName)
	c.Redirect(302, "/")
}
