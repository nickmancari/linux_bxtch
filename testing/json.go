package main

import (
	"fmt"
	"github.com/thedevsaddam/gojsonq"
)

func main() {
		
		search := "fedora"
		res := gojsonq.New().File("../data/test.json").From("system.linux").WhereContains("name", search).Only("download", "name")

		fmt.Sprint("%s\n", res)
}
