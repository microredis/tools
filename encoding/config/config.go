package config

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

func Unmarshal(data []byte, v interface{}) error {
	rv := reflect.ValueOf(v)

	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return errors.New("invalid type of v")
	}

	if !rv.IsValid() {
		return errors.New("invalid")
	}
	input := string(data)
	config := parseConfig(input)

	return fillObjectFields(config, rv)
}

func fillObjectFields(config map[string]string, v reflect.Value) error {
	for i := 0; i < v.Elem().NumField(); i += 1 {
		field := v.Elem().Type().Field(i)

		tag := field.Tag.Get("config")

		if tag == "" {
			continue
		}

		value := config[tag]

		switch field.Type.Kind() {
		case reflect.Int64:
			result, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return err
			}
			v.Elem().Field(i).SetInt(result)
		case reflect.String:
			v.Elem().Field(i).SetString(value)
		case reflect.Float64:
			result, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return err
			}
			v.Elem().Field(i).SetFloat(result)
		}
	}
	return nil
}

func parseConfig(data string) map[string]string {
	result := make(map[string]string)

	for _, i := range strings.Split(data, "\n") {
		if strings.HasPrefix(i, "#") || strings.TrimSpace(i) == "" {
			continue
		}
		if kv := strings.SplitN(i, " ", 2); len(kv) == 2 {
			result[kv[0]] = strings.TrimSpace(kv[1])
		}
	}

	return result
}
