package repositories

import (
	"fmt"
	"time"
)

type Filter struct {
	Column   string `json:"column"`
	Mode     string `json:"mode"`
	DataType string `json:"dataType"`
	Value    string `json:"value,omitempty"`
	From     string `json:"from,omitempty"`
	To       string `json:"to,omitempty"`
}

type FilteredResult[T any] struct {
	Data      []*T       `json:"data"`
	PageIndex int        `json:"pageIndex"`
	TotalPage int        `json:"totalPage"`
	PageSize  int        `json:"pageSize"`
	TotalSize int64      `json:"totalSize"`
	Pages     []PageLink `json:"pages"`
}

type PageLink struct {
	Page      string `json:"page"`
	PageIndex int    `json:"pageIndex"`
}

func (r *Repository[T]) ApplyFilters(filters []Filter, preloads []string, pageIndex, pageSize int) (*FilteredResult[T], error) {
	var data []*T
	var totalSize int64

	db := r.DB
	for _, preload := range preloads {
		db = db.Preload(preload)
	}

	for _, filter := range filters {
		query, args := buildQuery(filter)
		db = db.Where(query, args...)
	}

	if err := db.Model(new(T)).Count(&totalSize).Error; err != nil {
		return nil, err
	}

	// Paginate
	offset := (pageIndex - 1) * pageSize
	if err := db.Limit(pageSize).Offset(offset).Find(&data).Error; err != nil {
		return nil, err
	}

	// Construct pagination metadata
	totalPage := int((totalSize + int64(pageSize) - 1) / int64(pageSize))
	pages := make([]PageLink, totalPage)
	for i := 1; i <= totalPage; i++ {
		pages[i-1] = PageLink{
			Page:      fmt.Sprintf("/api/data?pageIndex=%d", i),
			PageIndex: i,
		}
	}

	return &FilteredResult[T]{
		Data:      data,
		PageIndex: pageIndex,
		TotalPage: totalPage,
		PageSize:  pageSize,
		TotalSize: totalSize,
		Pages:     pages,
	}, nil
}

// Helper function to build query for a filter
func buildQuery(filter Filter) (string, []interface{}) {
	column := filter.Column
	switch filter.DataType {
	case "text":
		return buildTextQuery(column, filter)
	case "number":
		return buildNumberQuery(column, filter)
	case "date":
		return buildDateQuery(column, filter)
	default:
		return "", nil
	}
}

func buildTextQuery(column string, filter Filter) (string, []interface{}) {
	switch filter.Mode {
	case "equal":
		return fmt.Sprintf("%s = ?", column), []interface{}{filter.Value}
	case "nequal":
		return fmt.Sprintf("%s != ?", column), []interface{}{filter.Value}
	case "contains":
		return fmt.Sprintf("%s LIKE ?", column), []interface{}{"%" + filter.Value + "%"}
	case "ncontains":
		return fmt.Sprintf("%s NOT LIKE ?", column), []interface{}{"%" + filter.Value + "%"}
	case "startswith":
		return fmt.Sprintf("%s LIKE ?", column), []interface{}{filter.Value + "%"}
	case "endswith":
		return fmt.Sprintf("%s LIKE ?", column), []interface{}{"%" + filter.Value}
	case "isempty":
		return fmt.Sprintf("%s = ''", column), nil
	case "isnotempty":
		return fmt.Sprintf("%s != ''", column), nil
	default:
		return "", nil
	}
}

func buildNumberQuery(column string, filter Filter) (string, []interface{}) {
	switch filter.Mode {
	case "equal":
		return fmt.Sprintf("%s = ?", column), []interface{}{filter.Value}
	case "nequal":
		return fmt.Sprintf("%s != ?", column), []interface{}{filter.Value}
	case "gt":
		return fmt.Sprintf("%s > ?", column), []interface{}{filter.Value}
	case "gte":
		return fmt.Sprintf("%s >= ?", column), []interface{}{filter.Value}
	case "lt":
		return fmt.Sprintf("%s < ?", column), []interface{}{filter.Value}
	case "lte":
		return fmt.Sprintf("%s <= ?", column), []interface{}{filter.Value}
	case "range":
		return fmt.Sprintf("%s BETWEEN ? AND ?", column), []interface{}{filter.From, filter.To}
	default:
		return "", nil
	}
}

func buildDateQuery(column string, filter Filter) (string, []interface{}) {
	from, _ := time.Parse(time.RFC3339, filter.From)
	to, _ := time.Parse(time.RFC3339, filter.To)
	switch filter.Mode {
	case "equal":
		return fmt.Sprintf("DATE(%s) = ?", column), []interface{}{from}
	case "nequal":
		return fmt.Sprintf("DATE(%s) != ?", column), []interface{}{from}
	case "before":
		return fmt.Sprintf("DATE(%s) < ?", column), []interface{}{from}
	case "after":
		return fmt.Sprintf("DATE(%s) > ?", column), []interface{}{from}
	case "gte":
		return fmt.Sprintf("DATE(%s) >= ?", column), []interface{}{from}
	case "lte":
		return fmt.Sprintf("DATE(%s) <= ?", column), []interface{}{to}
	case "range":
		return fmt.Sprintf("DATE(%s) BETWEEN ? AND ?", column), []interface{}{from, to}
	default:
		return "", nil
	}
}
