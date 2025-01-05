package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

func TransformToStruct[T any](data interface{}) *T {
	resourceMap := transformToMap[any](data)
	fmt.Printf("Resource Map: %+v\n", resourceMap)

	jsonBytes, err := json.Marshal(resourceMap)
	if err != nil {
		panic(fmt.Errorf("failed to marshal resource map: %w", err))
	}
	fmt.Printf("Intermediate JSON: %s\n", string(jsonBytes))

	var resource T
	err = json.Unmarshal(jsonBytes, &resource)
	if err != nil {
		panic(fmt.Errorf("failed to unmarshal into resource struct: %w", err))
	}

	return &resource
}

func transformToMap[T any](data interface{}) map[string]interface{} {
	tagKey := "resource"
	if data == nil {
		return nil
	}

	resource := make(map[string]interface{})
	v := reflect.ValueOf(data)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil
	}

	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		fmt.Printf("Transforming Field: %s, Value: %+v\n", field.Name, value.Interface())

		tagValue := field.Tag.Get(tagKey)
		if tagValue == "" {
			continue
		}

		if value.Kind() == reflect.Ptr {
			if value.IsNil() {
				resource[tagValue] = nil
			} else if value.Elem().Kind() == reflect.Struct {
				fmt.Printf("Recursively transforming nested struct for field: %s\n", field.Name)
				nestedMap := transformToMap[T](value.Elem().Interface())
				fmt.Printf("Nested Map for %s: %+v\n", field.Name, nestedMap)
				resource[tagValue] = nestedMap
			} else {
				resource[tagValue] = value.Elem().Interface()
			}
			continue
		}

		if value.Type() == reflect.TypeOf(time.Time{}) {
			resource[tagValue] = value.Interface().(time.Time).Format(time.RFC3339)
			continue
		}

		if value.Kind() == reflect.Slice {
			slice := []interface{}{}
			for j := 0; j < value.Len(); j++ {
				item := value.Index(j).Interface()
				if reflect.TypeOf(item).Kind() == reflect.Struct || (reflect.TypeOf(item).Kind() == reflect.Ptr && reflect.TypeOf(item).Elem().Kind() == reflect.Struct) {
					slice = append(slice, transformToMap[T](item))
				} else {
					slice = append(slice, item)
				}
			}
			resource[tagValue] = slice
			continue
		}

		if value.Kind() == reflect.Struct {
			nestedMap := transformToMap[T](value.Interface())
			fmt.Printf("Transforming nested struct for field: %s\n", field.Name)
			if nestedMap != nil {
				resource[tagValue] = nestedMap
			}
			continue
		}

		resource[tagValue] = value.Interface()
	}

	return resource
}

type Footstep struct {
	Action    string `resource:"action"`
	Timestamp string `resource:"timestamp"`
}

type Media struct {
	URL string `json:"url" resource:"url"`
}

type Admin struct {
	FirstName string      `json:"first_name" gorm:"type:varchar(255);not null" resource:"firstName" validate:"required,max=255"`
	LastName  string      `resource:"lastName"`
	Footsteps []*Footstep `resource:"footsteps"`
	CreatedAt time.Time   `resource:"createdAt"`
	Media     *Media      `resource:"media"`
}

func main() {
	admin := &Admin{
		FirstName: "John",
		LastName:  "Doe",
		Footsteps: []*Footstep{
			{Action: "Login", Timestamp: time.Now().Format(time.RFC3339)},
			{Action: "Logout", Timestamp: time.Now().Format(time.RFC3339)},
		},
		CreatedAt: time.Now(),
		Media:     &Media{URL: "https://example.com/image.png"},
	}

	resource := TransformToStruct[Admin](admin)

	fmt.Printf("First Name: %s\n", resource.FirstName)
	fmt.Printf("Last Name: %s\n", resource.LastName)
	if resource.Media != nil {
		fmt.Printf("Media URL: %s\n", resource.Media.URL)
	} else {
		fmt.Println("Media is nil.")
	}
}
