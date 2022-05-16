package main

import "fmt"

func sum(s1 string, s2 string) string {
	var res string

	var len1 int = len(s1)
	var len2 int = len(s2) 
	var num1 uint8
	var num2 uint8
	var maxlen int
	var carry uint8 = 0

	if len1 > len2{
		maxlen = len1
	}else{
		maxlen = len2
	}

	for i := 0; i < maxlen; i++{
		num1 = s1[len(s1) - i - 1] - '0'
		num2 = s2[len(s2) - i - 1] - '0'
		res = fmt.Sprint((d1 + d2 + carry ) % 10,res)
		carry = 0
		if d1 + d2 > 10{
			carry += 1
		}
	}
	if carry != 0{
		res = fmt.Sprint(carry,res)
	}
	return res
}

func main() {
	fmt.Println("Welcome to Sum program!")
	param1 := "929292929292"
	param2 := "212121212129"
	total := sum(param1, param2)
	fmt.Println("Result:", total)
}
