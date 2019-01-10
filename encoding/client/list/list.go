package list

import (
	"strings"
	"reflect"
	"strconv"
)

type Client struct {

}

func Unmarshal(data []byte, v interface{}) {
	list := make(map[string]interface{})
	for _, line := range strings.Split(string(data), "\r\n") {

		for _, item := range strings.Split(client, " ") {
			kv := strings.SplitN(item, "=", 2)
			result[kv[0]] = getClientValue(kv[1])
		}
	}

	result := make(map[string]interface{})

	return result

	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return errors.New("invalid type of v")
	}
	if !rv.IsValid() {
		return errors.New("")
	}
	for i := 0; i < rv.Elem().NumField(); i++ {
		field := rv.Type().Elem().Field(i)

		tag := field.Tag.Get("info")

		key := info[tag]

		switch field.Type.Kind() {
		case reflect.Int64:
			result, err := strconv.ParseInt(key, 10, 64)
			if err != nil {
				return err
			}
			rv.Elem().Field(i).SetInt(result)
		case reflect.String:
			rv.Elem().Field(i).SetString(key)
		case reflect.Float64:
			reflect.TypeOf(new(int))
			result, err := strconv.ParseFloat(key, 64)
			if err != nil {
				return err
			}
			rv.Elem().Field(i).SetFloat(result)
		}
	}
	return nil
}