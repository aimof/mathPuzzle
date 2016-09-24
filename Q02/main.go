package main

import "fmt"

func main() {
	for a := 1; a <= 9; a++ {
		for b := 0; b <= 9; b++ {
			for c := 0; c <= 9; c++ {
				for d := 1; d <= 9; d++ {
					if cal1(a, b, c, d) || cal2(a, b, c, d) || cal3(a, b, c, d) || cal4(a, b, c, d) || cal5(a, b, c, d) || cal6(a, b, c, d) || cal7(a, b, c, d) {
						fmt.Println(1000*a + 100*b + 10*c + d)
					}
				}
			}
		}
	}
}

func correctNum(a, b, c, d int) int {
	return 1000*d + 100*c + 10*b + a
}

func cal1(a, b, c, d int) bool {
	return (100*a+10*b+c)*d == correctNum(a, b, c, d)
}
func cal2(a, b, c, d int) bool {
	return (10*a+b)*(10*c+d) == correctNum(a, b, c, d)
}
func cal3(a, b, c, d int) bool {
	return a*(100*b+10*c*d) == correctNum(a, b, c, d)
}
func cal4(a, b, c, d int) bool {
	return (10*a+b)*c*d == correctNum(a, b, c, d)
}
func cal5(a, b, c, d int) bool {
	return a*(10*b+c)*d == correctNum(a, b, c, d)
}
func cal6(a, b, c, d int) bool {
	return a*b*(10*c+d) == correctNum(a, b, c, d)
}
func cal7(a, b, c, d int) bool {
	return a*b*c*d == correctNum(a, b, c, d)
}
