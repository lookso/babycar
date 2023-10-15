package pkg

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type Excel struct {
	File *excelize.File
}

func NewExcel() *Excel {
	return &Excel{
		File: excelize.NewFile(),
	}
}

func (e *Excel) SetHeader(sheetName string, headers []string) {
	for i, header := range headers {
		col := string('A' + i)
		e.File.SetCellValue(sheetName, col+"1", header)
	}
}

func (e *Excel) Write(ctx http.Context, content []byte, fileName string) error {
	_, err := ctx.Response().Write(content)
	if err != nil {
		return err
	}
	ctx.Response().Header().Set("Content-Disposition", "attachment; filename="+fileName)
	ctx.Response().Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

	return nil
}
