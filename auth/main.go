package auth

import (
    "fmt"
    "net/http"
    "log"
    "io"
)

type user struct {
    Name string
    Password string 
}

func BasicAuth(baseURL string) {
    me := user {Name: "test", Password: "12345"}

    url := fmt.Sprintf("%s/basic-auth/%s/%s", baseURL, me.Name, me.Password)

    req, err := http.NewRequest(http.MethodGet, url, nil)
    if err != nil {
	log.Fatal(err)
    }
    req.SetBasicAuth(me.Name, me.Password)

    res, err := http.DefaultClient.Do(req)
    if err != nil {
	log.Fatal(err)
    }
    defer res.Body.Close()

    if res.StatusCode == http.StatusOK {
	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))
    } else {
	fmt.Println("Username or password are incorrect")
    }
}

func BearerAuth(baseURL string) {
    url := fmt.Sprintf("%s/bearer", baseURL)
    bearer := "Bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.D2BDBY1e4Ln7DwjnoCvdTDYywx7xbNh7jvPmfHf6mOM"

    req, err := http.NewRequest(http.MethodGet, url, nil)
    if err != nil {
	log.Fatal(err)
    }
    req.Header.Add("Authorization", bearer)

    res, err := http.DefaultClient.Do(req)
    if err != nil {
	log.Fatal(err)
    }
    defer res.Body.Close()

    if res.StatusCode == http.StatusOK {
	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))
    } else {
	fmt.Println("Username or password are incorrect")
    }
}
