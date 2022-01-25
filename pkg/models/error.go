package models

type Error struct {
	Package    string
	File       string
	Func       string
	Err        string // err.Error()
	Message    string // custom message
	StatusCode int
}

const (
	UserNotFound              = "user not found"
	UserAlreadyRegistered     = "user already registered"
	InvalidPasswordOrUsername = "invalid password or username"
	InvalidId                 = "invalid id"
	EmptyAuthHeader           = "auth header is empty"
	InvalidAuthHeader         = "invalid auth header"
	UserIdNotFound            = "user id not found"
	InvalidListIdParam        = "invalid list id param"
	InvalidPasswordIdParam    = "invalid password id param"
	ErrorWhileInputBinding    = "error while input binding"
	PasswordNotFound          = "password not found"
)
