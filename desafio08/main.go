/*
Faça um Programa que peça a temperatura em graus Farenheit, transforme e mostre
a temperatura em graus Celsius.
C = (5 * (F-32) / 9).
*/

/*
Lógica:
Este programa tem como objetivo converter uma temperatura fornecida pelo usuário em graus Fahrenheit para graus Celsius.

1. Primeiro, declaramos a função `fahrenheitParaCelsius` que contém a lógica necessária para realizar a conversão de temperatura.

2. Dentro da função, definimos uma variável:
   - `fahrenheit` do tipo `float64`, para armazenar a temperatura em Fahrenheit inserida pelo usuário.

3. Usamos a função `fmt.Println` para exibir uma mensagem solicitando ao usuário que insira a temperatura em graus Fahrenheit.

4. Capturamos o valor inserido pelo usuário usando a função `fmt.Scan` e armazenamos na variável `fahrenheit`.

5. Para converter a temperatura de Fahrenheit para Celsius, utilizamos a fórmula:
   - `celsius = (5 * (fahrenheit - 32) / 9)`
   - Aqui, subtraímos 32 do valor em Fahrenheit, multiplicamos por 5, e dividimos o resultado por 9 para obter o valor em Celsius.

6. Em seguida, usamos `fmt.Println` para exibir o valor calculado de `celsius` na tela, mostrando assim a temperatura correspondente em graus Celsius.

7. Finalmente, a função `fahrenheitParaCelsius` é chamada dentro da função `main`, que é o ponto de entrada do programa, fazendo com que a execução comece e siga o fluxo descrito acima.

Este programa é um exemplo simples e eficaz de como realizar uma conversão de temperatura com entrada do usuário em Go.
*/

package main

import "fmt"

func farenheitParaCelsius() {
	var farenheit float64
	fmt.Println("Digite a temperatura em graus farenheit: ")
	fmt.Scan(&farenheit)

	celsius := (5 * (farenheit - 32) / 9)

	fmt.Println("Convertendo farenheit para celcius...")
	fmt.Println("A temperatura em graus celsius é: ", celsius)
}

func main() {
	farenheitParaCelsius()
}
