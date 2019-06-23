package main

import "fmt"

/**
  	1 -> A
   2 -> B
   3 -> C
   ...
   26 -> Z
   27 -> AA
   28 -> AB
   ...
*/

func convertToTitle(n int) string {
	var result = ""
	for n != 0 {

		if n%26 == 0 { // 判断是否整除
			result = "Z" + result
			n = n/26 - 1 //
		} else {
			// acsii处理
			result = string('A'+n%26-1) + result
			n = n / 26
		}
	}

	return string(result)

}

func main() {
	fmt.Println(convertToTitle(25))
}
