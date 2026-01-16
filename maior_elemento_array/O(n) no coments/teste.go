package main

func main() {
	var arr = []int{7, 4, 7, 3, 3, 2, 8, 8}
	println(foo(arr))
}

func foo(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	
	values := make(map[int]int)

	for _, numbers := range arr {
		values[numbers]++
	}
	
	max := -1
	found := false 
	
	for numArray, ocorrencia := range values {
		if ocorrencia == 1 {
			if !found || numArray > max {
				max =  numArray				
				found = true
			}
		}
	}
	
	if !found {
		return -1
	}
	
	return max
}