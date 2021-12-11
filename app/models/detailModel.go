package models

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/nwf-report/app/models/partModel"
	"github.com/nwf-report/services"
)

type DetailModel struct {
	ReportModel partModels.ReportModel
}

func NewDetailModel(c *gin.Context) *DetailModel {
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

	model := DetailModel{
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

func (m *DetailModel) ConvertModel(s services.DetailService) {

}
