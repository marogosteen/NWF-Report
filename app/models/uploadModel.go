package models

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"

	partModels "github.com/nwf-report/app/models/partModel"
	"github.com/nwf-report/services"
)

type UploadModel struct {
	ReportModel partModels.ReportModel
}

func NewUploadModel(c *gin.Context) *UploadModel {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		log.Fatalln(err)
	}
	var item partModels.ReportModel
	err = json.NewDecoder(file).Decode(&item)
	if err != nil {
		log.Fatalln(err)
	}
	m := UploadModel{
		ReportModel: item,
	}
	return &m
}

func (m *UploadModel) ConvertService(s *services.UploadService) {
	b, err := json.Marshal(m.ReportModel)
	if err != nil {
		log.Fatalln(err)
	}
	s.FileName = m.ReportModel.CaseName + ".json"
	s.ReportBlob = b
}
