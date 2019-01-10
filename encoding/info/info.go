package info

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

	info := parseInfo(string(data))

	return fillObjectFields(info, rv)
}

func fillObjectFields(info map[string]string, v reflect.Value) error {
	for i := 0; i < v.Elem().NumField(); i += 1 {
		field := v.Elem().Type().Field(i)

		tag := field.Tag.Get("info")

		value := info[tag]

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
			reflect.TypeOf(new(int))
			result, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return err
			}
			v.Elem().Field(i).SetFloat(result)
		}
	}
	return nil
}

func parseInfo(data string) map[string]string {
	info := make(map[string]string)

	for _, line := range strings.Split(data, "\r\n") {
		if len(line) != 0 && !strings.HasPrefix(line, "#") {
			result := strings.Split(line, ":")
			if len(result) == 2 {
				info[result[0]] = result[1]
			}
		}
	}

	return info
}