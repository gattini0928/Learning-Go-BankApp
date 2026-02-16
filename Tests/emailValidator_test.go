package tests

import (
	"bankapp/validators"
	"testing"
)

func TestValidateEmail(t *testing.T) {
	testCases := []struct {
		name string
		input string
		wantMsg string
		wantOk bool
		
	} {
		{
			name: "Invalid Email Don't Have @", 
			input :"gabriel.com",
			wantMsg: "Email gabriel.com inválido, por favor digite um email válido", 
			wantOk: false,
		},
		{
			name: "Invalid Email Don't Have .", 
			input :"gabriel@gmailcom", 
			wantMsg: "Email gabriel@gmailcom inválido, por favor digite um email válido", 
			wantOk: false,
		},
		{name: "Valid Email", 
			input :"teste@gmail.com", 
			wantMsg: "Email válido", 
			wantOk: true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			gotMsg, gotOk := validators.ValidateEmail(tt.input)
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