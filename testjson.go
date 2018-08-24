package main

import (
	"github.com/buger/jsonparser"
	"fmt"
)

func main() {
	data := []byte(`{
  "person": {
    "name": {
      "first": "Leonid",
      "last": "Bugaev",
      "fullName": "Leonid Bugaev"
    },
    "github": {
      "handle": "buger",
      "followers": 109
    },
    "avatars": [
      { "url": "https://avatars1.githubusercontent.com/u/14009?v=3&s=460", "type": "thumbnail" },
      { "url": "https://avatars1.githubusercontent.com/u/14111?v=3&s=460", "type": "thumbnail" },
      { "url": "https://avatars1.githubusercontent.com/u/14113?v=3&s=460", "type": "thumbnail" }
    ]
  },
  "company": {
    "name": "Acme"
  }
}`)

	// You can use `ArrayEach` helper to iterate items [item1, item2 .... itemN]
	//jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
	//	fmt.Println(jsonparser.Get(value, "url"))
	//}, "person", "avatars")

	// Or use can access fields by index!
	con,err := jsonparser.GetString(data, "person", "avatars", "[1]", "url")
	fmt.Println(string(con), err)
	//
	//// You can use `ObjectEach` helper to iterate objects { "key1":object1, "key2":object2, .... "keyN":objectN }
	//jsonparser.ObjectEach(data, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
	//	fmt.Printf("Key: '%s'\n Value: '%s'\n Type: %s\n", string(key), string(value), dataType)
	//	return nil
	//}, "person", "name")
}
