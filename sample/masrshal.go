package main

import "fmt"
import "encoding/json"

var dataz = `{"foo": ["bar", "baz"], "boff": {"foo": "bar", "baz": "boff"}}`

type Dataz struct {
	Foo  []string `json:"foo"`
	Boff struct {
		Foo string `json:"foo"`
		Baz string `json:"baz"`
	} `json:"boff"`
}

func main() {
	// Method 1
	var d interface{}
	json.Unmarshal([]byte(dataz), &d)
	fmt.Println(d.(map[string]interface{})["foo"].([]interface{})[0])

	// Method 2
	var D Dataz
	json.Unmarshal([]byte(dataz), &D)
	fmt.Println(D.Foo[0])
}
