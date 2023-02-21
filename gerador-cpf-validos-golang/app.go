package main

import (
    "fmt"
	"bufio"
    "math/rand"
	"os"
    "time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Cria um novo scanner para leitura das entradas do usuário
	scanner := bufio.NewScanner(bufio.NewReader((os.Stdin)))

	// Aguarda o usuário pressionar "Enter" e gera um novo CPF a cada vez
	for {
		fmt.Print("Pressione Enter para gerar um novo CPF...")
		scanner.Scan()
		cpf := generateCPF()
		fmt.Println("CPF gerado:", cpf)
	}
}

func generateCPF() string {
    // Gera os nove primeiros dígitos do CPF aleatoriamente
    var cpf []int
    for i := 0; i < 9; i++ {
        cpf = append(cpf, rand.Intn(9))
    }

    // Calcula o primeiro dígito verificador
    sum := 0
    for i := 0; i < 9; i++ {
        sum += cpf[i] * (10 - i)
    }
    digit := 11 - sum%11
    if digit >= 10 {
        digit = 0
    }
    cpf = append(cpf, digit)

    // Calcula o segundo dígito verificador
    sum = 0
    for i := 0; i < 10; i++ {
        sum += cpf[i] * (11 - i)
    }
    digit = 11 - sum%11
    if digit >= 10 {
        digit = 0
    }
    cpf = append(cpf, digit)

    // Retorna o CPF formatado
    return fmt.Sprintf("%d%d%d.%d%d%d.%d%d%d-%d%d", cpf[0], cpf[1], cpf[2], cpf[3], cpf[4], cpf[5], cpf[6], cpf[7], cpf[8], cpf[9], cpf[10])
}
