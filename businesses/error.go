package businesses

import "errors"

var (
	ErrInternalServer           = errors.New("something gone wrong, contact administrator")
	ErrNotFound                 = errors.New("data not found")
	ErrIDNotFound               = errors.New("id not found")
	ErrDuplicateData            = errors.New("duplicate data")
	ErrUsernamePasswordNotFound = errors.New("incorrect (Username) or (Password)")
	ErrEmailOrNIKNotFound       = errors.New("incorrect (Email) or (NIK)")
	ErrIncorrectPassword        = errors.New("incorrect (Password)")
)
