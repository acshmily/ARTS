package main

import "fmt"

/**
  A -> 1
  B -> 2
  C -> 3
  ...
  Z -> 26
  AA -> 27
  AB -> 28
  ...

*/

func titleToNumber(s string) int {
	var r int
	for _, e := range s {

		r = r*26 + int(e-'A'+1)
	}
	return r

}

func main() {
	str := "AA"
	fmt.Println(titleToNumber(str))
}
