package internal

import (
	"encoding/json"
	_ "encoding/json"
	"fmt"
	_ "fmt"
)

// WikiPage describes simple wiki page
type WikiPage struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  []byte `json:"body"`
}

func (wp WikiPage) save() error {
	jsonData := []byte(`{"id": 1, "title": "First", "Body": "Lorum ipsen"}`)
	// content, err := ioutil.ReadFile("test.json")
	// if err != nil {
	// 	panic(err.Error())
	// }
	var test WikiPage
	err := json.Unmarshal(jsonData, &test)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", test)
	return nil
}

// func (wp *WikiPage) load(filename string) error {
// content, err := ioutil.ReadFile()
// err := json.Unmarshal(jsonData, &wp)
// if err != nil {
// 	fmt.Println("error:", err)
// }
// return nil
// }
