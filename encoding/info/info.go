package info

import (
	"bytes"
	"encoding/json"
	"strings"
)

var newLine = []byte("\r\n")

var sharp = []byte("#")

func Unmarshal(data []byte, v interface{}) error {
	as := map[string]map[string]interface{}{
		"cmdstats": make(map[string]interface{}),
		"keyspace": make(map[string]interface{}),
		"stats":    make(map[string]interface{}),
	}
	for _, line := range bytes.Split(data, newLine) {
		if len(line) != 0 && !bytes.HasPrefix(line, sharp) {
			result := strings.SplitN(string(line), ":", 2)
			if strings.HasPrefix(result[0], "cmdstat") {
				kvp := make(map[string]string)
				for _, stat := range strings.Split(result[1], ",") {
					kv := strings.SplitN(stat, "=", 2)
					kvp[kv[0]] = kv[1]
				}
				as["cmdstats"][result[0]] = kvp
			} else if strings.HasPrefix(result[0], "db") {
				kvp := make(map[string]string)
				for _, keyspace := range strings.Split(result[1], ",") {
					kv := strings.SplitN(keyspace, "=", 2)
					kvp[kv[0]] = kv[1]
				}
				as["keyspace"][result[0]] = kvp
			} else {
				as["stats"][result[0]] = result[1]
			}
		}
	}
	info, err := json.Marshal(as)
	if err != nil {
		return err
	}
	return json.Unmarshal(info, v)
}
