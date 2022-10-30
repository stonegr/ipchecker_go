package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	dns := "qq.c"
	// 解析ip地址
	ns, err := net.LookupHost(dns)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Err: %s", err.Error())
	} else {
		fmt.Println(ns[0])
	}
}
