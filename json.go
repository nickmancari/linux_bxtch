package main

import (
	"fmt"
	"github.com/thedevsaddam/gojsonq"
)

func main() {

		res := gojsonq.New().File("./data/test.json").From("system.linux").WhereContains("name", "ar").Pluck("name", "download")
		fmt.Println(res)
}
