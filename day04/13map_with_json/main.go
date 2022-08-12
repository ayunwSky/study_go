package main

import (
	"encoding/json"
	"fmt"
)

// map 和 json 互相转换

func jsonToMap() {
	jsonStr := `
	{
		"name": "ayunwSky",
		"age": 24
	}
	`
	var mapResult map[string]interface{}

	// 用 json.Unmarshal(data []byte, v interface{})进行转换,返回 error 信息
	if err := json.Unmarshal([]byte(jsonStr), &mapResult); err != nil {
		fmt.Printf("json to map failed, err:%v\n", err)
		return
	}
	fmt.Println(&mapResult)

	if key, isOK := mapResult["age"]; isOK {
		fmt.Println(key)
	}

}

func mapToJson() {
	mapInstance := make(map[string]interface{})

	mapInstance["Name"] = "ayunwSky"
	mapInstance["Age"] = 24
	mapInstance["Province"] = "浙江"
	mapInstance["City"] = "杭州"

	jsonStr, err := json.Marshal(mapInstance)
	if err != nil {
		fmt.Printf("map to json failed, err:%v\n", err)
		return
	}
	fmt.Printf("json string is: %s\n", jsonStr)
}

func main() {
	jsonToMap()
	fmt.Println()
	mapToJson()
}
