package authentication

import (
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
	"bankapp/app"
	"bankapp/validators"

)

func SignUp(db *sql.DB, user app.User) bool {
	msg, ok :=  validateUser(db, user)
	if !ok {
		fmt.Println(msg)
		return false
	}
	fmt.Println(msg)
	return true
}

func validateUser(db *sql.DB, user app.User) (string, bool) {
	var msg string
	userExists, err := UserExists(db, user.Email)
	if err != nil {
		return "Erro ao validar usuário", false
	}
	if userExists {
		msg = fmt.Sprintf("O usuário %s já existe, faça login ou tente novamente", user.Email)
		return msg, false
	}
	msg, ok := validators.ValidateName(user.Name)
	if !ok {
		return msg, false
	}
	
	msg, ok = validators.ValidateEmail(user.Email)
    if !ok {
        return msg, false
    }
    
    msg, ok = validators.ValidateCpf(user.Cpf)
    if !ok {
        return msg, false
    }
    
    msg, ok = validators.ValidatePassword(user.Password)
    if !ok {
        return msg, false
    }
    
    msg = fmt.Sprintf("Bem-vindo ao GoBank %s", user.Name)
    return msg, true
}

func UserExists(db *sql.DB, email string) (bool, error) {
	query := `SELECT 1 FROM users WHERE email = ?`

	var exists int
	err := db.QueryRow(query, email).Scan(&exists)

	if err != nil {
		if err == sql.ErrNoRows{
			return false, nil
		}
		return false, nil
	}
	return true, nil
}