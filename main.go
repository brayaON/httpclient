package main

import ( 
    "github.com/brayaON/httpclient/methods"
)

const (
    baseURL = "http://localhost:80"
)

func main() {
    methods.Get(baseURL)
    methods.Post(baseURL)
    methods.Delete(baseURL)
    methods.Put(baseURL)
    methods.Patch(baseURL)
}
