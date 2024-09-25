package repository

func isValidField(field string) bool {
	// Example validation: field contains only letters, numbers, and underscores
	for _, r := range field {
		if !((r >= 'a' && r <= 'z') ||
			(r >= 'A' && r <= 'Z') ||
			(r >= '0' && r <= '9') ||
			r == '_') {
			return false
		}
	}
	return true
}

func (op Operator) IsValid() bool {
	switch op {
	case OpEquals, OpNotEquals, OpGreaterThan, OpGreaterThanOrEqual,
		OpLessThan, OpLessThanOrEqual, OpRange, OpStartsWith,
		OpEndsWith, OpIncludes, OpBooleanEquals,
		OpIn, OpNotIn, OpIsNull, OpIsNotNull:
		return true
	default:
		return false
	}
}
