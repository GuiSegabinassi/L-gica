// Escreva um programa que imprima na tela o valor digitado por um usuario

/*
Logica: Para esse Desafio, Devemos primeiro declarar uma variavel, essa variavel pode conter qualquer nome, mas para ser algo relacionado ao desafios irei chama-la de "numero",
ou seja, usaremos a sintaxe (var numero int) para declarar a nossa variavel. O "INT" significa que essa variavel chamada numero pode apenas receber valores inteiros, sem ponto flutuante.
Em seguida usaremos (Println) para imprimir na tela para que o usuario digite um valor e então iremos ler o valor digitado pelo usuário usando a função (Scanf)

E para finalizar, novamente usaremos (Println) para imprimir uma mensagem na tela, mas dessa vez fora das aspas, iremos passar a variavel "numero" que foi a variavel que guardou o valor digitado pelo usuário
*/
package main

import "fmt"

func main() {
	var numero int
	fmt.Println("Digite um número")
	fmt.Scan(&numero)

	fmt.Println("O numero digitado pelo usuário foi: ", numero)
}
