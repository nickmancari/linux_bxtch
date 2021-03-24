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
	
	docs, err := LoadDocuments("test_template.xml")
	if err != nil {
		fmt.Println(err)
	}

	elements := SearchData(docs, search)
	
	elementsMap := make(map[string]string)
	for i := 0; i < len(elements); i +=2 {
    		elementsMap[elements[i]] = elements[i+1]
	}

	data := struct {
		Res string
		Search map[string]string
	}{
		Search:	elementsMap, 
		Res:	search,
	}
	tpl.ExecuteTemplate(w, "search.html", data)
}
