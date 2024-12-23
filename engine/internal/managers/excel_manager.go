package managers

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/xuri/excelize/v2"
)

// ExportExcelBuilder facilitates building and exporting Excel files.
type ExportExcelBuilder struct {
	sheets []Sheet
}

// Sheet represents a single sheet in the Excel file.
type Sheet struct {
	Name     string
	Data     interface{}
	Headers  []string
	Filters  func(interface{}) bool
	Styles   map[string]int // Cell address to style ID
	AutoSize bool
}

// NewExportExcelBuilder initializes a new ExportExcelBuilder with default settings.
func NewExportExcelBuilder() *ExportExcelBuilder {
	return &ExportExcelBuilder{
		sheets: []Sheet{},
	}
}

// AddSheet adds a new sheet to the Excel file.
func (b *ExportExcelBuilder) AddSheet(name string, data interface{}) *ExportExcelBuilder {
	sheet := Sheet{
		Name:     name,
		Data:     data,
		Headers:  []string{},
		AutoSize: true,
	}
	b.sheets = append(b.sheets, sheet)
	return b
}

// SetHeaders sets the headers for a specific sheet by name.
func (b *ExportExcelBuilder) SetHeaders(sheetName string, headers []string) *ExportExcelBuilder {
	for i, sheet := range b.sheets {
		if sheet.Name == sheetName {
			b.sheets[i].Headers = headers
			break
		}
	}
	return b
}

// SetFilters sets a filter function for a specific sheet by name.
func (b *ExportExcelBuilder) SetFilters(sheetName string, filter func(interface{}) bool) *ExportExcelBuilder {
	for i, sheet := range b.sheets {
		if sheet.Name == sheetName {
			b.sheets[i].Filters = filter
			break
		}
	}
	return b
}

// SetStyles sets cell styles for a specific sheet by name.
func (b *ExportExcelBuilder) SetStyles(sheetName string, styles map[string]int) *ExportExcelBuilder {
	for i, sheet := range b.sheets {
		if sheet.Name == sheetName {
			b.sheets[i].Styles = styles
			break
		}
	}
	return b
}

// EnableAutoSize toggles auto-sizing of columns for a specific sheet by name.
func (b *ExportExcelBuilder) EnableAutoSize(sheetName string, enable bool) *ExportExcelBuilder {
	for i, sheet := range b.sheets {
		if sheet.Name == sheetName {
			b.sheets[i].AutoSize = enable
			break
		}
	}
	return b
}

func (b *ExportExcelBuilder) CreateStyle(f *excelize.File, style *excelize.Style) (int, error) {
	return f.NewStyle(style)
}

func (b *ExportExcelBuilder) Build() ([]byte, error) {
	f := excelize.NewFile()
	defaultSheetName := f.GetSheetName(0)
	f.DeleteSheet(defaultSheetName)

	for idx, sheet := range b.sheets {
		if idx == 0 {
			if existingSheetIndex, err := f.GetSheetIndex(sheet.Name); err == nil && existingSheetIndex > 0 {
				f.SetSheetName(f.GetSheetName(existingSheetIndex), sheet.Name)
			} else {
				f.NewSheet(sheet.Name)
			}
		} else {
			f.NewSheet(sheet.Name)
		}

		data := reflect.ValueOf(sheet.Data)
		if data.Kind() != reflect.Slice && data.Kind() != reflect.Array {
			return nil, fmt.Errorf("data for sheet %s is not a slice or array", sheet.Name)
		}

		var filteredData []interface{}
		for i := 0; i < data.Len(); i++ {
			item := data.Index(i).Interface()
			if sheet.Filters == nil || sheet.Filters(item) {
				filteredData = append(filteredData, item)
			}
		}

		if len(sheet.Headers) > 0 {
			for colIdx, header := range sheet.Headers {
				cell, _ := excelize.CoordinatesToCellName(colIdx+1, 1)
				if err := f.SetCellValue(sheet.Name, cell, header); err != nil {
					return nil, err
				}
			}
		} else {
			if len(filteredData) > 0 {
				itemType := reflect.TypeOf(filteredData[0])
				if itemType.Kind() == reflect.Ptr {
					itemType = itemType.Elem()
				}
				if itemType.Kind() == reflect.Struct {
					for j := 0; j < itemType.NumField(); j++ {
						field := itemType.Field(j)
						cell, _ := excelize.CoordinatesToCellName(j+1, 1)
						if err := f.SetCellValue(sheet.Name, cell, field.Name); err != nil {
							return nil, err
						}
						sheet.Headers = append(sheet.Headers, field.Name)
					}
				}
			}
		}

		for rowIdx, item := range filteredData {
			v := reflect.ValueOf(item)
			if v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			if v.Kind() != reflect.Struct {
				return nil, fmt.Errorf("data item is not a struct in sheet %s", sheet.Name)
			}
			for colIdx := 0; colIdx < v.NumField(); colIdx++ {
				field := v.Field(colIdx)
				var value interface{}
				if field.CanInterface() {
					value = field.Interface()
				}
				cell, _ := excelize.CoordinatesToCellName(colIdx+1, rowIdx+2)
				if err := f.SetCellValue(sheet.Name, cell, value); err != nil {
					return nil, err
				}

				if sheet.Styles != nil {
					if styleID, exists := sheet.Styles[cell]; exists {
						if err := f.SetCellStyle(sheet.Name, cell, cell, styleID); err != nil {
							return nil, err
						}
					}
				}
			}
		}

		if sheet.AutoSize && len(sheet.Headers) > 0 {
			for colIdx := range sheet.Headers {
				colLetter, err := excelize.ColumnNumberToName(colIdx + 1)
				if err != nil {
					return nil, err
				}
				maxWidth := len(sheet.Headers[colIdx])
				for rowIdx := range filteredData {
					cell, _ := excelize.CoordinatesToCellName(colIdx+1, rowIdx+2)
					cellValue, err := f.GetCellValue(sheet.Name, cell)
					if err == nil && len(cellValue) > maxWidth {
						maxWidth = len(cellValue)
					}
				}
				if err := f.SetColWidth(sheet.Name, colLetter, colLetter, float64(maxWidth+2)); err != nil {
					return nil, err
				}
			}
		}
	}

	var buffer bytes.Buffer
	if err := f.Write(&buffer); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
