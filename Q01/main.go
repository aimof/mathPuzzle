package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	finish := false
	for i := 10; !finish; i++ {
		str2, str8, str10 := getStr(i)
		finish = check(str2, str8, str10)
		if finish {
			fmt.Printf("%d\n", i)
		}
	}

}

func getStr(num int) (str2 string, str8 string, str10 string) {
	str10 = strconv.Itoa(num)
	str8 = fmt.Sprintf("%o", num)
	str2 = fmt.Sprintf("%b", num)
	return
}

func check(str2 string, str8 string, str10 string) bool {
	if palindrome(str2) == true && palindrome(str8) == true && palindrome(str10) == true {
		return true
	}
	return false
}

func palindrome(str string) bool {
	splitedStrSlice := strings.Split(str, "")
	lenStrInt := len(splitedStrSlice)
	palindromeBool := true
	for i := 0; i < lenStrInt/2; i++ {
		if splitedStrSlice[i] == splitedStrSlice[lenStrInt-i-1] {
			continue
		} else {
			palindromeBool = false
			break
		}
	}
	return palindromeBool
}
