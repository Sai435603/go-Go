package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Id       string
	Prize    int      `json:"coursecost"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("about json")
	myCourses := []course{
		{"React", "1", 20, "123", []string{"web-dev", "js", "frontend"}},
		{"Node", "2", 30, "124", []string{"web-dev", "js", "backend"}},
		{"Express", "3", 40, "125", nil},
	}

	//Encoding json data

	fmt.Println(myCourses)
	encoadedJson, _ := json.Marshal(myCourses)
	fmt.Println(string(encoadedJson))
	// fmt.Println("\n\n\n")
	indentEncoadedJson, _ := json.MarshalIndent(myCourses, "", "\t")
	fmt.Printf("%s\n", indentEncoadedJson)

	//Decoding json data
	jsonDataFromWeb := []byte(`
	  {
		"coursename": "React",
		"Id": "1",
		"coursecost": 20,
		"tags": [
			"web-dev",
			"js",
			"frontend"
		]
	}
	`)
	fmt.Println("decoding json data from web")
	fmt.Println(string(jsonDataFromWeb))
	var myCourse course
	checkValidJson := json.Valid(jsonDataFromWeb)
	json.Unmarshal(jsonDataFromWeb, &myCourse)
	var myData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myData)
	if checkValidJson {
		fmt.Println("Yep the data is valid json")
		fmt.Println(myCourse)
		fmt.Println("map json")
		fmt.Println(myData)
		for key,val := range myData {
			 fmt.Printf("%v %v %T\n",key,val,val)
		}
	} else {
		fmt.Println("This is not a valid json")
	}
	// fmt.Printf("%T\n", checkValidJson)

}
