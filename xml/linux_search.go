package main

import (
	"github.com/antchfx/xmlquery"
	"fmt"
	"os"
)

func main() {

	f, err := os.Open("/root/linux_bxtch/xml/linux.xml")
	if err != nil {
		fmt.Println(err)
	}

	doc, _ := xmlquery.Parse(f)

	s := xmlquery.FindOne(doc, "//os")

	if n := s.SelectElement("link"); n != nil {
		fmt.Printf("Debian: %s\n", n.InnerText())
	}
}
