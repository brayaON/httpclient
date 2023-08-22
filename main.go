package main

import ( 
    "github.com/brayaON/httpclient/methods"
    "github.com/brayaON/httpclient/auth"
    "github.com/brayaON/httpclient/cookies"
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

func callAuth() {
    auth.BasicAuth(baseURL)
    auth.BearerAuth(baseURL)
}

func main() {
    cookies.GetCookies(baseURL)
    cookies.SetCookie(baseURL)
}
