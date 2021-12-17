package models

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/nwf-report/app/models/partModel"
	"github.com/nwf-report/services"
)

type UploadModel struct {
	ReportModel partModels.ReportModel
}

func NewUploadModel(c *gin.Context) *UploadModel {
	file, err := c.FormFile("file")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(file.Size)
	reader, err := file.Open()
	if err != nil {
		log.Fatalln(err)
	}
	var b []byte
	_, err = reader.Read(b)
	if err != nil {
		log.Fatalln(err)
	}

	var model UploadModel
	err = json.Unmarshal(b, model)
	if err != nil {
		log.Fatalln(err)
	}
	// model := UploadModel{
	// 	ReportModel: partModels.ReportModel{
	// 		Title:     c.PostForm("reporttitle"),
	// 		BestEpoch: bestEpoch,
	// 		BestLoss:  bestLoss,
	// 		Observed:  observed,
	// 		Predicted: predicted,
	// 	},
	// }

	return &model
}

func (m *UploadModel) ConvertService(s *services.UploadService) {
	b, err := json.Marshal(m.ReportModel)
	if err != nil {
		log.Fatalln(err)
	}
	s.FileName = m.ReportModel.CaseName + ".json"
	s.ReportBlob = b
}

func readFormFile(c *gin.Context, formKey string) [][]float64 {
	file, _, err := c.Request.FormFile(formKey)
	if err != nil {
		log.Fatalln(err)
	}

	dataArray := [][]float64{}
	reader := csv.NewReader(file)

	for {
		strline, err := reader.Read()
		if err != nil {
			break
		}

		var floatline []float64
		for _, strv := range strline {
			floatv, err := strconv.ParseFloat(strv, 64)
			if err != nil {
				log.Fatalln(err)
			}
			floatline = append(floatline, floatv)
		}

		dataArray = append(dataArray, floatline)
	}

	return dataArray
}
