package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

type Data struct {
	Metadata interface{} `json:"meta,omitempty"`
	Users []Employee `json:"data,omitemtpy"`
}

type Employee struct {
	Id     int    `json:"id, omitempty"`
	Name   string `json:"name, omitempty"`
	Email  string `json:"email, omitempty"`
	Gender string `json:"gender, omitempty"`
	Status string `json:"status, omitempty"`
}

func appendresponse(client *http.Client, req *http.Request, ch chan Employee, wg *sync.WaitGroup) {
	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("1: %s", err.Error())
		return
	}
	body, err := ioutil.ReadAll(response.Body)
	// fmt.Printf("%s",body)
	if err != nil {
		fmt.Printf("2: %s", err.Error())
		return
	}
	// fmt.Printf("%s",body)
	 var d  = new(Data)
	err = json.Unmarshal(body, &d)
	if err != nil {
		fmt.Printf("3: %s", err.Error())
		return
	}
	// fmt.Println(d)
	for _, v := range d.Users {
		ch <- v
	}
	close(ch)
	wg.Done()
}
func getNewRequest(method , url string) (*http.Request){
	req, err := http.NewRequest(method,url, nil)
	if err != nil {
		fmt.Errorf("Error while creating new Request")
		return nil
	}
	return req
}
func main() {
	client := &http.Client{}
	var ch chan Employee
	ch = make(chan Employee)
	req1, err := http.NewRequest("GET", "https://gorest.co.in/public/v1/users?page=1", nil)
	if err != nil {
		fmt.Errorf("Error while creating new Request")
	}

	req2, err := http.NewRequest("GET", "https://gorest.co.in/public/v1/users?page=2", nil)
	if err != nil {
		fmt.Errorf("Error while creating new Request")
	}

	req3, err := http.NewRequest("GET", "https://gorest.co.in/public/v1/users?page=3", nil)
	if err != nil {
		fmt.Errorf("Error while creating new Request")
	}
	var wg sync.WaitGroup
	req := []*http.Request{req1,req2,req3}
	for _, r := range req{
		wg.Add(1)
		go appendresponse(client, r, ch, &wg)	
		wg.Wait()
	}

	for v := range ch {
		fmt.Println("Name:",strings.ToLower(v.Name))
		fmt.Println("Email:",strings.ToLower(v.Email))
		fmt.Println("gender:",v.Gender)
		fmt.Println("Status:",strings.ToUpper(v.Status))
		fmt.Println()
	}

}

// https://gorest.co.in/public/v1/users?page=1
// https://gorest.co.in/public/v1/users?page=2
// https://gorest.co.in/public/v1/users?page=3

// "data": [
//     {
//       "id": 15,
//       "name": "Brajesh Shah",
//       "email": "brajesh_shah@kreiger-bosco.biz",
//       "gender": "female",
//       "status": "active"
//     },
