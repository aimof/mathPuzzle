package main

import "fmt"

func main() {
	cards := make([]bool, 101)
	for i := 1; i <= 100; i++ {
		cards = reverse(i, cards)
	}
	for i := 1; i <= 100; i++ {
		if cards[i] {
			fmt.Println(i)
		}
	}
}

func switchBool(a bool) bool {
	if a {
		return false
	}
	return true
}

func reverse(i int, cards []bool) []bool {
	for j := 1; j*i <= 100; j++ {
		cards[i*j] = switchBool(cards[i*j])
	}
	return cards
}
