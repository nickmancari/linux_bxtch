// POS example of 'search & serve' using data index

package main

import (
	"net/http"
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

	//query into Bleve and then result output to var that gets written to 	
	//search.html file

	data := struct {
		Search string
	}{
		Search: search,
	}
	tpl.ExecuteTemplate(w, "search.html", data)
}
