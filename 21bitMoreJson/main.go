package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type person struct {
	Name       string   `json:"name"`
	Age        int      `json:"age"`
	Gender     string   `json:"gender"`
	Occupation string   `json:"occupation"`
	Hobbies    []string `json:"hobbies,omitempty"`
	Tags       []string `json:"-"`
}

func encodeJson() {
	persons := []person{
		{
			Name:       "Pratik Dev",
			Age:        22,
			Gender:     "Male",
			Occupation: "Web Developer",
			Hobbies:    []string{"reading", "writing", "debating"},
			Tags:       []string{"tagOne", "tagTwo"},
		},
		{
			Name:       "Tasnuva Chowdhury",
			Age:        22,
			Gender:     "Female",
			Occupation: "Student",
			Hobbies:    []string{"reading", "dancing"},
			Tags:       []string{"tagOne", "tagTwo"},
		},
		{
			Name:       "Third Person",
			Age:        25,
			Gender:     "Male",
			Occupation: "Student",
			Hobbies:    nil,
			Tags:       []string{"tagOne", "tagTwo"},
		},
	}

	finalJson, err := json.MarshalIndent(persons, "", "\t")
	checkErrNil(err)

	responseBuilder := strings.Builder{}
	responseBuilder.Write(finalJson)

	fmt.Println(responseBuilder.String())
}

func decodeJson() {
	jsonByte := []byte(`{"name":"Pratik Dev","age":22,"gender":"Male","occupation":"Web Developer","hobbies":["reading","writing","debating"]}`)

	if !(json.Valid(jsonByte)) {
		checkErrNil(fmt.Errorf("invalid json"))
	}

	var data person
	json.Unmarshal(jsonByte, &data)
	fmt.Printf("%+v\n", data)

	var dataWithoutStruct map[string]interface{}
	json.Unmarshal(jsonByte, &dataWithoutStruct)
	fmt.Printf("%+v\n", dataWithoutStruct)
}

func main() {
	encodeJson()
	decodeJson()
}

func checkErrNil(err error) {
	if err != nil {
		panic(err)
	}
}
