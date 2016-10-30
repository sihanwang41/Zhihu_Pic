package main

import "fmt"
import "net/http"
import "io/ioutil"
import "regexp"
import "strings"
//import "encoding/json"
//import "bytes"
import "os"
import "net/url"
var agent string = "Mozilla/5.0 (Windows NT 5.1; rv:33.0) Gecko/20100101 Firefox/33.0"

func get_xsrf() string{
    index_url := "http://www.zhihu.com"
    client := &http.Client{}
    req, err := http.NewRequest("GET", index_url, nil)
    if err != nil {
        fmt.Println(err)
        return "error"
    }
    req.Header.Set("User-Agent", agent)

    resp, err := client.Do(req)
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    pattern := `name="_xsrf" value="(.*?)"`
    reg := regexp.MustCompile(pattern)
    _xsrf := reg.FindString(string(body))
    return strings.Split(_xsrf, "=")[2]
}

func login() {
    var username string
    var password string
    fmt.Println("username: ")
    fmt.Scanf("%s", &username)
    fmt.Println("password: ")
    fmt.Scanf("%s", &password)

    client := &http.Client{}
    post_url := "http://www.zhihu.com/login/email"

    postValues := url.Values{}
    postValues.Add("remember_me", "true")
    postValues.Add("_xsrf", get_xsrf())
    postValues.Add("uname", username)
    postValues.Add("password", password)

    req, err := http.NewRequest("POST", post_url, strings.NewReader(postValues.Encode()))
    if err != nil {
        fmt.Println(err)
        return
    }
    req.Header.Set("User-Agent", agent)
    resp, err := client.Do(req)
    defer resp.Body.Close()
    if err != nil {
        fmt.Println(err)
        return
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(body)
    os.Stdout.Write(body)
    fmt.Println("hehe")
}


func main() {
    login()
}