package list

import (
	"strings"
	"testing"
)

type Client struct {
	Id   int64  `json:"id,string"`
	Name string `json:"name"`
	Age  int64  `json:"age,string"`
}

func TestUnmarshal(t *testing.T) {
	data := strings.Join(strings.Split(data, "\n"), "\r\n")
	var clients []Client
	if err := Unmarshal([]byte(data), &clients); err != nil {
		t.Fatal(err.Error())
	}
	if clients[0].Id != 3 || clients[1].Id != 3 {
		t.Fatal("fuck")
	}
	if clients[0].Name != "" || clients[1].Name != "" {
		t.Fatal("fuck")
	}
	if clients[0].Age != 6 || clients[1].Age != 6 {
		t.Fatal("fuck")
	}
}

const data = `id=3 addr=172.17.0.1:42244 fd=8 name= age=6 idle=0 flags=N db=0 sub=0 psub=0 multi=-1 qbuf=26 qbuf-free=32742 obl=0 oll=0 omem=0 events=r cmd=client
id=3 addr=172.17.0.1:42244 fd=8 name= age=6 idle=0 flags=N db=0 sub=0 psub=0 multi=-1 qbuf=26 qbuf-free=32742 obl=0 oll=0 omem=0 events=r cmd=client`
