package main

import (
	"fmt"
	"github.com/thedevsaddam/gojsonq"
)

func main() {
		
		search := "e"
		res := gojsonq.New().File("./data/test.json").From("system.linux").WhereContains("name", search).Only("download", "name")
		fmt.Printf("%+v\n", res)
}
