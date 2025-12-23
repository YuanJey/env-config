package load

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func LoadEnv(config interface{}) error {
	val := reflect.ValueOf(config).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		// Check if the field is a struct and need to recursively process it.
		if field.Kind() == reflect.Struct {
			// Recursively call LoadEnv for nested structs
			if err := LoadEnv(field.Addr().Interface()); err != nil {
				return err
			}
			continue
		}
		envTag := fieldType.Tag.Get("env")
		defTag := fieldType.Tag.Get("def")

		if envTag == "" && defTag == "" {
			continue
		}

		var envValue string
		if envTag != "" {
			envValue = os.Getenv(envTag)
		}
		if envValue == "" {
			envValue = defTag
		}

		if envValue == "" {
			continue
		}

		switch field.Kind() {
		case reflect.String:
			field.SetString(envValue)
		case reflect.Slice:
			// Assuming we want to handle it as a slice of strings
			if field.Type().Elem().Kind() == reflect.String {
				parts := strings.Split(envValue, ",")
				slice := reflect.MakeSlice(field.Type(), len(parts), len(parts))
				for i, part := range parts {
					slice.Index(i).SetString(part)
				}
				field.Set(slice)
			}
		case reflect.Int:
			intValue, err := strconv.Atoi(envValue)
			if err != nil {
				return fmt.Errorf("invalid value for %s: %v", envTag, err)
			}
			field.SetInt(int64(intValue))
		case reflect.Bool:
			boolValue, err := strconv.ParseBool(envValue)
			if err != nil {
				return fmt.Errorf("invalid value for %s: %v", envTag, err)
			}
			field.SetBool(boolValue)
		default:
			return fmt.Errorf("unsupported type for %s", envTag)
		}
	}
	return nil
}
