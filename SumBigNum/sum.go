package main

import (
	//"errors"
	"fmt"
	"regexp"
)

var regexNum = regexp.MustCompile("[0-9]*")


func sum(s1 *string, s2 *string) string {
	//////////////  check case nil ////////////////
	if s1 == nil{
		s1 = new(string)
		*s1 = ""
	}

	if s2 == nil{
		s2 = new(string)
		*s2 = ""
	}
	//////////////init variable/////////////////
	var res string
	var len1 int = len(*s1)
	var len2 int = len(*s2) 
	var num1 uint8
	var num2 uint8
	var maxlen int
	var carry uint8 = 0

	//////////////////// check case not correct formate
	if !regexNum.MatchString(*s1){
		fmt.Println("s1 not correct formate")
	}
	if !regexNum.MatchString(*s2){
		fmt.Println("s2 not correct formate")
	}
	//////////////////////////////////////////
	if len1 > len2{
		maxlen = len1
	}else{
		maxlen = len2
	}

	for i := 0; i < maxlen; i++{

		if len(*s1) - i - 1 >= 0{
			s1_val := *s1
			num1 = s1_val[(len(*s1) - i - 1)] - '0'
		}else{
			num1 = 0
		}

		if len(*s2) - i - 1 >= 0{
			s2_val := *s2
			num2 = s2_val[(len(*s2) - i - 1)] - '0'
		}else{
			num2 = 0
		}

		res = fmt.Sprint((num1 + num2 + carry ) % 10,res)
		carry = 0
		if num1 + num2 > 10{
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
	param1 := "1a23"
	param2 := "123"
	total := sum(&param1, &param2)
	fmt.Println("Result:", total)
}
