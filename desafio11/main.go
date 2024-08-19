/*Descrição: Escreva um programa que leia um valor de tempo em minutos e converta para horas e minutos.
 */

package main

import "fmt"

func conversaoDeTempo() {
	var tempoEmMinutos int
	fmt.Println("Digite o tempo em minutos: ")
	fmt.Scan(&tempoEmMinutos)

	horas := tempoEmMinutos / 60
	minutos := tempoEmMinutos % 60

	fmt.Printf("O tempo convertido é de: %d horas e %d minutos\n", horas, minutos)
}

func main() {
	conversaoDeTempo()
}
