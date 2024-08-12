/*
Faça um Programa que pergunte quanto você ganha por hora e o número de horas
trabalhadas no mês. Calcule e mostre o total do seu salário no referido mês.
*/

/*
Lógica:
Este programa tem como objetivo calcular o salário mensal de uma pessoa com base no valor que ela ganha por hora e na quantidade de horas trabalhadas no mês.

1. Primeiro, declaramos a função `salarioMes` que contém a lógica necessária para realizar o cálculo do salário mensal.

2. Dentro da função, definimos duas variáveis:
   - `horasTrabalhadasNoMes` do tipo `int`, para armazenar o número de horas que a pessoa trabalhou no mês.
   - `valorDoSalarioPorHora` do tipo `float64`, para armazenar o valor que a pessoa ganha por hora.

3. Usamos a função `fmt.Println` para exibir uma mensagem solicitando ao usuário que insira o valor que ganha por hora e o número de horas trabalhadas no mês.

4. Capturamos os valores inseridos pelo usuário usando a função `fmt.Scan` e armazenamos nas variáveis correspondentes.

5. Para calcular o salário mensal, multiplicamos o valor que a pessoa ganha por hora (`valorDoSalarioPorHora`) pelo número de horas trabalhadas no mês (`horasTrabalhadasNoMes`). Este cálculo é realizado e o resultado é armazenado na variável `salarioMensal`.

6. Em seguida, usamos `fmt.Println` para exibir o valor calculado de `salarioMensal` na tela, mostrando assim o salário mensal do usuário.

7. Finalmente, a função `salarioMes` é chamada dentro da função `main`, que é o ponto de entrada do programa, fazendo com que a execução comece e siga o fluxo descrito acima.

Este programa é um exemplo simples e prático de como realizar cálculos básicos com entrada do usuário em Go.
*/

package main

import "fmt"

func salarioMes() {
	var horasTrabalhadasNoMes int
	var valorDoSalarioPorHora float64

	fmt.Println("Digite quanto você ganha por hora: ")
	fmt.Scan(&horasTrabalhadasNoMes)

	fmt.Println("Digite quantas horas você trabalhou no mês: ")
	fmt.Scan(&valorDoSalarioPorHora)

	salarioMensal := (valorDoSalarioPorHora * float64(horasTrabalhadasNoMes))

	fmt.Println("O seu salario mensal é: ", salarioMensal)
}

func main() {
	salarioMes()
}
