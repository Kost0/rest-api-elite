package pkg

import (
	"reflect"
)

type TableView struct {
	Title     string
	Data      interface{}
	GetKeys   func(interface{}) []string
	GetValues func(interface{}) []interface{}
}

func GetKeys(obj interface{}) []string {
	var keys []string
	val := reflect.ValueOf(obj)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return nil
	}

	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		jsonTag := field.Tag.Get("json")
		if jsonTag != "" {
			keys = append(keys, jsonTag)
		} else {
			keys = append(keys, field.Name)
		}
	}
	return keys
}

func GetValues(obj interface{}) []interface{} {
	var values []interface{}
	val := reflect.ValueOf(obj)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return nil
	}

	for i := 0; i < val.NumField(); i++ {
		fieldValue := val.Field(i)
		values = append(values, fieldValue.Interface())
	}
	return values
}
