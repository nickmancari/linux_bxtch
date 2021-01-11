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

	http.HandleFunc("/test", results)
        http.ListenAndServe(":80", nil)
}

func results(w http.ResponseWriter, r *http.Request) {
	
	type item struct {
                        	Name string `json:"name"`
                        	Download string `json:"download"`
               }
                v:= item{}

              gojsonq.New().File("./data/test.json").From("system.linux.[3]").Out(&v)
                

		out := fmt.Sprintf("Name: %s", "Download: %s", v.Name, v.Download)	

	tpl.ExecuteTemplate(w, "results.html", out)
}
