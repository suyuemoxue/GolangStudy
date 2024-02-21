package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func testStructJson() {
	monster := Monster{
		Name:  "皮卡丘",
		Age:   1,
		Skill: "十万伏特",
	}
	//json序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println(err)
	}
	str := string(data)
	fmt.Println(str)
	//反序列化
	var monster01 Monster
	json.Unmarshal([]byte(str), &monster01)
	fmt.Println(monster01)
}

func testMapJson() {
	a := map[string]interface{}{}
	a["name"] = "喷火龙"
	a["age"] = 3
	a["skill"] = "喷射火焰"

	//json序列化
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}
	str := string(data)
	fmt.Println(str)
	//反序列化
	a1 := map[string]interface{}{}
	json.Unmarshal([]byte(str), &a1)
	fmt.Println(a1)
}

func testSliceJson() {
	var slice []map[string]interface{}

	a1 := map[string]interface{}{}
	a1["name"] = "杰尼龟"
	a1["age"] = 2
	a1["skill"] = "水炮"

	a2 := map[string]interface{}{}
	a2["name"] = "洛奇亚"
	a2["age"] = 200
	a2["skill"] = "神鸟猛击"

	a3 := map[string]interface{}{}
	a3["name"] = "裂空座"
	a3["age"] = 500
	a3["skill"] = "龙之怒"

	a4 := map[string]interface{}{}
	a4["name"] = "妙蛙种子"
	a4["age"] = 2
	a4["skill"] = "阳光烈焰"

	slice = append(slice, a1, a2, a3, a4)

	//json序列化
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Println(err)
	}
	str := string(data)
	fmt.Println(str)
	//反序列化
	var slice01 []map[string]interface{}
	json.Unmarshal([]byte(str), &slice01)
	fmt.Println(slice01)

}

func main() {
	testStructJson()
	testMapJson()
	testSliceJson()
}
