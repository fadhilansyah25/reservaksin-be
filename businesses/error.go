package businesses

import "errors"

var (
	ErrInternalServer           = errors.New("something gone wrong, contact administrator")
	ErrNotFound                 = errors.New("data not found")
	ErrIDNotFound               = errors.New("id not found")
	ErrDuplicateData            = errors.New("duplicate data")
	ErrUsernamePasswordNotFound = errors.New("incorrect (Username) or (Password)")
	ErrEmailPasswordNotFound    = errors.New("incorrect (email) or (Password)")
	ErrNIKPasswordNotFound      = errors.New("incorrect (NIK) or (Password)")
)
