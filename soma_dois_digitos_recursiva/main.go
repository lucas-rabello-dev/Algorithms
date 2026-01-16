package main

func main() {
	println(foo(9875))
}

func foo(number uint) int {
	// verifica se o numero eh menor que 10 (sendo assim sera um digito e retorna)
	if number < 10 {
		return int(number)
	}
	
	var soma int
	for number > 0 {
		// pega o ultimo digito do numero e coloca ele no soma
  		soma += int(number) % 10
        number /= 10
	}
	
	// recursao da funcao
	return foo(uint(soma))
}