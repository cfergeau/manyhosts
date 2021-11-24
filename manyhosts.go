package main

import (
	"fmt"
	"os"

	"github.com/code-ready/admin-helper/pkg/hosts"
)

func main() {
	os.Setenv("HOSTS_PATH", "myhosts")
	hosts, err := hosts.New()
	if err != nil {
		panic(err.Error())
	}
	for i := 0; i < 20; i++ {
		err := hosts.Add("127.0.0.1", []string{fmt.Sprintf("foo%d.crc.testing", i)})
		if err != nil {
			panic(err.Error())
		}
	}
}
