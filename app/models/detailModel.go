package models

import (
	"encoding/json"
	"log"

	partModels "github.com/nwf-report/app/models/partModel"
	"github.com/nwf-report/services"
)

type DetailModel struct {
	ReportModel partModels.ReportModel
}

func (m *DetailModel) ConvertModel(s services.DetailService) {
	err := json.Unmarshal(s.ReportBlob, &m.ReportModel)
	if err != nil {
		log.Fatalln(err)
	}
}
