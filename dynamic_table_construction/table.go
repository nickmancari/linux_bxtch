package main

import (
    "html/template"
    "net/http"
)

type Data struct {
    Items []Devicevalue_view
}

type Devicevalue_view struct {
    Devicetype string
    Iddevice   string
    Devicename string
    Oidname    string
    Value      string
}

func page_1(w http.ResponseWriter, r *http.Request) {
    data := Data{}
    for i := 1; i < 10; i++ {
        view := Devicevalue_view{
            Devicetype: "devicetype",
            Iddevice:   "iddevice",
            Devicename: "devicename",
            Oidname:    "oidname",
            Value:      "value",
        }

        data.Items = append(data.Items, view)
    }

    tmpl, _ := template.ParseFiles("./index.html")
    tmpl.Execute(w, data)
}
func main() {
    http.HandleFunc("/table", page_1)
    http.ListenAndServe(":80", nil)
}
