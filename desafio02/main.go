/* Escrever um algoritmo para verificar se um numero é impar ou par
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
