package managers

import (
	"encoding/csv"

	"github.com/gin-gonic/gin"
)

// CSVManager manages CSV file creation with configurable parameters.
type CSVManager struct {
	FileName       string
	Headers        []string
	Records        [][]string
	FilterFunc     func([]string) bool
	TransformFunc  func([]string) []string
	Delimiter      rune // Optional: Allow custom delimiter
	UseCustomDelim bool // Flag to indicate if a custom delimiter is used
}

// NewCSVManager initializes a new CSVManager with default settings.
func NewCSVManager() *CSVManager {
	return &CSVManager{
		FileName: "output.csv",
		Headers:  []string{},
		Records:  [][]string{},
	}
}

// SetFileName sets the name of the CSV file.
func (cm *CSVManager) SetFileName(name string) {
	cm.FileName = name
}

// SetHeaders sets the header columns for the CSV.
func (cm *CSVManager) SetHeaders(headers []string) {
	cm.Headers = headers
}

// AddRecords appends multiple records to the CSV.
func (cm *CSVManager) AddRecords(records [][]string) {
	cm.Records = append(cm.Records, records...)
}

// AddRecord appends a single record to the CSV.
func (cm *CSVManager) AddRecord(record []string) {
	cm.Records = append(cm.Records, record)
}

// SetFilter sets a filter function to include only specific records.
func (cm *CSVManager) SetFilter(filterFunc func([]string) bool) {
	cm.FilterFunc = filterFunc
}

// SetTransform sets a transformation function to modify records.
func (cm *CSVManager) SetTransform(transformFunc func([]string) []string) {
	cm.TransformFunc = transformFunc
}

// SetDelimiter sets a custom delimiter for the CSV.
func (cm *CSVManager) SetDelimiter(delim rune) {
	cm.Delimiter = delim
	cm.UseCustomDelim = true
}

// WriteCSV generates the CSV and writes it to the provided Gin context.
func (cm *CSVManager) WriteCSV(c *gin.Context) error {
	// Set the headers for the HTTP response
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+cm.FileName)
	c.Header("Content-Type", "text/csv")

	// Create a new CSV writer using Gin's ResponseWriter
	writer := csv.NewWriter(c.Writer)
	if cm.UseCustomDelim {
		writer.Comma = cm.Delimiter
	}
	defer writer.Flush()

	// Write the header if set
	if len(cm.Headers) > 0 {
		if err := writer.Write(cm.Headers); err != nil {
			return err
		}
	}

	// Iterate over the records
	for _, record := range cm.Records {
		// Apply filter if set
		if cm.FilterFunc != nil && !cm.FilterFunc(record) {
			continue
		}

		// Apply transformation if set
		if cm.TransformFunc != nil {
			record = cm.TransformFunc(record)
		}

		// Write the record
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}
