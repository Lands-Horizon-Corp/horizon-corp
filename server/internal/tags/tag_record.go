package tags

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/rotisserie/eris"
)

type RecordTag struct{}

func NewRecordTag() *RecordTag {
	return &RecordTag{}
}

func (rt *RecordTag) ExtractFields(obj interface{}) ([]string, []string, error) {
	objValue := reflect.ValueOf(obj)
	if objValue.Kind() != reflect.Ptr || objValue.Elem().Kind() != reflect.Struct {
		return nil, nil, eris.Errorf("expected a pointer to a struct, got %T", obj)
	}

	objValue = objValue.Elem()
	objType := objValue.Type()

	var values []string
	var headers []string

	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		value := objValue.Field(i)

		if tagValue, ok := field.Tag.Lookup("record"); ok && tagValue == "true" {
			header := field.Name
			if jsonTag, ok := field.Tag.Lookup("json"); ok && jsonTag != "" {
				header = strings.Split(jsonTag, ",")[0]
			}
			headers = append(headers, header)
			formattedValue, err := formatFieldValue(value)
			if err != nil {
				return nil, nil, eris.Wrapf(err, "error processing field '%s'", field.Name)
			}
			values = append(values, formattedValue)
		}
	}
	return values, headers, nil
}
func formatFieldValue(value reflect.Value) (string, error) {
	if !value.IsValid() {
		return "N/A", nil
	}

	if value.Kind() == reflect.Ptr {
		if value.IsNil() {
			return "N/A", nil
		}
		value = value.Elem()
	}

	switch value.Kind() {
	case reflect.String:
		return sanitizeCSVField(value.String()), nil
	case reflect.Int, reflect.Int64:
		return strconv.FormatInt(value.Int(), 10), nil
	case reflect.Uint, reflect.Uint64:
		return strconv.FormatUint(value.Uint(), 10), nil
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%.6f", value.Float()), nil
	case reflect.Bool:
		return strconv.FormatBool(value.Bool()), nil
	default:
		return "", eris.Errorf("unsupported field type: %s", value.Kind())
	}
}

func sanitizeCSVField(field string) string {
	problematicPrefixes := []string{"=", "+", "-", ""}
	for _, prefix := range problematicPrefixes {
		if strings.HasPrefix(field, prefix) {
			return "'" + field
		}
	}
	return field
}
