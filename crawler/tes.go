package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
)

func main() {

        file, err := os.Open("/root/linux_bxtch/crawler/output.txt")
        if err != nil {
                log.Print(err)
        }

        scanner := bufio.NewScanner(file)

        scanner.Split(bufio.ScanLines)
        var text []string

        for scanner.Scan() {
                text = append(text, scanner.Text())
        }

        file.Close()

        for _, each_ln := range text {
                fmt.Println(each_ln)
        }

}
