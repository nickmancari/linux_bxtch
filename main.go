// POC example of 'search & serve' using data from JSON file

package main

import (
	"net/http"
	"fmt"
	"html/template"
	"github.com/thedevsaddam/gojsonq"	
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
	
	docs, err := loadDocuments("linux_database.xml")
	if err != nil {
		fmt.Println(err)
	}

	res := searchData(docs, search)

	data := map[string]string{"Search": search, "Link": link, "Image": image,}
	tpl.ExecuteTemplate(w, "search.html", data)
}
