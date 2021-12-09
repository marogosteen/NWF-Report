package models

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nwf-report/services"
)

type UploadModel struct {
	Title     string
	BestEpoch int
	BestLoss  float64
	Observed  []float64
	Predicted []float64
}

func NewUploadModel(c *gin.Context) *UploadModel {
	bestEpoch, err := strconv.Atoi(c.PostForm("bestepoch"))
	if err != nil {
		// TODO 　log.Fatal()使わずに、errorはViewに返す
		log.Fatal(err)
	}
	bestLoss, err := strconv.ParseFloat(c.PostForm("bestloss"), 64)
	if err != nil {
		log.Fatal(err)
	}
	observed := readFormFile(c, "observed")
	predicted := readFormFile(c, "predicted")

	model := UploadModel{
		Title:     c.PostForm("reporttitle"),
		BestEpoch: bestEpoch,
		BestLoss:  bestLoss,
		Observed:  observed,
		Predicted: predicted,
	}

	return &model
}

func (m *UploadModel) ConvertToService(s *services.UploadService) {
	s.Title = m.Title
	s.BestEpoch = m.BestEpoch
	s.BestLoss = m.BestLoss
	s.Observed = m.Observed
	s.Predicted = m.Predicted
}

func readFormFile(c *gin.Context, formKey string) []float64 {
	file, fileHeader, err := c.Request.FormFile(formKey)
	if err != nil {
		log.Fatal(err)
	}

	dataArray := make([]float64, fileHeader.Size)
	reader := csv.NewReader(file)
	i := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(nil)
		}
		value, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}
		dataArray[i] = value
		i++
	}

	return dataArray
}
