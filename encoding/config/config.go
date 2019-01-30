package config

import (
	"strings"
)

type Configuration struct {
	Requirepass   string
	RenameCommand map[string]string
}

func New(data []byte) Configuration {
	c := Configuration{
		RenameCommand: make(map[string]string),
	}
	for _, line := range strings.Split(string(data), "\n") {
		if strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "" {
			continue
		}
		var (
			argv    = strings.Split(line, " ")
			argc    = len(argv)
			command = argv[0]
		)
		if strings.EqualFold(command, "requirepass") && argc == 2 {
			c.Requirepass = argv[1]
		} else if strings.EqualFold(command, "rename-command") && argc == 3 {
			c.RenameCommand[argv[1]] = argv[2]
		}
	}
	return c
}
