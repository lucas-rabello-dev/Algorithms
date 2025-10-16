package main

import "fmt"

func main() {

	array := []int{}

	for i := 0; i < 101; i++ {
		array = append(array, i)
	}
	index, tentativas := BinarySearch(array, 1)

	fmt.Println(array[index], tentativas)

}

func BinarySearch(array []int, item int) (int, int) {
	menor := 0 // refente ao indice 0
	maior := len(array) -1 // refente ao indice ultimo valor do array
	tentativas := 0

	for menor <= maior {
		tentativas += 1
		meio := (menor + maior) / 2 // pega o meio
		chute := array[meio] // pega o valor do meio no array

		if chute == item {
			return meio, tentativas
		}
		if chute < item {
			menor = meio + 1 // um pra frente do meio
		} else {
			maior = meio - 1 // um pra trás do meio
		}
	}
	return 0, tentativas
}