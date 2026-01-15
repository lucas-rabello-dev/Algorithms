package main

import (
	"fmt"
)

func main() {
	v := Atoi(" +34")
	fmt.Printf("%T | (%d)\n", v, v)
}

func Atoi(s string) int {
	i := 0 // indice para percorrer 
	sign := 1 // controla o numero pra ver se ele eh negativo ou positivo
	result := 0 // resultado final para o return
	
	// ignoramos os espacos e verificamos onde realmente comeca o numero
	// logo se o input for ( -32)
	// o segundo loop o i vai apontar para o '-' e assim encerrando o loop, porem o indice (i) tera o valor 2 onde realmente comeca o numero
	for i < len(s) && s[i] == ' ' {
		i++
	}
	
	// percorre a string e ve se existem o sinal de + e - na string 
	// se tiver o de menos invertemos o sinal de sign para que possamos nos orientar mais tarde
	// se for + ele nao faz nada
	for i < len(s) && (s[i] == '+' || s[i] == '-') {
		// inverte o sinal de sign
		if s[i] == '-' {
			sign = -1
		}
		// incrementa o indice (i) para que agora aponte para o inicio do numero sem o sinal
		i++
	}
	
	// convertendo os digitos
	// no for 
	// verificamos se os digitos (que iteramos pelo for each) sejam numeros de 0 a 9
	
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		// int(s[i] - '0') -> isso converte os bytes para numeros atraves do calculo da ascii table (igual quando converte no assembly kkkk)
		// exemplo: '7' - '0' = 7
		// o result*10 desloca os digitos para a esquerda
		// exemplo de 123
		/*
	
		i | char | cálculo   | result
		0 | '1'	 | 0*10 + 1	 | 1
		1 | '2'	 | 1*10 + 2	 | 12
		2 | '3'	 | 12*10 + 3 | 123
	     */
		
		result = result*10 + int(s[i] - '0')
		i++
	}
	// se o sign for negativo ele converter positivo para negativo
	return sign * result
}