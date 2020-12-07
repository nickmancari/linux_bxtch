package main

import (
	
	"fmt"

	"github.com/thedevsaddam/gojsonq" 
)

func main() {
	jq := gojsonq.New().File("./linux.json")
	res := jq.From("os").Where("name", "=", "Fedora").Only("download")
	fmt.Println(res)

}
