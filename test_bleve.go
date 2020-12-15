//Testing Bleve indexing and querying with JSON files in database

package main

import (
	"testing"
	"io/ioutil"
	"encoding/json"
	"path/filepath"

	"github.com/blevesearch/bleve"
)

func TestBleveSearch( t *testing.T) {
	
	mapping, err := buildIndexMapping()
	if err != nil {
		t.Fatal(err)
	}
	index, err := bleve.New("search.bleve", mapping)
	if err != nil {
		t.Fatal(err)
	}
	defer index.Close()
		

	// open directory
	dirEntries, err := ioutil.ReadDir("db/")
	if err != nil {
		t.Fatal(err)
	}

	indexBatchSize := 100
	batch := index.NewBatch()
	batchCount := 0

	for _, dirEntry := range dirEntries {
		filename := dirEntry.Name()
	
		jsonBytes, err := ioutil.ReadFile("db/" + filename)
			if err != nil {
				t.Fatal(err)
			}
	
		var jsonDoc interface{}
		err = json.Unmarshal(jsonBytes, &jsonDoc)
		if err != nil {
			t.Fatal(err)
		}
		ext := filepath.Ext(filename)
		docId := filename[:(len(filename) - len(ext))]
		batch.Index(docId, jsonDoc)
		batchCount++
	
		if batchCount >= indexBatchSize {
			err = index.Batch(batch)
			if err != nil {
				t.Fatal(err)
			}
			batch = index.NewBatch()
			batchCount = 0
		}
	}
	
	termQuery := bleve.NewTermQuery("Arch")
	termQuery.SetField("name")
	termSearchRequest := bleve.NewSearchRequest(termQuery)
	termSearchResult, err := index.Search(termSearchRequest)
	if err != nil {
		t.Error(err)
	}

}

type jsonFile struct {
	
	filename string
	contents []byte

}

