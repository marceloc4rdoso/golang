package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var numeros []int
	var soma, maior, menor int
	var media float64
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Digite um número inteiro (ou 0 para sair): ")
		if !scanner.Scan() {
			fmt.Println("Erro ao ler entrada, tente novamente.")
			continue
		}

		// Lê a entrada como string e converte para inteiro
		entrada := strings.TrimSpace(scanner.Text())
		numero, err := strconv.Atoi(entrada)
		if err != nil {
			fmt.Println("Entrada inválida! Por favor, digite um número inteiro.")
			continue
		}

		// Condição para sair do loop
		if numero == 0 {
			break
		}

		numeros = append(numeros, numero)
	}

	if len(numeros) > 0 {
		// Inicializar maior e menor com o primeiro elemento da lista
		maior, menor = numeros[0], numeros[0]

		for _, num := range numeros {
			soma += num
			if num > maior {
				maior = num
			}
			if num < menor {
				menor = num
			}
		}
		media = float64(soma) / float64(len(numeros))

		fmt.Printf("A média dos números é: %.2f\n", media)
		fmt.Printf("O maior número é: %d\n", maior)
		fmt.Printf("O menor número é: %d\n", menor)
	} else {
		fmt.Println("Nenhum número foi digitado.")
	}
}
