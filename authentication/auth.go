package authentication

import (
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
	"bankapp/models"
	"bankapp/validators"

)

func SignUp(db *sql.DB, user models.User) bool {
	msg, ok :=  validateUser(db, user)
	fmt.Println(msg)
	return ok
}

func Login(db *sql.DB, user models.User) bool {
    msg, ok := validateLogin(db, user.Email, user.Password)
    fmt.Println(msg)
    return ok
}

func validateUser(db *sql.DB, user models.User) (string, bool) {
	var msg string
	uExists, err := userExists(db, user.Email)
	if err != nil {
		return "Erro ao validar usuário", false
	}
	if uExists {
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


func validateLogin(db *sql.DB, email, password string) (string, bool) {
    query := `SELECT 1 FROM users WHERE email = ? AND password = ?`

    var exists int
    err := db.QueryRow(query, email, password).Scan(&exists)
    if err != nil {
        if err == sql.ErrNoRows {
            msg := fmt.Sprintf("Usuário %s ou senha não encontrados", email)
            return msg, false
        }
        return "Erro ao validar usuário", false
    }

    msg := fmt.Sprintf("Bem-vindo novamente ao GoBank %s", email)
    return msg, true
}

func userExists(db *sql.DB, email string) (bool, error) {
	query := `SELECT 1 FROM users WHERE email = ?`

	var exists int
	err := db.QueryRow(query, email).Scan(&exists)

	if err != nil {
		if err == sql.ErrNoRows{
			return false, nil
		}
		return false, err
	}
	return true, nil
}

