package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"
)

func main() {
    var card string
    fmt.Print("Simple Tool By Soud :)\n")
    fmt.Print("Enter Card: ")
    fmt.Scanln(&card)
    req, _ := http.Get("https://soud.me/api/CC?card=" + card)
    body, _ := ioutil.ReadAll(req.Body)
    if strings.Contains(string(body), "\"Status\":\"Dead\"") {
        fmt.Print("[-] Dead Card")
    } else if strings.Contains(string(body), "\"Status\":\"Unknown\"") {
        fmt.Print("[-] Unknown Card")
    } else if strings.Contains(string(body), "\"Status\":\"Live\"") {
        fmt.Print("[+] Live Card")
    }
    fmt.Scanln("\n")
}
