package validators

import (
	"fmt"
	"net/mail"
	_ "modernc.org/sqlite"
)

func ValidateEmail(email string) (string, bool) {
	_, err := mail.ParseAddress(email)
	if err != nil {
        msg := fmt.Sprintf("Email %s inválido, por favor digite um email válido", email)
        return msg, false
    }
	return "Email válido", true
}