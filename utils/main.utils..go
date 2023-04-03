package utils

import (
	"net/http"
)

type SCookieFunc struct {
	Name  string
	Value string
}

func CreateCookie(cookie SCookieFunc) *http.Cookie {
	return &http.Cookie{
		Value:    cookie.Value,
		Name:     cookie.Name,
		Secure:   false,
		HttpOnly: true,
		Path:     "/",
	}
}
