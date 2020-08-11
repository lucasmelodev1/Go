package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//funcao auxiliar para calcular cada dígito, tanto o penultimo quanto o ultimo
func calculaSubtraendo(cpf string, multiplicador int) int {
	subtrativo := 0
	primeiros9Digitos := cpf[:9]
	for idx, value := range primeiros9Digitos {
		v, _ := strconv.Atoi(string(value))
		subtrativo += v * (multiplicador - idx)
	}
	return subtrativo
}

// funcao que calcula os dígitos do CPF
func calculaDigitosCpf(cpfDigitado string) string {
	if len(cpfDigitado) == 13 {
		penultimoDigito := calculaSubtraendo(cpfDigitado, 10)
		penultimoDigito = 11 - penultimoDigito%11
		if penultimoDigito > 9 {
			penultimoDigito = 0
		}

		ultimoDigito := calculaSubtraendo(cpfDigitado, 11)
		ultimoDigito = 11 - (ultimoDigito+(penultimoDigito*2))%11
		if ultimoDigito > 9 {
			ultimoDigito = 0
		}
		ultimosDigitos := strconv.Itoa(penultimoDigito) + strconv.Itoa(ultimoDigito)
		return ultimosDigitos
	}
	return ""
}

func isCpfValido(cpfInput string) bool {
	ultimosDoisDigitos := calculaDigitosCpf(cpfInput)
	if len(ultimosDoisDigitos) == 2 {
		cpfArray := []string{}
		for _, caractere := range cpfInput {
			cpfArray = append(cpfArray, string(caractere))
		}
		ultimosDigitosCpfInputByte := cpfArray[9] + cpfArray[10]
		ultimosDigitosCpfInput := string(ultimosDigitosCpfInputByte)
		if ultimosDoisDigitos == ultimosDigitosCpfInput {
			if cpfInput[0] == cpfInput[1] && cpfInput[1] == cpfInput[2] && cpfInput[2] == cpfInput[3] &&
				cpfInput[3] == cpfInput[4] && cpfInput[4] == cpfInput[5] && cpfInput[5] == cpfInput[6] &&
				cpfInput[6] == cpfInput[7] && cpfInput[7] == cpfInput[8] && cpfInput[8] == cpfInput[9] &&
				cpfInput[9] == cpfInput[10] {
				return false
			}
			return true
		}
		return false
	}
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite seu CPF: ")
	cpfInput, _ := reader.ReadString('\n')

	if isCpfValido(cpfInput) {
		fmt.Println("CPF valido")
	} else {
		fmt.Println("CPF invalido")
	}
}
