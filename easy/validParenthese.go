package main

import "fmt"

func isValid(s string) bool {

	for i := 0; i < len(s)-1; i++ {
		if s[i] != s[i+1] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isValid("{}[]()"))
}
