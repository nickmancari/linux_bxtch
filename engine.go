package main

import (
	"fmt"
	"strings"
	"os"
	"encoding/xml"

)

type document struct {
	Title string `xml:"Title"`
	Artist string `xml:"Artist"`
	Text string `xml:"Company"`
	ID	int
}

func main() {

	q, _ := loadDocuments("/root/linux_bxtch/data/test.xml")

	res := search(q, "CBS")
	fmt.Println(res)
}

func loadDocuments(path string) ([]document, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	dec := xml.NewDecoder(f)
	dump := struct {
		Documents []document `xml:"doc"`
	}{}
	if err := dec.Decode(&dump); err != nil {
		return nil, err
	}

	docs := dump.Documents
	for i := range docs {
		docs[i].ID = i
	}
	return docs, nil
}

func search(docs []document, term string) []document {
	var r []document
	for _, doc := range docs {
		if strings.Contains(doc.Text, term) {
			r = append(r, doc)
		}
	}
	return r
}

