package methods

import (
    "fmt"
    "net/http"
    "log"
    "io"
    "encoding/json"
    "net/url"
    "bytes"
)

func Get(baseURL string) {
    resp, err := http.Get(baseURL+"/get")
    if err != nil {
	log.Fatal(err)
    }

    defer resp.Body.Close()
    if resp.StatusCode == http.StatusOK {
	fmt.Println("-------GET Method-------")
	fmt.Println("Protocol:", resp.Proto)
	fmt.Println("Status Code:", resp.Status)

	data1, _ := io.ReadAll(resp.Body)
	fmt.Println(string(data1))
    }
}

func Post(baseURL string) {
    mapData := url.Values {
	"name": {"brayan ortiz (bof)"}, 
	"role": {"software engineer"},
    }

    resp, err := http.PostForm(baseURL+"/post?param1=value1", mapData)
    if err != nil {
	log.Fatal(err)
    }
    defer resp.Body.Close()

    if resp.StatusCode == http.StatusOK {
	var respBody map[string] interface{}
	json.NewDecoder(resp.Body).Decode(&respBody)

	jsonRequest, _ := json.MarshalIndent(respBody["form"], "", " ")

	fmt.Println("-------POST Method-------")
	fmt.Println(string(jsonRequest))
    }
}

func Delete(baseURL string) {
    data := map[string] int {
	"productId": 3333,
    }

    jsonData, _ := json.Marshal(data)

    url := fmt.Sprintf("%s/delete?param1=value1&param2=value2", baseURL)

    req, err := http.NewRequest(http.MethodDelete, url, bytes.NewReader(jsonData))
    if err != nil {
	log.Fatal(err)
    }

    res, err := http.DefaultClient.Do(req)
    if err != nil {
	log.Fatal(err)
    }

    defer res.Body.Close()

    if res.StatusCode == http.StatusOK {
	var resBody map[string] interface{}

	json.NewDecoder(res.Body).Decode(&resBody)
	jsonRes, _ := json.MarshalIndent(resBody["json"], "", " ")

	fmt.Println("-------DELETE Method-------")
	fmt.Println(string(jsonRes))
    }
}

func Put(baseURL string) {
    data := map[string] string {
	"productId": "1234",
	"productName": "Computer",
	"productPrice": "200000",
    }
    jsonData, _ := json.Marshal(data) 

    url := fmt.Sprintf("%s/put?param1=value1", baseURL)

    req, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(jsonData))
    if err != nil {
	log.Fatal(err)
    }

    res, err := http.DefaultClient.Do(req)
    if err != nil {
	log.Fatal(err)
    }
    defer res.Body.Close()

    if res.StatusCode == http.StatusOK {
	var resBody map[string] interface{}
	json.NewDecoder(res.Body).Decode(&resBody)

	bodyJson, _ := json.MarshalIndent(resBody["json"], "", " ")
	fmt.Println("-------PUT Method-------")
	fmt.Println(string(bodyJson))
    }
}

func Patch(baseURL string) {
    data := map[string] string {
	"productId": "1234",
	"productPrice": "300000",
    }

    jsonData, _ := json.Marshal(data)
    url := fmt.Sprintf("%s/patch", baseURL)
    req, err := http.NewRequest(http.MethodPatch, url, bytes.NewReader(jsonData))
    if err != nil {
	log.Fatal(err)
    }

    res, err := http.DefaultClient.Do(req)
    if err != nil {
	log.Fatal(err)
    }
    defer res.Body.Close()

    if res.StatusCode == http.StatusOK {
	var rawBody map[string] interface{}
	json.NewDecoder(res.Body).Decode(&rawBody)

	jsonBody, _ := json.MarshalIndent(rawBody["json"], "" , " ")
	fmt.Println("-------PATCH Method-------")
	fmt.Println(string(jsonBody))
    }
}
