package helpers

import (
	"strconv"

	"github.com/xuri/excelize/v2"
)

type ExcelSheetBuilder struct {
	file      *excelize.File
	sheetName string
	headers   []string
	data      [][]interface{}
	styles    map[string]*excelize.Style
}

// NewExcelSheetBuilder initializes a new sheet builder.
func NewExcelSheetBuilder(sheetName string) *ExcelSheetBuilder {
	file := excelize.NewFile()
	file.NewSheet(sheetName)

	return &ExcelSheetBuilder{
		file:      file,
		sheetName: sheetName,
		styles:    make(map[string]*excelize.Style),
	}
}

// SetHeaders sets the headers for the sheet.
func (b *ExcelSheetBuilder) SetHeaders(headers []string) *ExcelSheetBuilder {
	b.headers = headers
	return b
}

// AddRow adds a data row to the sheet.
func (b *ExcelSheetBuilder) AddRow(row []interface{}) *ExcelSheetBuilder {
	b.data = append(b.data, row)
	return b
}

// AddStyle adds a style to a specific column.
func (b *ExcelSheetBuilder) AddStyle(column string, style *excelize.Style) *ExcelSheetBuilder {
	b.styles[column] = style
	return b
}

// Build constructs the sheet with headers, data, and styles.
func (b *ExcelSheetBuilder) Build() ([]byte, error) {
	// Write headers to the sheet.
	for colIdx, header := range b.headers {
		cell := getExcelColumnName(colIdx) + "1"
		b.file.SetCellValue(b.sheetName, cell, header)
	}

	// Write data rows.
	for rowIdx, row := range b.data {
		for colIdx, value := range row {
			cell := getExcelColumnName(colIdx) + strconv.Itoa(rowIdx+2)
			b.file.SetCellValue(b.sheetName, cell, value)
		}
	}

	// Apply styles.
	for column, style := range b.styles {
		styleID, err := b.file.NewStyle(style)
		if err != nil {
			return nil, err
		}

		// Apply style to the entire column.
		for rowIdx := 2; rowIdx <= len(b.data)+1; rowIdx++ {
			cell := column + strconv.Itoa(rowIdx)
			b.file.SetCellStyle(b.sheetName, cell, cell, styleID)
		}
	}

	// Write to a buffer and return as a byte slice.
	buf, err := b.file.WriteToBuffer()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Helper function to convert a column index to Excel column letters.
func getExcelColumnName(colIdx int) string {
	colName := ""
	for colIdx >= 0 {
		remainder := colIdx % 26
		colName = string(rune('A'+remainder)) + colName
		colIdx = (colIdx / 26) - 1
	}
	return colName
}
