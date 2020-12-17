//Testing db files in Go

package main

import (
	
	"encoding/json"
	"os"
	"fmt"

)

func main() {
	jsonFile, err := os.Open("db/Cubes.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened Cubes.json")
	
	defer jsonFile.Close()
	
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var 
	
}
