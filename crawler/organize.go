package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {

	titles := OutText("/root/linux_bxtch/crawler/output.txt")

	data := struct {
		Titles []string
	}{
		Titles: titles,
	}

	f, _ := os.Create("/root/linux_bxtch/crawler/template.xml")
	f.Write([]byte(`
<feed>
{{range $v := .Titles}}
  <doc>
    <title>{{$v}}</title>
    <url>Testing</url>
    <abstract>linux, all, Linux, All, ALL, LINUX</abstract>
  </doc>
{{end}}
</feed>`))

	f, _ = os.Open("/root/linux_bxtch/crawler/template.xml")

	CreateTemplate("/root/linux_bxtch/crawler/template.xml", data)
}

func CreateTemplate(s string, m interface{}) {
	t, err := template.ParseFiles(s)
	if err != nil {
		log.Print(err)
		return
	}

	f, err := os.Create(s)
	if err != nil {
		log.Print(err)
		return
	}

	config := m

	err = t.Execute(f, config)
	if err != nil {
		log.Print("Execute:", err)
	}

	f.Close()
}

func OutText(s string) []string {

	file, err := os.Open(s)
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
	return text
}
