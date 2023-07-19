package model

import (
	"regexp"
)

type Signup struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type SessionID struct {
	SID string `json:"id"`
}
type Note struct {
	ID      uint32 `json:"id"`
	Note    string `json:"note"`
	Session string `json:"-"`
}

func IsValidEmail(email string) bool {
	// Simple email validation using regular expression
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(regex, email)
	return match
}

func IsValidPassword(password string) bool {
	// Password validation rules
	if len(password) < 8 {
		return false
	}
	if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		return false
	}
	if !regexp.MustCompile(`[a-z]`).MatchString(password) {
		return false
	}
	if !regexp.MustCompile(`[0-9]`).MatchString(password) {
		return false
	}
	if !regexp.MustCompile(`[@#!$^]`).MatchString(password) {
		return false
	}
	return true
}
