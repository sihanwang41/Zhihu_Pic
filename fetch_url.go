package main

import "fmt"
//import "io/ioutil"
import "net/http"
import "golang.org/x/net/html"
import "os"
import "strings"

func get_href(t html.Token) (ok bool, href string){
    for _, a := range t.Attr {
        if a.Key == "href" {
            href = a.Val
            ok = true
        }
    }
    return
}

func main() {
    url := "https://godoc.org/golang.org/x/net/html"
    resp, err := http.Get(url)
    if err != nil {
        fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
        os.Exit(1)
    }
    
    z := html.NewTokenizer(resp.Body)

    for {
        tt := z.Next()
        switch {
        case tt == html.ErrorToken:
            return
        case tt == html.StartTagToken:
            t := z.Token()
            isAnchor := t.Data == "a"
            if isAnchor {
                ok, href := get_href(t)
                if ok {
                    isValid_hrep := strings.Index(href, "http") == 0
                    if isValid_hrep {
                        fmt.Println(href)
                    }
                }
            }
        }
    }
    resp.Body.Close()
}