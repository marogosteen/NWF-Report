package models

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	partModels "github.com/nwf-report/app/models/partModel"
	"github.com/nwf-report/services"
)

type DetailModel struct {
	ReportModel partModels.ReportModel
	Datetime    []string
}

const (
	shortForm = "2006-01-02 00:00"
	strlayout = "2006-01-02"
)

func (m *DetailModel) ConvertModel(s services.DetailService) {
	err := json.NewDecoder(s.Reader).Decode(&m.ReportModel)
	if err != nil {
		log.Fatalln(err)
	}

	duration, _ := time.ParseDuration("1h")
	strStartDate := fmt.Sprintf("%v-01-01 00:00", m.ReportModel.TargetYear)
	date, _ := time.Parse(shortForm, strStartDate)
	for {
		if date.Year() == m.ReportModel.TargetYear+1 {
			break
		}
		strhour := fmt.Sprintf(" %v:00", date.Hour())
		strdate := date.Format(strlayout) + strhour
		m.Datetime = append(m.Datetime, strdate)
		date = date.Add(duration)
	}
}
