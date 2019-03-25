package main

import (
	"fmt"
	"time"
)

/**
Tips: init string type map you should not to add ","  in init code
*/
func main() {
	map1 := map[string]string{"a": "b", "c": "d"}
	map2 := map[string]string{
		"a": "b",
		"c": "d", //you must add "," to the end or it will be mistake
	}
	// consider map1 cap is equal map2
	fmt.Println("map1 len is:", len(map1), "map2 len is:", len(map2))
	// to check key
	if _, ok := map1[""]; ok {
		fmt.Println(ok)
	}
	if _, ok := map2[""]; ok {
		fmt.Println(ok)
	}

	startTime := time.Now().Nanosecond()
	fmt.Println("for 1000s to string start")
	str := "123"
	for i := 0; i < 1000; i++ {
		str = str + str[1:1]
	}
	fmt.Println("for 1000s to string end,cost time:", time.Now().Nanosecond()-startTime)
	strList := []string{"1", "2", "3"}
	startTime = time.Now().Nanosecond()
	fmt.Println("for 1000s to array start")
	for i := 0; i < 1000; i++ {
		strList = append(strList, "1")
	}
	fmt.Println("for 1000s to array end,cost time:", time.Now().Nanosecond()-startTime)

}
