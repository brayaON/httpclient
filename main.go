package main

import ( 
    "github.com/brayaON/httpclient/methods"
    "github.com/brayaON/httpclient/auth"
)

const (
    baseURL = "http://localhost:80"
)

func callMethods() {
    methods.Get(baseURL)
    methods.Post(baseURL)
    methods.Delete(baseURL)
    methods.Put(baseURL)
    methods.Patch(baseURL)
}

func main() {
    auth.BasicAuth(baseURL)
}
