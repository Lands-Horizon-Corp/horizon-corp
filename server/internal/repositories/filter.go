package repositories

import (
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type TSearchFilter struct {
	DataType string      `json:"dataType"`        // 'number' | 'text' | 'date' | 'enum'
	Value    interface{} `json:"value,omitempty"` // Single value or array for 'enum'
	Mode     string      `json:"mode"`            // Filter mode
	From     interface{} `json:"from,omitempty"`  // For range filters
	To       interface{} `json:"to,omitempty"`    // For range filters
}

type TFinalFilter struct {
	Field         string      `json:"field"`   // Field to filter on
	Preload       string      `json:"preload"` // Preload related data
	Value         interface{} `json:"value"`   // Value or range
	TSearchFilter             // Embedding TSearchFilter
}

// FilterParam represents the entire filter payload.
type FilterParam struct {
	Filters   []TFinalFilter `json:"filters"`   // Array of filters
	PageIndex int            `json:"pageIndex"` // Current page index
	PageSize  int            `json:"pageSize"`  // Number of items per page
}

func (fb *Repository[T]) ApplyFilters(filters []TFinalFilter) (*gorm.DB, error) {
	query := fb.DB

	for _, filter := range filters {
		switch strings.ToLower(filter.Mode) {
		case "equal":
			query = query.Where(fmt.Sprintf("%s = ?", filter.Field), filter.Value)
		case "nequal":
			query = query.Where(fmt.Sprintf("%s != ?", filter.Field), filter.Value)
		case "contains":
			query = query.Where(fmt.Sprintf("%s LIKE ?", filter.Field), "%"+filter.Value.(string)+"%")
		case "ncontains":
			query = query.Where(fmt.Sprintf("%s NOT LIKE ?", filter.Field), "%"+filter.Value.(string)+"%")
		case "startswith":
			query = query.Where(fmt.Sprintf("%s LIKE ?", filter.Field), filter.Value.(string)+"%")
		case "endswith":
			query = query.Where(fmt.Sprintf("%s LIKE ?", filter.Field), "%"+filter.Value.(string))
		case "isempty":
			query = query.Where(fmt.Sprintf("%s IS NULL OR %s = ''", filter.Field, filter.Field))
		case "isnotempty":
			query = query.Where(fmt.Sprintf("%s IS NOT NULL AND %s != ''", filter.Field, filter.Field))
		case "gt":
			query = query.Where(fmt.Sprintf("%s > ?", filter.Field), filter.Value)
		case "gte":
			query = query.Where(fmt.Sprintf("%s >= ?", filter.Field), filter.Value)
		case "lt":
			query = query.Where(fmt.Sprintf("%s < ?", filter.Field), filter.Value)
		case "lte":
			query = query.Where(fmt.Sprintf("%s <= ?", filter.Field), filter.Value)
		case "range":
			rangeValues, ok := filter.Value.(map[string]interface{})
			if !ok {
				return nil, errors.New("invalid range value")
			}
			from, ok1 := rangeValues["from"]
			to, ok2 := rangeValues["to"]
			if !ok1 || !ok2 {
				return nil, errors.New("range filter requires 'from' and 'to' values")
			}
			query = query.Where(fmt.Sprintf("%s BETWEEN ? AND ?", filter.Field), from, to)
		case "before":
			query = query.Where(fmt.Sprintf("%s < ?", filter.Field), filter.Value)
		case "after":
			query = query.Where(fmt.Sprintf("%s > ?", filter.Field), filter.Value)
		default:
			return nil, fmt.Errorf("unsupported filter mode: %s", filter.Mode)
		}

		// Handle preloads
		// if filter.Preload != "" && !contains(fb.Preloads, filter.Preload) {
		// 	fb.Preloads = append(fb.Preloads, filter.Preload)
		// }
	}

	return query, nil
}
