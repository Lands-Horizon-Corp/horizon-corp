package tags

import (
	"encoding/json"
	"reflect"
	"time"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/models"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/rotisserie/eris"
)

type ResourceTag[T any] struct {
	storage *providers.StorageProvider
	tagKey  string
}

func NewResourceTag[T any](storage *providers.StorageProvider, tagKey string) *ResourceTag[T] {
	if tagKey == "" {
		tagKey = "resource"
	}
	return &ResourceTag[T]{storage: storage, tagKey: tagKey}
}

func (rt *ResourceTag[T]) TransformToStruct(data interface{}) (*T, error) {
	resourceMap, err := rt.transformToMap(data)
	if err != nil {
		return nil, eris.Wrap(err, "failed to transform data to map")
	}

	resource := new(T)
	if err := mapToStruct(resourceMap, resource); err != nil {
		return nil, eris.Wrap(err, "failed to map resource data")
	}

	return resource, nil
}

func mapToStruct(data map[string]interface{}, target interface{}) error {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonBytes, target)
}

func (rt *ResourceTag[T]) transformToMap(data interface{}) (map[string]interface{}, error) {
	if data == nil {
		return nil, nil
	}

	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, eris.New("data is not a struct or pointer to a struct")
	}

	resource := make(map[string]interface{})
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		tagValue := field.Tag.Get(rt.tagKey)
		if tagValue == "" || !value.IsValid() {
			continue
		}

		processedValue, err := rt.processField(value, field)
		if err != nil {
			return nil, eris.Wrapf(err, "error processing field: %s", field.Name)
		}
		resource[tagValue] = processedValue
	}
	return resource, nil
}

func (rt *ResourceTag[T]) processField(value reflect.Value, field reflect.StructField) (interface{}, error) {
	if value.Kind() == reflect.Ptr {
		if value.IsNil() {
			return nil, nil
		}
		value = value.Elem()
	}

	switch value.Kind() {
	case reflect.Struct:
		if value.Type() == reflect.TypeOf(models.Media{}) {
			return rt.processMedia(value.Interface().(models.Media))
		}
		if value.Type() == reflect.TypeOf(time.Time{}) {
			return value.Interface().(time.Time).Format(time.RFC3339), nil
		}
		return rt.transformToMap(value.Interface())
	case reflect.Slice:
		return rt.processSlice(value, field)
	default:
		return value.Interface(), nil
	}
}

func (rt *ResourceTag[T]) processMedia(media models.Media) (map[string]interface{}, error) {
	url, err := rt.storage.GeneratePresignedURL(media.StorageKey)
	if err != nil {
		return nil, eris.Wrap(err, "failed to generate presigned URL")
	}

	mediaData, err := json.Marshal(media)
	if err != nil {
		return nil, eris.Wrap(err, "failed to marshal media")
	}

	var mediaMap map[string]interface{}
	if err := json.Unmarshal(mediaData, &mediaMap); err != nil {
		return nil, eris.Wrap(err, "failed to unmarshal media")
	}
	mediaMap["downloadUrl"] = url
	return mediaMap, nil
}

func (rt *ResourceTag[T]) processSlice(value reflect.Value, field reflect.StructField) ([]interface{}, error) {
	slice := make([]interface{}, value.Len())
	for i := 0; i < value.Len(); i++ {
		item := value.Index(i)
		processedItem, err := rt.processField(item, field)
		if err != nil {
			return nil, eris.Wrapf(err, "failed to process slice item for field: %s", field.Name)
		}
		slice[i] = processedItem
	}
	return slice, nil
}
