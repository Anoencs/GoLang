package main

import (
	"fmt"
	"regexp"
)

func sum(s1 *string, s2 *string) (string, error) {
	///////// regex
	var regexNum = regexp.MustCompile("[-+*/a-zA-Z]")
	//////////////  check case nil ////////////////
	if s1 == nil{
		s1 = new(string)
		*s1 = ""
	}

	if s2 == nil{
		s2 = new(string)
		*s2 = ""
	}
	//////////////init variable//////////////////////////////////////////////////////
	var res string
	var len1 int = len(*s1)
	var len2 int = len(*s2) 
	var num1 uint8
	var num2 uint8
	var maxlen int
	var carry uint8 = 0

	//////////////////// check case not correct formate ////////////////
	if regexNum.MatchString(*s1){
		return "",ExNumberFormationException{"ExNumberFormationException in string 1 at index:",regexNum.FindStringIndex(*s1)[0]}
	}
	if regexNum.MatchString(*s2){
		return "",ExNumberFormationException{"ExNumberFormationException in string 2 at index:",regexNum.FindStringIndex(*s2)[0]}
	}
	////////////////////////////////////////////////////////////////////////////
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
		carry = (num1 + num2 + carry ) / 10
	}
	
	if carry != 0{
		res = fmt.Sprint(carry,res)
	}


	return res,nil
}


