package validators

import ("strings"
	"unicode"
)

func ValidateName(name string) (string, bool) {
	parts := strings.Fields(name)
	if len(parts) < 2 {
		return "Por favor digite seu nome completo.", false
	}
	if len(parts[0]) < 3 {
		return "Seu primeiro nome precisa ter pelo menos 3 letras.", false
	}
	if parts[1] == parts[0] {
		return "O nome e sobrenome não podem ser iguais", false
	}

	for _, char := range name {
		if !unicode.IsLetter(char) && char != ' ' {
			return "Seu nome só pode conter letras", false
		} 
	}
	return "Nome está correto", true 
}