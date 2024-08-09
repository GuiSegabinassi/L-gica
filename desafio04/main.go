/*
Faça um Programa que peça as 4 notas bimestrais e mostre a média.
*/

/*
Lógica:
Para resolver este desafio, precisamos calcular a média de quatro notas bimestrais fornecidas pelo usuário e exibir o resultado.

1. Primeiro, declaramos quatro variáveis (`nota1`, `nota2`, `nota3`, `nota4`), todas do tipo `float64`. O tipo `float64` foi escolhido para permitir a entrada de números com casas decimais, garantindo maior precisão no cálculo das médias.

2. Em seguida, utilizamos a função `fmt.Println` para exibir uma mensagem solicitando que o usuário digite as quatro notas.

3. Para capturar a entrada do usuário, utilizamos a função `fmt.Scan`, que lê os valores digitados e os armazena nas respectivas variáveis `nota1`, `nota2`, `nota3`, e `nota4`. A leitura é feita por referência, utilizando o operador `&`.

4. Depois de capturar os valores, somamos as quatro notas e dividimos o resultado por 4 para calcular a média. Essa operação é feita utilizando a expressão `(nota1 + nota2 + nota3 + nota4) / 4`, que garante que todas as notas sejam consideradas no cálculo da média.

5. Por fim, usamos a função `fmt.Println` para exibir a média calculada na tela, mostrando ao usuário o resultado final.

Este programa é um exemplo básico de como coletar múltiplas entradas, realizar cálculos com essas entradas e exibir o resultado em Go.
*/

package main

import "fmt"

func main() {
	var nota1, nota2, nota3, nota4 float64
	fmt.Println("Digite as notas")
	fmt.Scan(&nota1, &nota2, &nota3, &nota4)

	media := (nota1 + nota2 + nota3 + nota4) / 4
	fmt.Println("A média foi: ", media)
}
