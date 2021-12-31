package models

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"

	partModels "github.com/nwf-report/app/models/partModel"
	"github.com/nwf-report/services"
)

type UploadModel struct {
	Error       string
	ReportModel partModels.ReportModel
}

func NewUploadModel(c *gin.Context) (*UploadModel, error) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		err = errors.New("reportファイルに異常があります。入力したファイルを再度確認してください。")
		return nil, err
	}

	var item partModels.ReportModel
	err = json.NewDecoder(file).Decode(&item)
	if err != nil {
		err = errors.New("ファイルのデコードに失敗しました。reportファイルの形式を確認してください。")
		return nil, err
	}

	m := UploadModel{
		ReportModel: item,
	}

	return &m, nil
}

func (m *UploadModel) ConvertService(s *services.UploadService) error {
	b, err := json.Marshal(m.ReportModel)
	if err != nil {
		return errors.New("")
	}
	s.FileName = m.ReportModel.CaseName + ".json"
	s.ReportBlob = b
	return nil
}
