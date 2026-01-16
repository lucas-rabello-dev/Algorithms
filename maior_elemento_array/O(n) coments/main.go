package main


func main() {
	var arr = []int{7, 4, 7, 3, 3, 2, 8, 8}
	println(foo(arr))
}

func foo(arr []int) int {
	// verifica se o arr eh vazio
	if len(arr) == 0 {
		return -1
	}
	// inicializa o map
	count := make(map[int]int)
	
	// contando a ocorrencia
	// preenche todo o map com a chave sendo os numeros do array e o valor sendo 1 para todos
	// se reptirem, ou seja, se o numero (num) ja estiver no map ele acessa ele e incrementa seu valor
	for _, num := range arr {
		count[num]++
	}
	
	// o maior valor ja encontrado fica aqui
	max := -1 // ele tem esse valor por que se tivesse 0 ou 1 no array poderia causar confusao
	// ajuda a identificar se ja encontrou pelo menos um numero
	found := false
	// exemplo se fosse com o numero 7 dps das ocorrencias
	// i = 7 | num = 2 (numero de ocorrencias)
	for i, num := range count {
		// se o numero de ocorrencia for igual a 1 (so cai os numeros que ocorreram uma vez ex: 4, 2) ele cai nesse if
		// se nao for um numero com uma ocorrencia ele nem entra no if
		if num == 1 {
			// cai na condicional se found == false OU o i = numero do array e chave do map for maior do que o max que guarda a maior ocorrencia
			if !found || i > max {
				// o max recebe o valor do i = (os numeros com uma ocorrencia)
				max = i
				// found fica true
				found = true
			}
		}
	}
	
	if !found {
		return -1
	}

	return max
}