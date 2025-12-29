package errorsmessages

import "errors"

var (
	ErrInvalidEmail       = errors.New("Invalid email")
	ErrUserNotFound       = errors.New("User not found")
	ErrUserAlreadyExist   = errors.New("Email already exists")
	ErrInvalidId          = errors.New("Invalid id! ID missing...")
	ErrInvalidCredentials = errors.New("Invalid email or password")
	ErrInternalServer     = errors.New("Something went wrong, try again later...")
	ErrInvalidPassword    = errors.New("Invalid password, password is blank or has less than 4 characters")
)
