package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// func main() {
// 	file, err := os.Open("data.json")
// 	if err != nil {
// 		log.Fatalf("Error opening file: %v", err)
// 	}
// 	defer file.Close()

// 	var raw json.RawMessage
// 	err = json.NewDecoder(file).Decode(&raw)
// 	if err != nil {
// 		log.Fatalf("Error decoding JSON: %v", err)
// 	}

// 	fmt.Printf("Raw JSON: %s\n", raw)

// 	var data map[string]interface{}
// 	err = json.Unmarshal(raw, &data)
// 	if err != nil {
// 		log.Fatalf("Error unmarshaling JSON: %v", err)
// 	}

// 	value, exists := data["number"]
// 	if exists {
// 		switch v := value.(type) {
// 		case map[string]interface{}:
// 			for k := range v {
// 				fmt.Println("Key:", k)
// 			}
// 		case []interface{}:
// 			for i, element := range v {
// 				fmt.Printf("Element %d: %+v\n", i, element)
// 			}
// 		default:
// 			fmt.Printf("Value is of type %T\n", value)
// 		}
// 	} else {
// 		fmt.Println("Key 'number' does not exist in the map")
// 	}
// }

func enumerateValues(value interface{}) {
	switch v := value.(type) {
	case map[string]interface{}:
		for k, subValue := range v {
			fmt.Printf("Key: %s\n", k)
			enumerateValues(subValue)
		}
	case []interface{}:
		for i, subValue := range v {
			fmt.Printf("Element %d:\n", i)
			enumerateValues(subValue)
		}
	default:
		fmt.Printf("Value: %+v\n", v)
	}
}

func main() {
	file, err := os.Open("data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var data map[string]interface{}
	err = decoder.Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	enumerateValues(data)
}
