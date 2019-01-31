package main

import "github.com/kgoval/erresp"

const(
	LangID = "ID"
	LangEN = "EN"
	ErrInvalidPhoneFormat = "ErrInvalidPhoneFormat"
	ErrInternalServer = "ErrInternalServer"
)

var ErrorMessage = map[string]*erresp.MessageBody{
	ErrInvalidPhoneFormat: &erresp.MessageBody{
		Lang: map[string]string{
			LangID: "Format nomor handphone salah",
			LangEN: "Invalid phone format",
		},
		Code: 400,
	},
	ErrInternalServer: &erresp.MessageBody{
		Lang: map[string]string{
			LangID: "Format nomor handphone salah",
			LangEN: "Invalid phone format",
		},
		Code: 500,
	},
}
