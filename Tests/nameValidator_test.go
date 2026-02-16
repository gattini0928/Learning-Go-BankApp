package tests

import ("testing"
	"bankapp/validators"
)

func TestValidateName(t *testing.T) {
	testCases := []struct {
		name    string
		input   string
		wantMsg string
		wantOk  bool
	}{
		{
			name:    "Only First Name",
			input:   "Gabriel",
			wantMsg: "Por favor digite seu nome completo.",
			wantOk:  false,
		},
		{
			name:    "First Name Minimum Length",
			input:   "Ga Gattini",
			wantMsg: "Seu primeiro nome precisa ter pelo menos 3 letras.",
			wantOk:  false,
		},
		{
			name:    "Same Name and Surname",
			input:   "Gabriel Gabriel",
			wantMsg: "O nome e sobrenome não podem ser iguais",
			wantOk:  false,
		},
		{
			name:    "Digits in the Names",
			input:   "Gabriel1 Gattini",
			wantMsg: "Seu nome só pode conter letras",
			wantOk:  false,
		},
		{
			name:    "Full Name",
			input:   "Gabriel Gattini",
			wantMsg: "Nome está correto",
			wantOk:  true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			gotMsg, gotOk := validators.ValidateName(tt.input)
			if gotMsg != tt.wantMsg {
				t.Fatalf("Esperava msg: %v,  mas recebeu msg: %v",
					tt.wantMsg, gotMsg)
			}
			if gotOk != tt.wantOk {
				t.Fatalf("Esperava ok: %v, recebeu ok: %v",
					tt.wantOk, gotOk)
			}
		})
	}
}
