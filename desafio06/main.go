/*
Faça um Programa que peça o raio de um círculo, calcule e mostre sua área.
*/

/*
Lógica:
Este programa tem como objetivo calcular a área de um círculo com base no raio fornecido pelo usuário.

1. Primeiro, importamos os pacotes `fmt` e `math`:
   - `fmt`: Usado para exibir mensagens e capturar a entrada do usuário.
   - `math`: Importado para usar a constante `math.Pi`, que representa o valor de \(\pi\) com alta precisão.

2. Declaramos a função `calcularArea` que contém a lógica para calcular a área do círculo:
   - Dentro dessa função, declaramos uma variável `raio` do tipo `float64`, que armazenará o valor do raio fornecido pelo usuário.
   - Usamos `fmt.Print` para exibir uma mensagem pedindo ao usuário para inserir o valor do raio.
   - A função `fmt.Scanln(&raio)` captura o valor digitado pelo usuário e o armazena na variável `raio`.

3. Em seguida, calculamos a área do círculo usando a fórmula:
   \[
   \text{Área} = \pi \times \text{raio}^2
   \]
   - Isso é feito multiplicando `math.Pi` pelo valor do raio ao quadrado (`raio * raio`), e o resultado é armazenado na variável `area`.

4. Por fim, utilizamos `fmt.Println` para exibir o valor calculado da área na tela, mostrando a mensagem "A área do círculo é: " seguida do valor da área.

5. A função `calcularArea` é chamada dentro da função `main`, que é o ponto de entrada do programa, fazendo com que a operação de cálculo seja realizada quando o programa é executado.
*/

package main

import (
	"fmt"
	"math"
)

func calcularArea() {
	var raio float64
	fmt.Print("Digite o raio do círculo: ")
	fmt.Scanln(&raio)

	area := math.Pi * raio * raio

	fmt.Println("A Area do circulo é: ", area)

}

func main() {
	calcularArea()
}
