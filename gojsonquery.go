package main

import (
	"fmt"

	"github.com/thedevsaddam/gojsonq"
)

func main() {
	jq := gojsonq.New().File("./master.json")
	res := jq.Find("download")
	fmt.Println(res)
}
