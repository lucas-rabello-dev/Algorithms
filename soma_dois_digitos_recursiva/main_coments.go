package main

func main() {
	println(foo(9875))
}

func foo(number uint) int {
	if number < 10 {
		return int(number)
	}
	
	var soma int
	for number > 0 {
  		soma += int(number) % 10
        number /= 10
	}
	
	return foo(uint(soma))
}