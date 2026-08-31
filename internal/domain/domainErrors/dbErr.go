package domainErrors

import "errors"

var ErrPersonNotFound = errors.New("person not found")
var ErrInvalidLoginOrPass = errors.New("invalid login/pass")
