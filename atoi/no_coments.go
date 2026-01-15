package main

import (
	"fmt"
)

func main() {
	v := Atoi(" +34")
	fmt.Printf("%T | (%d)\n", v, v)
}

func Atoi(s string) int {
	i := 0 
	sign := 1
	result := 0
	
	for i < len(s) && s[i] == ' ' {
		i++
	}
	
	for i < len(s) && (s[i] == '+' || s[i] == '-') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}
	
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		result = result*10 + int(s[i] - '0')
		i++
	}
	return sign * result
}