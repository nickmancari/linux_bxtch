package main

import (
	"fmt"

	"github.com/thedevsaddam/gojsonq"
)

func main() {
	jq := gojsonq.New().Folder("./db")
	res := jq.From("vendor.items").SortBy("price").Get()
	fmt.Println(res)
