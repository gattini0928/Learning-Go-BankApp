package tests

import (
	"bankapp/validators"
	"testing"
)

func TestValidatePassword(t *testing.T) {
	testCases := []struct {
		name string
		input string
		wantMsg string
		wantOk bool
		
	} {
		{
			name: "Password Invalid Length", 
			input :"Tst3*",
			wantMsg: "Sua senha deve ter pelo menos 8 caracteres", 
			wantOk: false,
		},
				{
			name: "Don't Have Uppercase", 
			input :"teste123%", 
			wantMsg: "Sua senha deve conter pelo menos um(a) letra maiúscula", 
			wantOk: false,
		},
		{name: "Don't Have Uppercase", 
			input :"TESTE123%", 
			wantMsg: "Sua senha deve conter pelo menos um(a) letra minúscula", 
			wantOk: false,
		},
		{
			name: "Don't Have Digit", 
			input :"Teste&&%", 
			wantMsg: "Sua senha deve conter pelo menos um(a) dígito", 
			wantOk: false,
		},
		{
			name: "Don't Have Special", 
			input :"Teste1234", 
			wantMsg: "Sua senha deve conter pelo menos um(a) caractere especial", 
			wantOk: false,
		},
		{
			name: "Valid Password", 
			input :"Teste123%", 
			wantMsg: "Senha válida", 
			wantOk: true,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			gotMsg, gotOk := validators.ValidatePassword(tt.input)
			if gotMsg != tt.wantMsg {
				t.Fatalf("Esperava msg: %v, recebeu msg: %v",
				 tt.wantMsg, gotMsg)
			}
			if gotOk != tt.wantOk {
				t.Fatalf("Esperava ok %v, recebeu ok: %v",
				tt.wantOk, gotOk)
			}
		})
	}
}