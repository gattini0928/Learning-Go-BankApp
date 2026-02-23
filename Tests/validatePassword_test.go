package tests

import (
	"bankapp/authentication"
	"testing"
)

func TestValidateAccountPassword(t *testing.T) {
	testCases := []struct {
		name string
		input string
		password string
		wantOk bool
	} {
	   {
		name:"Empty password",
		input:"",
		password: "",
		wantOk: false,
	   },
		{
		name:"Too short",
		input:"12345",
		password: "12345",
		wantOk: false,
	   },
		{
		name:"Too large",
		input:"1234567",
		password: "1234567",
		wantOk: false,
	   },
		{
		name:"With characters",
		input:"12345a",
		password: "12345a",
		wantOk: false,
	   },
		{
		name:"Valid Password",
		input:"123456",
		password: "123456",
		wantOk: true,
	   }, 
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			_, gotOk := authentication.ValidatePassword(tt.input)
			if gotOk != tt.wantOk {
				t.Fatalf("Esperava ok %v, recebeu ok: %v",
				tt.wantOk, gotOk)
			}
		})
	}
}