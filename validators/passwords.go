package validators

import (
	"fmt"
	"strings"
	"unicode"
)

func ValidatePassword(password string) (string, bool) {
	specials := func(r rune) bool {
		return unicode.IsSymbol(r) || unicode.IsPunct(r) 
	} 
	rules := map[string]func(rune) bool {
		"letra maíscula": unicode.IsUpper,
		"letra minúscula": unicode.IsLower,
		"dígito": unicode.IsDigit,
		"caractere especial": specials,
	}
	for ruleName, ruleFunc := range rules {
		found := false
		for _, char : range password {
			if ruleFunc(char) {
				found = true
				break
			}
		}
		if !found {
			msg := fmt.Sprintf("Sua senha deve conter pelo menos um(a) %s", ruleName)
			return msg, false
		}
	}
	if len(password) < 8 {
		return "Sua senha deve ter pelo menos 8 caracteres", false
	}

	return "Senha válida", true
}