package main

import (
	"encoding/json"
	"fmt"
)

type Human struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}

// Marshal: 从 struct 转换至 string
func marshal2JsonString(h Human) string {
	updateBytes, err := json.Marshal(&h)
	if err != nil {
		println(err)
	}
	return string(updateBytes)
}

// Unmarshal: 从 string 转换至 struct
func unmarshal2Struct(humanStr string) Human {
	h := Human{}

	if err := json.Unmarshal([]byte(humanStr), &h); err != nil {
		println(err)
	}
	return h
}

func main() {
	h2s := Human{"Eric", 18, "male"}
	humanStr := marshal2JsonString(h2s)
	fmt.Println(humanStr)
	s2h := unmarshal2Struct(humanStr)
	fmt.Println(s2h)

	var obj interface{}
	err := json.Unmarshal([]byte(humanStr), &obj)
	objMap, ok := obj.(map[string]interface{})
	if !ok {
		println(err)
	}

	// Type switch on interface{}
	for k, v := range objMap {
		switch value := v.(type) {
		case string:
			fmt.Printf("type of %s is string, value is %v\n", k, value)
		case interface{}:
			fmt.Printf("type of %s is interface{}, value is %v\n", k, value)
		default:
			fmt.Printf("type of %s is wrong, value is %v\n", k, value)
		}
	}
}
