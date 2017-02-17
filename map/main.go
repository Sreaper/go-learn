package main

import (
	"fmt"
	"time"
)

var m map[string]map[string]string

func main() {
	m = make(map[string]map[string]string)
	//m["Bell Labs"] = make(map[string]string)
	//m["Bell Labs"]["szm"] = "lyl"
	fmt.Println(m["Bell Labs"]["szm"])
	for key,_:=range m {
		fmt.Println(key)
	}
	// 创建动态的map
	data:= make(map[string]interface{})
	data["1"] =22
	data["2"] ="2"
	for key, value := range data {
		fmt.Println(key)
		fmt.Println(value)
	}

	indexName := "metrics-"
	t := time.Now()
	indexName += t.Format("2006.01.02")
	fmt.Println(indexName)

	//if map
	info :=make(map[string]interface{})
	info["used_memory_rss"] =33
	if used_memory_rss, found := info["used_memory_rss"]; found {
		fmt.Println(found)
		fmt.Println(used_memory_rss)
	}
}

