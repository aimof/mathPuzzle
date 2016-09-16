package main

import "testing"

func TestParindrome(t *testing.T) {
	actual := palindrome("1234567890")
	expected := false
	if actual != expected {
		t.Errorf("非回文がtrueになる")
	}
}
