package validators

import (
	"strconv"
	"strings"
)

func ValidateCpf(cpf string) (string, bool) {
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")
	cpf = strings.ReplaceAll(cpf, " ", "")

	if len(cpf) != 11 {
		return "Seu CPF precisa ter 11 dígitos",false
	}
	_, err := strconv.Atoi(cpf) 
	if err != nil {
		return "CPF deve conter apenas números", false
	}

	if cpf == "00000000000" || cpf == "11111111111" || 
       cpf == "22222222222" || cpf == "33333333333" ||
       cpf == "44444444444" || cpf == "55555555555" ||
       cpf == "66666666666" || cpf == "77777777777" ||
       cpf == "88888888888" || cpf == "99999999999" {
        return "CPF inválido", false
    }

	return "CPF válido", true
}