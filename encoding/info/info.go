package info

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

func Unmarshal(data []byte, v interface{}) error {
	info := make(map[string]string)

	for _, line := range strings.Split(string(data), "\r\n") {
		if len(line) != 0 && !strings.HasPrefix(line, "#") {
			result := strings.Split(line, ":")
			if len(result) == 2 {
				info[result[0]] = result[1]
			}
		}
	}

	rv := reflect.ValueOf(v)

	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return errors.New("invalid type of v")
	}

	if !rv.IsValid() {
		return errors.New("")
	}

	for i := 0; i < rv.Elem().NumField(); i++ {
		field := rv.Elem().Type().Field(i)

		tag := field.Tag.Get("info")

		value := info[tag]

		switch field.Type.Kind() {
		case reflect.Int64:
			result, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return err
			}
			rv.Elem().Field(i).SetInt(result)
		case reflect.String:
			rv.Elem().Field(i).SetString(value)
		case reflect.Float64:
			reflect.TypeOf(new(int))
			result, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return err
			}
			rv.Elem().Field(i).SetFloat(result)
		}
	}

	return nil
}