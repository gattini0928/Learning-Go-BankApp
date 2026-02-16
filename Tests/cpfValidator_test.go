package tests

import (
	"bankapp/validators"
	"testing"
)

func TestValidateCpf(t *testing.T) {
	testCases := []struct {
		name string
		input string
		wantMsg string
		wantOk bool
		
	} {
		{
			name: "CPF > 11 Length", 
			input :"039409902011",
			wantMsg: "Seu CPF precisa ter 11 dígitos", 
			wantOk: false,
		},
		{
			name: "CPF < 11 Length", 
			input :"1092012345", 
			wantMsg: "Seu CPF precisa ter 11 dígitos", 
			wantOk: false,
		},
		{name: "CPF Only 1 digit", 
			input :"11111111111", 
			wantMsg: "CPF inválido", 
			wantOk: false,
		},
		{name: "CPF Valid", 
			input :"12930080899", 
			wantMsg: "CPF válido", 
			wantOk: true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			gotMsg, gotOk := validators.ValidateCpf(tt.input)
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