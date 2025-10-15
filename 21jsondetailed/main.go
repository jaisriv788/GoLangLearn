package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to the JSON Detailed.")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJs Bootcamo", 299, "LearnCodeOnline.in", "abc123", nil},
		{"MERN Bootcamo", 199, "LearnCodeOnline.in", "abc1234", []string{"fullstack", "js/ts"}},
		{"GoLang Bootcamo", 599, "LearnCodeOnline.in", "abc12345", []string{"Backend", "Go"}},
	}

	//package this data as json data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
	// fmt.Println(finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	 {
                "coursename": "GoLang Bootcamo",
                "price": 599,
                "website": "LearnCodeOnline.in",
                "tags": [
                        "Backend",
                        "Go"
                ]
        }
	`)

	var lcoCourse course

	check := json.Valid(jsonDataFromWeb)

	if check {
		fmt.Println("Json is Valid.")

		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("Json Was Not Valid.")
	}

	//Case 2 when we want key value pairs

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("The key is %v and the value is %v and type is %T\n", k, v, v)
	}
}
