package main

const (
	BadRequest = iota
	Ok
)
const (
	EmptyFields  = iota
	UserNotFound
	LoggedIn
)

type Answer struct {
	code   uint8
	status uint
	body   string
}
