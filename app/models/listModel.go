package models

import (
	"regexp"

	"github.com/nwf-report/services"
)

type ListModel struct {
	ReportList []string
}

var r *regexp.Regexp

func init() {
	r = regexp.MustCompile(`.json`)
}

func (m *ListModel) ConvertModel(s services.ListService) {
	reportList := m.ReportList
	for _, blobname := range s.ReportList {
		s := r.ReplaceAllString(blobname, "")
		reportList = append(reportList, s)
	}
	m.ReportList = reportList
}
