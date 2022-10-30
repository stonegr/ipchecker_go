package base_f

import (
	"fmt"
	"net"
	"time"
)

func Ipcs(host string, port string) (int, string) {
	connTimeout := 1 * time.Second
	// d := net.Dialer{Timeout: connTimeout}
	address := host + ":" + port
	// conn, err := d.Dial("tcp", address)
	conn, err := net.DialTimeout("tcp", address, connTimeout)
	if err != nil {
		fmt.Println(address, " connect faild.")
		return 1, address + " connect faild."
	} else {
		fmt.Println(address, " connect success.")
		conn.Close()
		return 0, address + " connect success."
	}
}

// func New() {
// 	fmt.Println("mypackage.New")
// }
