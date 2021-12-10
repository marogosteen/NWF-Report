package models

import (
	"encoding/csv"
	"encoding/json"
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
	bestEpoch, err := strconv.Atoi(c.PostForm("bestepoch"))
	if err != nil {
		// TODO 　log.Fatal()使わずに、errorはViewに返す
		log.Fatalln(err)
	}
	bestLoss, err := strconv.ParseFloat(c.PostForm("bestloss"), 64)
	if err != nil {
		// TODO 　log.Fatal()使わずに、errorはViewに返す
		log.Fatalln(err)
	}
	observed := readFormFile(c, "observed")
	predicted := readFormFile(c, "predicted")

	model := UploadModel{
		ReportModel: partModels.ReportModel{
			Title:     c.PostForm("reporttitle"),
			BestEpoch: bestEpoch,
			BestLoss:  bestLoss,
			Observed:  observed,
			Predicted: predicted,
		},
	}

	return &model
}

func (m *UploadModel) ConvertToService(s *services.UploadService) {
	b, err := json.Marshal(s)
	if err != nil {
		log.Fatalln(err)
	}
	s.FileName = m.ReportModel.Title + ".json"
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
