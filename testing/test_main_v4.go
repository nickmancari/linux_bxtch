// POC example of 'search & serve' using data index

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

        res := gojsonq.New().File("./data/test.json").From("system.linux").WhereContains("name", search).Only("download")
        link := fmt.Sprint(res)
	
	imagesearch := gojsonq.New().File("./data/icons.json").Find(search)
	image := fmt.Sprint(imagesearch)	

	data := struct {
		Search string
		Link string
		Image string
	}{
		Search: search,
		Link: link,
		Image: image,
	}
	tpl.ExecuteTemplate(w, "search.html", data)
}
