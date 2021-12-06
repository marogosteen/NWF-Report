package models

import (
	"github.com/gin-gonic/gin"
)

type AdminModel struct {
}

func NewAdminModel() *AdminModel {
	model := AdminModel{}
	return &model
}

func (m *AdminModel) ViewModel() *gin.H {
	viewModel := gin.H{
		"title": nil,
		"hoge":  nil,
	}
	return &viewModel
}
