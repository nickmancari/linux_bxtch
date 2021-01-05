package main

import (
	"fmt"

	"github.com/thedevsaddam/gojsonq"
)

func main() {
	jq := gojsonq.New().File("./master.json")
	res := jq.From("Linux.OS").Where("name", "=", "Debian").Pluck("download")
	fmt.Println(res)
}
