package user

import "errors"

var ErrInvalidEmail = errors.New("Invalid e-mail.")
var ErrEmptyFields = errors.New("All fields are required.")
var ErrEmailAlreadyExists = errors.New("This e-mail already exists.")
var ErrGeneric = errors.New("We can't create your account. Please, try again.")