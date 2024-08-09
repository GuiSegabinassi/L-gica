/*
Faça um Programa que converta metros para centímetros.
*/

/*
Lógica:
Para resolver este desafio, precisamos converter uma distância fornecida pelo usuário em metros para centímetros e exibir o resultado.

1. Primeiro, declaramos a função `metroEmCm` para encapsular a lógica da conversão. Dentro dessa função, declaramos uma variável chamada `metros` do tipo `float64`. O tipo `float64` permite que o usuário insira valores com casas decimais, o que é útil para garantir precisão na conversão de medidas.

2. Em seguida, utilizamos a função `fmt.Println` para exibir uma mensagem na tela, solicitando que o usuário digite uma distância em metros. Essa mensagem orienta o usuário sobre o que ele deve inserir.

3. Para capturar a entrada do usuário, usamos a função `fmt.Scan`, que lê o valor digitado pelo usuário e o armazena na variável `metros`. O uso do operador `&` (endereço de memória) garante que o valor lido seja armazenado diretamente na variável.

4. Depois de capturar o valor em metros, realizamos a conversão multiplicando o valor de `metros` por 100, já que 1 metro é equivalente a 100 centímetros. O resultado dessa operação é armazenado na variável `centimetros`, que também é do tipo `float64` para manter a precisão.

5. Por fim, utilizamos `fmt.Println` para exibir a distância convertida em centímetros. A mensagem apresentada ao usuário inclui o valor final calculado.

6. No `func main()`, chamamos a função `metroEmCm` para executar a lógica de conversão quando o programa é executado.

Este programa é um exemplo básico e direto de como realizar uma conversão de unidades, desde a captura de entrada do usuário até o processamento e exibição do resultado em Go.
*/

package main

import "fmt"

func main() {
	metroEmCm()
}

func metroEmCm() {
	var metros float64
	fmt.Println("Digite a distancia em metros para converter em centimetros")
	fmt.Scan(&metros)

	centimetros := metros * 100
	fmt.Println("A distancia em centimetros é: ", centimetros)
}
