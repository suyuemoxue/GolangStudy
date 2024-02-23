package store

import (
	"encoding/json"
	"fmt"
	"os"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

// Store 方法，将Monster对象序列化后，保存到monster.json文件中
func (m *Monster) Store() bool {
	//将一个monster对象序列化为json格式
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println(string(data))
	//将序列化后的结果写进monster.json文件中
	filePath := "E:/monster.json"
	err = os.WriteFile(filePath, data, 0666)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// ReStore 方法，从monster.json文件中读取数据，并反序列化为一个Monster对象
func (m *Monster) ReStore() bool {
	//读取monster.json文件
	filePath := "E:/monster.json"
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println("file = ", string(data))
	//反序列化为一个Monster对象
	err = json.Unmarshal(data, &m)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println(*m)
	return true
}
