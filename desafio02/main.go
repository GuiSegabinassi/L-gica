/* Escrever um algoritmo para verificar se um numero é impar ou par
 */

/*
Lógica:
Este programa tem como objetivo verificar se um número fornecido pelo usuário é par ou ímpar, e exibir uma mensagem correspondente.

1. Primeiro, declaramos a função `parOuImpar` que contém toda a lógica necessária para realizar a verificação. Dentro dessa função, declaramos uma variável chamada `numero` do tipo `int`. O tipo `int` é adequado aqui, pois estamos lidando com números inteiros.

2. Em seguida, utilizamos a função `fmt.Print` para exibir uma mensagem na tela, solicitando que o usuário insira um número. A função `Print` é usada para manter o cursor na mesma linha após a mensagem, permitindo que o usuário insira o número na mesma linha da mensagem.

3. Para capturar a entrada do usuário, usamos a função `fmt.Scanln`, que lê o valor digitado pelo usuário e o armazena na variável `numero`. O operador `&` (endereço de memória) é usado para garantir que o valor lido seja armazenado diretamente na variável.

4. Após capturar o número, usamos uma estrutura condicional `if-else` para determinar se o número é par ou ímpar:
   - A operação de módulo (`%`) é usada para verificar o resto da divisão de `numero` por 2.
   - Se `numero % 2 == 0`, isso significa que o número é divisível por 2 sem deixar resto, logo, o número é par, e o programa exibe a mensagem "O numero é par".
   - Caso contrário, ou seja, se o resto da divisão for diferente de 0, o número é ímpar, e o programa exibe a mensagem "O numero é ímpar".

5. A função `parOuImpar` é chamada dentro da função `main`, que é o ponto de entrada do programa, fazendo com que a verificação seja realizada quando o programa é executado.

Este programa é um exemplo simples e eficaz de como usar condicionais e operações matemáticas básicas para verificar a paridade de um número, fornecendo um bom exercício para reforçar conceitos fundamentais em Go.
*/

package main

import "fmt"

func parOuImpar() {
	var numero int
	fmt.Print("Digite um número: ")
	fmt.Scanln(&numero)

	if numero%2 == 0 {
		fmt.Println("O numero é par")
	} else {
		fmt.Println("O numero é impar")
	}
}
func main() {
	parOuImpar()
}
