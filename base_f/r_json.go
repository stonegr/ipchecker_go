package base_f

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	Host    string        `json:"host"`
	Port    int           `json:"port"`
	Secreat []interface{} `json:"secreat"`
}

func R_json(path string) Config {
	var config Config

	// 打开json文件
	jsonFile, err := os.Open(path)

	// 最好要处理以下错误
	if err != nil {
		fmt.Println(err)
	}

	// 要记得关闭
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	// fmt.Println(string(byteValue))

	json.Unmarshal([]byte(byteValue), &config)
	// fmt.Println(config)
	// fmt.Printf("%T", config)

	// fmt.Println(config.Host)
	return config
}

func In(target string, str_array []interface{}) bool {
	for _, element := range str_array {
		if target == element {
			return true
		}
	}
	return false
}
