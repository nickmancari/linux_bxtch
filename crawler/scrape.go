// find_html_comments_with_regex.go
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "regexp"
)

func main() {
    // Make HTTP request
    response, err := http.Get("https://myhsts.org/tutorial-list-of-all-linux-operating-system-distributions.php")
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()

    // Read response data in to memory
    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal("Error reading HTTP body. ", err)
    }

    // Create a regular expression to find comments
    re := regexp.MustCompile("<b>(.|\n)*?</b>")
    comments := re.FindAllString(string(body), -1)
    if comments == nil {
        fmt.Println("No matches.")
    } else {
        for _, comment := range comments {
            fmt.Println(comment)
        }
    }
}
