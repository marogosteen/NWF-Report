package models

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"

	"example.com/osunnwf.report/services"
)

type ReportAddModel struct {
	Title     string
	BestEpoch int
	BestLoss  float64
	Observed  []float64
	Predicted []float64
}

func NewReportAddModel(c *gin.Context) *ReportAddModel {
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

	model := ReportAddModel{
		Title:     c.PostForm("reporttitle"),
		BestEpoch: bestEpoch,
		BestLoss:  bestLoss,
		Observed:  observed,
		Predicted: predicted,
	}

	return &model
}

func (m *ReportAddModel) ConvertToService() services.CreateService {
	service := services.CreateService{
		Title:     m.Title,
		BestEpoch: m.BestEpoch,
		BestLoss:  m.BestLoss,
		Observed:  m.Observed,
		Predicted: m.Predicted,
	}
	return service
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
