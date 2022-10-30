package main

import (
	"flag"
	"ipcs/base_f"
	"strconv"

	"github.com/gin-gonic/gin"
)

func start(host string, port int, secreat []interface{}) {

	// gin.SetMode(gin.DebugMode)
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	type User struct {
		Status int    `json:"status"` // 通过json标签定义struct字段转换成json字段的名字。
		Msg    string `json:"msg"`
	}
	r.GET("/", func(c *gin.Context) {
		secreat_ := c.Query("s")
		// fmt.Println(secreat_)
		access := base_f.In(secreat_, secreat)
		// fmt.Println(bo)
		if access {
			host := c.Query("host")
			if host != "" {
				if base_f.Test_ip(host) {
				} else {
					_host := base_f.Get_ip(host)
					if base_f.Test_ip(_host) {
						host = _host
					} else {
						c.String(404, "404")
						return
					}
				}
				port := c.DefaultQuery("port", "80")
				cs, msg := base_f.Ipcs(host, port)
				// cs := base_f.Ipcs("114.114.114.114", "53")

				u := &User{
					Status: cs,
					Msg:    msg,
				}
				c.JSON(200, u)
			} else {
				c.String(404, "404")
			}
		} else {
			// u := &User{
			// 	Status: 1,
			// 	Msg:    "secreat is wrong.",
			// }
			// c.JSON(200, u)
			c.String(404, "404")

		}
	})

	r.Run(host + ":" + strconv.Itoa(port)) // listen and serve on 0.0.0.0:8080
	// s := &http.Server{
	// 	Addr:           host + ":" + "8081",
	// 	Handler:        r,
	// 	ReadTimeout:    10 * time.Second,
	// 	WriteTimeout:   10 * time.Second,
	// 	MaxHeaderBytes: 1 << 20,
	// }
	// s.ListenAndServe()
}

func main() {
	var config string
	flag.StringVar(&config, "c", "", "配置文件")
	flag.Parse()
	if flag.NFlag() == 0 {
		flag.PrintDefaults()
	} else {
		// fmt.Println(config)
		// base_f.r_json(config)
		cs := base_f.R_json(config)
		// cs := base_f.R_json("./config.json")
		// fmt.Println(cs.Host)
		// fmt.Println(cs.Port)
		// fmt.Println(cs.Secreat)
		// fmt.Printf("%T", cs.Secreat)
		// fmt.Println(cs.Secreat[0])

		start(cs.Host, cs.Port, cs.Secreat)
	}
}
