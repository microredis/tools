# Tools
Tools for Redis client in go.
Allows you to parse info into valid go structures of client list into golang slice.

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
		if err := m.to.Info().Scan(info); err != nil {
			panic(err)
		}
		if info.MasterLinkStatus == "up" {
			return
		}
	}
}
```
