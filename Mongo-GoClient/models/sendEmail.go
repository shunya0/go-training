package models

type VerificationEmailEvent struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}