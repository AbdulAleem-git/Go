package main

import (
	"encoding/json"
	"fmt"
)

// type Student struct{
// 	name string
// 	age int
// }

// func main(){
// 	s1 := &Student{"abdul", 24}
// 	fmt.Println(s1)
// 	bytes_s1 , err := json.Marshal(s1)
// 	if err != nil {
// 		fmt.Println("errr: ",err)
// 	}
// 	fmt.Println(bytes_s1)
// 	fmt.Printf("%s",bytes_s1)
// }
type Product struct {
	Users              []User `json:"Users,omitempty"`
	Name               string `json:"Name,omitempty"`
	MonthOfManufacture int    `json:"MonthOfManufacture,omitempty"`
}

type User struct {
	Name   string `json:"name,omitempty"`
	Eid    int    `json:"eid,omitempty"`
	Age    int    `json:"age,omitempty"`
	Gender string `json:"gender,omitempty"`
}

func main() {
	e := &Product{
		Name:               "Mi",
		MonthOfManufacture: 12,
		Users: []User{
			{
				Name:   "abdul aleem ",
				Eid:    35,
				Age:    24,
				Gender: "male",
			},
			{
				Name:   "abdul aleem ",
				Eid:    35,
				Age:    24,
				Gender: "male",
			},
			{
				Name:   "abdul aleem ",
				Eid:    35,
				Age:    24,
				Gender: "male",
			},
			{
				Name:   "abdul aleem ",
				Eid:    35,
				Age:    24,
				Gender: "male",
			},
			{
				Name:   "abdul aleem ",
				Eid:    35,
				Age:    24,
				Gender: "male",
			},
			{
				Name:   "abdul aleem ",
				Eid:    35,
				Age:    24,
				Gender: "male",
			},
			{
				Name:   "abdul aleem ",
				Eid:    35,
				Age:    24,
				Gender: "male",
			},
			{
				Name:   "abdul aleem ",
				Eid:    35,
				Age:    24,
				Gender: "male",
			},
		},
	}

	a, err := json.Marshal(e)
	if err != nil {
		fmt.Println("err:", err)
	}
	var p = new(Product)
	err = json.Unmarshal(a, p)
	if err != nil {
		fmt.Println("err:", err)
	}
	var data struct {
		Key string `json:"TheKey"`
	}
	jsonData := []byte(`{"thekey": "Value"}`)
	err = json.Unmarshal(jsonData, &data)
	fmt.Printf("%v\n", data)
	fmt.Printf("%s\n", err)

}
