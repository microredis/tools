# Tools
[![Maintainability](https://api.codeclimate.com/v1/badges/158f8ef3eedf42d72c5b/maintainability)](https://codeclimate.com/github/microredis/tools/maintainability) [![Build Status](https://travis-ci.org/microredis/tools.svg?branch=master)](https://travis-ci.org/microredis/tools)

Tools for Redis client in go.
Allows you to parse info into or client list valid go structures of client list into golang slice.
Supports ```int64``` and ```string``` as well as ```map[string]map[string]int``` for Keyspace section of info command and more.

```go
package main

import "github.com/microredis/tools/encoding/info"

type Info struct {
	MasterLinkStatus string `info:"master_link_status"`
}

func (i *Info) UnmarshalBinary(data []byte) error {
	return info.Unmarshal(data, i)
}
```

```go
func (m *Migrator) waitForUp() {
	info := new(Info)
	for {
		if err := m.redisClient.Info().Scan(info); err != nil {
			panic(err)
		}
		if info.MasterLinkStatus == "up" {
			return
		}
	}
}
```
