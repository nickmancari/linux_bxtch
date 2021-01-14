package main

import (
	"net/http"
	"html/template"
	"fmt"
	"github.com/thedevsaddam/gojsonq"
)


var tpl *template.Template

func init() {
        tpl = template.Must(template.ParseGlob("*.html"))
}

func main() {

        http.HandleFunc("/test", results)
        http.ListenAndServe(":80", nil)
}

func results(w http.ResponseWriter, r *http.Request) {
		type item struct {
			Name string `json:"name"`
			Download string `json:"download"`
		}
		v:= item{}
		
		gojsonq.New().File("../data/test.json").From("system.linux.[3]").Out(&v)
		
		data := fmt.Sprintf(v.Download)
		tpl.ExecuteTemplate(w, "table.html", data)
	
}
