package repository

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

var OperatorHandlers = map[Operator]OperatorHandler{
	OpEquals:             handleEquals,
	OpNotEquals:          handleNotEquals,
	OpGreaterThan:        handleGreaterThan,
	OpGreaterThanOrEqual: handleGreaterThanOrEqual,
	OpLessThan:           handleLessThan,
	OpLessThanOrEqual:    handleLessThanOrEqual,
	OpRange:              handleRange,
	OpStartsWith:         handleStartsWith,
	OpEndsWith:           handleEndsWith,
	OpIncludes:           handleIncludes,
	OpBooleanEquals:      handleBooleanEquals,
	OpIn:                 handleIn,
	OpNotIn:              handleNotIn,
	OpIsNull:             handleIsNull,
	OpIsNotNull:          handleIsNotNull,
}

func handleEquals(query *gorm.DB, filter Filter) (*gorm.DB, error) {
	return query.Where(fmt.Sprintf("%s = ?", filter.Field), filter.Value), nil
}

func handleNotEquals(query *gorm.DB, filter Filter) (*gorm.DB, error) {
	return query.Where(fmt.Sprintf("%s != ?", filter.Field), filter.Value), nil
}

func handleGreaterThan(query *gorm.DB, filter Filter) (*gorm.DB, error) {
	return query.Where(fmt.Sprintf("%s > ?", filter.Field), filter.Value), nil
}

func handleGreaterThanOrEqual(query *gorm.DB, filter Filter) (*gorm.DB, error) {
	return query.Where(fmt.Sprintf("%s >= ?", filter.Field), filter.Value), nil
}

func handleLessThan(query *gorm.DB, filter Filter) (*gorm.DB, error) {
	return query.Where(fmt.Sprintf("%s < ?", filter.Field), filter.Value), nil
}

func handleLessThanOrEqual(query *gorm.DB, filter Filter) (*gorm.DB, error) {
	return query.Where(fmt.Sprintf("%s <= ?", filter.Field), filter.Value), nil
}

func handleRange(query *gorm.DB, filter Filter) (*gorm.DB, error) {
	values, ok := filter.Value.([]interface{})
	if !ok || len(values) != 2 {
		return nil, errors.New("invalid value for 'range' operator; expected slice with two elements")
	}
	return query.Where(fmt.Sprintf("%s BETWEEN ? AND ?", filter.Field), values[0], values[1]), nil
}

func handleStartsWith(query *gorm.DB, filter Filter) (*gorm.DB, error) {
	valueStr, ok := filter.Value.(string)
	if !ok {
		return nil, errors.New("invalid value for 'starts_with' operator; expected string")
	}
	return query.Where(fmt.Sprintf("%s LIKE ?", filter.Field), fmt.Sprintf("%s%%", valueStr)), nil
}

func handleEndsWith(query *gorm.DB, filter Filter) (*gorm.DB, error) {
	valueStr, ok := filter.Value.(string)
	if !ok {
		return nil, errors.New("invalid value for 'ends_with' operator; expected string")
	}
	return query.Where(fmt.Sprintf("%s LIKE ?", filter.Field), fmt.Sprintf("%%%s", valueStr)), nil
}

func handleIncludes(query *gorm.DB, filter Filter) (*gorm.DB, error) {
	valueStr, ok := filter.Value.(string)
	if !ok {
		return nil, errors.New("invalid value for 'includes' operator; expected string")
	}
	return query.Where(fmt.Sprintf("%s LIKE ?", filter.Field), fmt.Sprintf("%%%s%%", valueStr)), nil
}

func handleBooleanEquals(query *gorm.DB, filter Filter) (*gorm.DB, error) {
	boolVal, ok := filter.Value.(bool)
	if !ok {
		return nil, errors.New("invalid value for 'boolean_equals' operator; expected boolean")
	}
	return query.Where(fmt.Sprintf("%s = ?", filter.Field), boolVal), nil
}

func handleIn(query *gorm.DB, filter Filter) (*gorm.DB, error) {
	values, ok := filter.Value.([]interface{})
	if !ok || len(values) == 0 {
		return nil, errors.New("invalid value for 'in' operator; expected a non-empty slice")
	}
	return query.Where(fmt.Sprintf("%s IN ?", filter.Field), values), nil
}

func handleNotIn(query *gorm.DB, filter Filter) (*gorm.DB, error) {
	values, ok := filter.Value.([]interface{})
	if !ok || len(values) == 0 {
		return nil, errors.New("invalid value for 'not_in' operator; expected a non-empty slice")
	}
	return query.Where(fmt.Sprintf("%s NOT IN ?", filter.Field), values), nil
}

func handleIsNull(query *gorm.DB, filter Filter) (*gorm.DB, error) {
	return query.Where(fmt.Sprintf("%s IS NULL", filter.Field)), nil
}

func handleIsNotNull(query *gorm.DB, filter Filter) (*gorm.DB, error) {
	return query.Where(fmt.Sprintf("%s IS NOT NULL", filter.Field)), nil
}
