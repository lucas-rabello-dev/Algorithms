package main

import "fmt"


func main() {
	var arr = []int{7, 4, 7, 3, 3, 2, 8, 8}
	println(foo(arr))
}

func foo(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
		
	// criando o map
	values := make(map[int]int)
	for _, num := range arr {
		values[num] = 0
	}
	
	for _, num := range arr {
		for index := range values {
			if num == index {
				values[index]++
			}
		}
	}
	fmt.Println(values)
	
	var newArr []int
	for i, value := range values {
		if value == 1 {
			newArr = append(newArr, i)
		}
	}
	fmt.Println(newArr)
	
	if len(newArr) == 0 {
		return -1
	}
	
	return newArr[len(newArr)-1]
}