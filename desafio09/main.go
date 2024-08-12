/*
Tendo como dados de entrada a altura de uma pessoa,
construa um algoritmo que calcule seu peso ideal,
usando a seguinte fórmula: (72.7*altura) - 58
*/

/*
Lógica:
Este programa tem como objetivo calcular o peso ideal de uma pessoa com base na sua altura, utilizando a fórmula: `(72.7 * altura) - 58`.

1. Primeiro, declaramos a função `pesoIdeal` que contém a lógica necessária para realizar o cálculo do peso ideal.

2. Dentro da função, definimos uma variável:
   - `altura` do tipo `float64`, para armazenar a altura da pessoa inserida pelo usuário.

3. Usamos a função `fmt.Println` para exibir uma mensagem solicitando ao usuário que insira a sua altura.

4. Capturamos o valor inserido pelo usuário usando a função `fmt.Scan` e armazenamos na variável `altura`.

5. Para calcular o peso ideal, utilizamos a fórmula fornecida:
   - `pesoIdeal = (72.7 * altura) - 58`
   - Aqui, multiplicamos a altura por 72.7 e, em seguida, subtraímos 58 do resultado para obter o valor do peso ideal.

6. Em seguida, usamos `fmt.Println` para exibir o valor calculado de `pesoIdeal` na tela, mostrando assim o peso ideal correspondente à altura fornecida.

7. Finalmente, a função `pesoIdeal` é chamada dentro da função `main`, que é o ponto de entrada do programa, fazendo com que a execução comece e siga o fluxo descrito acima.

Este programa é um exemplo simples e eficaz de como realizar cálculos baseados em fórmulas matemáticas com entrada do usuário em Go.
*/

package main

import "fmt"

func pesoIdeal() {
	var altura float64
	var pesoIdeal float64
	fmt.Println("Digite a sua altura")
	fmt.Scan(&altura)

	pesoIdeal = (72.7 * altura) - 58
	fmt.Println("O peso ideal para a sua altura é: ", pesoIdeal)
}

func main() {
	pesoIdeal()
}
