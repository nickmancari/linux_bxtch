// POC example of 'search & serve' using data from XML file

package main

import (
	"net/http"
	"fmt"
	"html/template"
)


var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.html"))
} 

func main() {
	
	views := http.FileServer(http.Dir("views"))
	images := http.FileServer(http.Dir("images"))

	http.Handle("/views/", http.StripPrefix("/views/", views))
	http.Handle("/images/", http.StripPrefix("/images/", images))	
	http.HandleFunc("/", index)
	http.HandleFunc("/search", results)
	http.ListenAndServe(":80", nil)
}

func index( w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func results(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	
	search := r.FormValue("q")
	
	docs, err := LoadDocuments("linux_database.xml")
	if err != nil {
		fmt.Println(err)
	}

	res := SearchData(docs, search)
	
	elements := res
	
//	elementsMap := make(map[string]string)
//	for i, datap := range elements {
//		elementsMap[i] = datap
//	}

	elementsMap := make(map[string]string)
	for i := 0; i < len(elements); i +=2 {
    		elementsMap[elements[i]] = elements[i+1]
	}

	data := struct {
		Search map[string]string
	}{
		Search:	elementsMap, 
	}
	tpl.ExecuteTemplate(w, "search.html", data)
	fmt.Println(elementsMap)
}
