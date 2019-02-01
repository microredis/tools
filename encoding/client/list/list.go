package list

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func Unmarshal(data []byte, v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr {
		return fmt.Errorf("non-pointer %v", rv.Type())
	}
	rv = rv.Elem()
	if rv.Kind() != reflect.Slice {
		return fmt.Errorf("invalid type of v %s", rv.Kind())
	}
	clients := strings.Split(string(data), "\r\n")
	rv.Set(reflect.MakeSlice(rv.Type(), len(clients), len(clients)))
	rv.SetCap(len(clients))
	rv.SetLen(len(clients))
	return parseClientList(clients, rv)
}

func parseClientList(clients []string, rv reflect.Value) error {
	for i, line := range clients {
		client := rv.Index(i)
		details := strings.Split(line, " ")
		options := make(map[string]string, len(details))
		for _, item := range details {
			kv := strings.SplitN(item, "=", 2)
			options[kv[0]] = kv[1]
		}
		for i := 0; i < client.NumField(); i++ {
			field := client.Type().Field(i)
			tag := field.Tag.Get("client_list")
			value := options[tag]
			switch field.Type.Kind() {
			case reflect.Int64:
				result, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					return err
				}
				client.Field(i).SetInt(result)
			case reflect.String:
				client.Field(i).SetString(value)
			case reflect.Float64:
				reflect.TypeOf(new(int))
				result, err := strconv.ParseFloat(value, 64)
				if err != nil {
					return err
				}
				client.Field(i).SetFloat(result)
			}
		}
	}
	return nil
}