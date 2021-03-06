package main

import (
    "encoding/xml"
    "os"
    "strings"
)

type document struct {
    Title string `xml:"title"`
    URL   string `xml:"url"`
    Text  string `xml:"abstract"`
    ID    int
}



func LoadDocuments(path string) ([]document, error) {
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

func SearchData(docs []document, term string) []string {
    var r []string
    for _, doc := range docs {
        if strings.Contains(doc.Text, term) {
            r = append(r, doc.Title, doc.URL )
        }
    }
    return r
}



