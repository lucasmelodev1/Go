package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"unicode/utf8"
)

func calculaCpfValido(cpfDigitado string) (string, error) {
	if utf8.RuneCountInString(cpfDigitado) == 13 {
		cpfArray := []int{}
		for _, caractere := range cpfDigitado {
			numero, _ := strconv.Atoi(string(caractere))
			cpfArray = append(cpfArray, numero)
		}
		calculoPenultimoDigitoCpf := 11 - (cpfArray[0]*10+cpfArray[1]*9+cpfArray[2]*8+cpfArray[3]*7+cpfArray[4]*6+cpfArray[5]*5+cpfArray[6]*4+cpfArray[7]*3+cpfArray[8]*2)%11
		var penultimoDigitoCpf int
		if calculoPenultimoDigitoCpf > 9 {
			penultimoDigitoCpf = 0
		} else {
			penultimoDigitoCpf = calculoPenultimoDigitoCpf
		}
		calculoUltimoDigitoCpf := 11 - (cpfArray[0]*11+cpfArray[1]*10+cpfArray[2]*9+cpfArray[3]*8+cpfArray[4]*7+cpfArray[5]*6+cpfArray[6]*5+cpfArray[7]*4+cpfArray[8]*3+penultimoDigitoCpf*2)%11
		penultimoDigitoCpfStr := strconv.Itoa(penultimoDigitoCpf)
		var ultimoDigitoCpf string
		if calculoUltimoDigitoCpf > 9 {
			ultimoDigitoCpf = "0"
		} else {
			ultimoDigitoCpf = strconv.Itoa(calculoUltimoDigitoCpf)
		}
		ultimosDigitosCpf := penultimoDigitoCpfStr + ultimoDigitoCpf
		return ultimosDigitosCpf, nil
	}
	return "", errors.New("CPF inválido")
}

func validarNovoCpf(ultimosDoisDigitos, cpfInput string) bool {
	if len(ultimosDoisDigitos) == 2 {
		cpfArray := []string{}
		for _, caractere := range cpfInput {
			cpfArray = append(cpfArray, string(caractere))
		}
		ultimosDigitosCpfInputByte := cpfArray[9] + cpfArray[10]
		ultimosDigitosCpfInput := string(ultimosDigitosCpfInputByte)
		if ultimosDoisDigitos == ultimosDigitosCpfInput {
			if cpfInput[0] == cpfInput[1] && cpfInput[1] == cpfInput[2] && cpfInput[2] == cpfInput[3] && cpfInput[3] == cpfInput[4] && cpfInput[4] == cpfInput[5] && cpfInput[5] == cpfInput[6] && cpfInput[6] == cpfInput[7] && cpfInput[7] == cpfInput[8] && cpfInput[8] == cpfInput[9] && cpfInput[9] == cpfInput[10] {
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
	ultimosDigitos, _ := calculaCpfValido(cpfInput)
	fmt.Println(validarNovoCpf(ultimosDigitos, cpfInput))
}
