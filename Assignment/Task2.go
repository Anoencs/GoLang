package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	//"math"
)

var num_to19 = []string{"Không", "Một", "Hai", "Ba", "Bốn", "Năm", "Sáu", "Bảy", "Tám", "Chín", "Mười", "Mười một", "Mười hai", "Mười ba", "Mười bốn", "Mười lăm", "Mười sáu", "Mười bảy", "Mười tám", "Mười chín"}
var num_10s = []string{"Hai mươi", "Ba mươi", "Bốn mươi", "Năm mươi", "Sáu mươi", "Bảy mươi", "Tám mươi", "Chín mươi"}
var num_100s = []string{"trăm", "nghìn", "triệu", "tỉ"}

func convert_100(val int) string {
	if val < 20 {
		return num_to19[val]
	}

	var word string
	var word_val int
	for i := 0; i < len(num_10s); i++ {
		word = num_10s[i]
		word_val = 20 + 10*i
		if word_val+10 > val {
			if val%10 == 1 {
				return word + " mốt"
			} else if val%10 == 5 {
				return word + " lăm"
			} else if val%10 == 0 {
				return word
			} else {
				return word + " " + strings.ToLower(num_to19[val%10])
			}
		}
	}

	return ""
}

func convert_1000(val int) string {
	word_val := val / 100
	mod := val % 100
	var word string = ""
	if word_val > 0 {
		word = num_to19[word_val] + " " + num_100s[0]
		if mod > 0 {
			if mod <= 9 {
				word = word + " lẻ " + strings.ToLower(convert_100(mod))
			} else {
				word = word + " " + strings.ToLower(convert_100(mod))
			}
		}
	}
	if mod > 0 && word_val <= 0 {
		word = word + " " + convert_100(mod)
	}
	return word
}

// func powInt(x, y int) int {
//     return int(math.Pow(float64(x), float64(y)))
// }

// func convert(val int) string{
// 	if val < 100{
// 		return convert_100(val)
// 	}else if val < 1000{
// 		return convert_1000(val)
// 	}else{
// 		var word_val int
// 		var highOrderVal int
// 		var mod int
// 		var word string
// 		for i := 1; i < len(num_100s);i++{
// 			word_val = powInt(1000,i)
// 			if int(word_val) > val{
// 				highOrderVal = val/powInt(1000,i-1)
// 				mod = val - powInt(1000,i-1) * highOrderVal
// 				word = convert_1000(highOrderVal) + " " + num_100s[i-1]
// 				if mod > 0{
// 					word = word + convert(mod)
// 				}
// 			}

// 		}
// 		return word
// 	}
// 	return ""
// }

func main() {
	if len(os.Args) <= 1 {
		return
	}
	num, _ := strconv.Atoi(os.Args[1])
	fmt.Println(convert_1000(num))

}
