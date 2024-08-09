/*
Faça um Programa que peça dois números e imprima a soma.
*/

/*
Lógica:
Para resolver esse desafio, precisamos somar dois números fornecidos pelo usuário e exibir o resultado.

1. Primeiro, declaramos duas variáveis, `numero1` e `numero2`, ambas do tipo `float64`. O tipo `float64` foi escolhido para que possamos trabalhar com números que possuam casas decimais, oferecendo maior precisão nos cálculos.

2. Em seguida, utilizamos a função `fmt.Println` para exibir uma mensagem na tela pedindo ao usuário para digitar os números que deseja somar.

3. Para capturar a entrada do usuário, utilizamos a função `fmt.Scan`, que lê os valores digitados pelo usuário e os armazena nas variáveis `numero1` e `numero2`. Aqui, os valores lidos são passados por referência utilizando o operador `&`.

4. Depois de capturar os valores, somamos `numero1` e `numero2` e armazenamos o resultado na variável `resultado`. A operação de soma é simples e direta, utilizando o operador `+`.

5. Por fim, usamos a função `fmt.Println` novamente para exibir o resultado da soma na tela, mostrando ao usuário o valor final calculado.

*/

package main

import "fmt"

func main() {
	var numero1 float64
	var numero2 float64

	fmt.Println("Digite os numero a serem somados: ")
	fmt.Scan(&numero1, &numero2)

	resultado := numero1 + numero2

	fmt.Println("O Resultado da soma é: ", resultado)
}
