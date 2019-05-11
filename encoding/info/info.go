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
		if tag == "" {
			continue
		}
		if result := strings.Split(tag, ","); len(result) == 2 && result[1] == "keyspace" {
			if err := parseKeySpace(info, v.Elem().Field(i)); err != nil {
				return err
			}
			continue
		}
		value := info[tag]
		if strings.HasPrefix(tag, "cmdstat") {
			if err := parseCommandStat(value, v.Elem().Field(i)); err != nil {
				return err
			}
			continue
		}
		switch field.Type.Kind() {
		case reflect.Int64, reflect.Int:
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

func parseCommandStat(cmdstat string, v reflect.Value) error {
	cmdstats := make(map[string]string)
	for _, line := range strings.Split(cmdstat, ",") {
		if len(line) != 0 {
			result := strings.Split(line, "=")
			cmdstats[result[0]] = result[1]
		}
	}
	v.Set(reflect.ValueOf(cmdstats))
	return nil
}

func parseKeySpace(info map[string]string, v reflect.Value) error {
	keyspace := make(map[string]map[string]int64)
	for key, value := range info {
		if strings.HasPrefix(key, "db") {
			keyspace[key] = make(map[string]int64)
			for _, item := range strings.Split(value, ",") {
				result := strings.Split(item, "=")
				metric, err := strconv.ParseInt(result[1], 10, 64)
				if err != nil {
					return err
				}
				keyspace[key][result[0]] = metric
			}
		}
	}
	v.Set(reflect.ValueOf(keyspace))
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
