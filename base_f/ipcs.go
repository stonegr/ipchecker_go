package base_f

import (
	"fmt"
	"net"
	"os"
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

// 检测ip地址是否合法
func Test_ip(ip string) bool {
	// ParseIP 这个方法 可以用来检查 ip 地址是否正确，如果不正确，该方法返回 nil
	address := net.ParseIP(ip)
	if address == nil {
		return false
	} else {
		return true
	}

}

// 解析ip地址
func Get_ip(domain string) string {

	// 解析ip地址
	ns, err := net.LookupHost(domain)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Err: %s", err.Error())
		return ""
	} else {
		// fmt.Println(ns[0])
		return ns[0]
	}
}

// func New() {
// 	fmt.Println("mypackage.New")
// }
