# go-pprint
一个能够格式化输出的golang小工具（pretty print)

为了方便开发和调试，能够清晰看到结构体的值:

```golang
package main

import (
	"time"
	"github.com/timest/go-pprint"
)

type Policy struct {
	Allow string
	Deny  string
}

type Host struct {
	ip     string
	alias  []string
	policy Policy
}
type Redis struct {
	name string // lower-case
	port uint   // lower-case
	host Host
}

type Web struct {
	Host    string
	port    int32 // lower-case
	Timeout time.Duration
	Rate    float32
	Score   []float32
	IP      []string
	MySQL struct {
		Name string
		port int64 // lower-case
	}
	redis *Redis
}

func main() {
	w := &Web{
		Host:    "web host",
		port:    1234,
		Timeout: 5 * time.Second,
		Rate:    0.32,
		Score:   []float32{1.1, 2.2, 3.3},
		IP:      []string{"192.168.1.1", "127.0.0.1", "localhost"},
		MySQL: struct {
			Name string
			port int64
		}{Name: "mysqldb", port: 3306},
		redis: &Redis{"rdb", 6379, Host{"adf", []string{"alias1", "alias2"}, Policy{"allow policy", "deny policy"}}},
	}
	// 调用pprint 格式化输出
	pprint.Format(w)
}

```



