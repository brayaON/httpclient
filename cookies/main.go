package cookies

import (
    "fmt"
    "net/http"
    "log"
    "io"
    "net/http/cookiejar"
)

func GetCookies(baseURL string) {
    url := fmt.Sprintf("%s/cookies", baseURL)

    req, err := http.NewRequest(http.MethodGet, url, nil)
    if err != nil {
	log.Fatal(err)
    }

    req.Header.Set("Cookie", "cookie1=value1;cookie2=value2;cookie3=value3")

    res, err := http.DefaultClient.Do(req)
    if err != nil {
	log.Fatal(err)
    }

    defer res.Body.Close()
    rawBody, _ := io.ReadAll(res.Body)
    fmt.Println(string(rawBody))
}

func SetCookie(baseURL string) {
    url := fmt.Sprintf("%s/cookies/set/user/brayan", baseURL)
    
    req, err := http.NewRequest(http.MethodGet, url, nil)
    if err != nil {
	log.Fatal(err)
    }

    ck1 := http.Cookie {
	Name: "cookie1",
	Value: "value1",
	Path: "/",
    }
    ck2 := http.Cookie {
	Name: "cookie2",
	Value: "value2",
	Path: "/",
    }
    req.AddCookie(&ck1)
    req.AddCookie(&ck2)

    jar, err := cookiejar.New(nil)
    if err != nil {
	log.Fatal(err)
    }

    client := http.Client{
	CheckRedirect: func(req *http.Request, via []*http.Request) error {
	    fmt.Printf("Redirected from %s to %s\n", req.Referer(), req.URL)
	    return nil
	},
	Jar: jar,
    }

    res, err := client.Do(req)
    if err != nil {
	log.Fatal(err)
    }

    defer res.Body.Close()

    rawBody, _ := io.ReadAll(res.Body)
    fmt.Println(string(rawBody))
}
