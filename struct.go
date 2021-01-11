package main

import (
	"fmt"
	"github.com/thedevsaddam/gojsonq"
)

func main() {
		type item struct {
			Name string `json:"name"`
			Download string `json:"download"`
		}
		v:= item{}
		
		gojsonq.New().File("./data/test.json").From("system.linux.[3]").Out(&v)
		fmt.Printf("%+v\n", v)
}
