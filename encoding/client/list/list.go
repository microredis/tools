package list

import (
	"bytes"
	"encoding/json"
	"strings"
)

var newLine = []byte("\r\n")

var space = []byte(" ")

func Unmarshal(data []byte, v interface{}) error {
	clients := bytes.Split(data, newLine)
	list := make([]map[string]string, len(clients))
	for i := range clients {
		details := bytes.Split(clients[i], space)
		options := make(map[string]string, len(details))
		for i := range details {
			kv := strings.SplitN(string(details[i]), "=", 2)
			options[kv[0]] = kv[1]
		}
		list[i] = options
	}
	result, err := json.Marshal(list)
	if err != nil {
		return err
	}
	return json.Unmarshal(result, v)
}
