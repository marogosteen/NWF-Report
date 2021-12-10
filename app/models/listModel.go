package models

import "github.com/nwf-report/app/models/partModel"

type ListModel struct {
	ReportList []partModels.ReportModel
}

func NewListModel() ListModel {
	
	m := ListModel{

	}
	return m
}