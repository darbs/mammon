package api

import "errors"

// todo custom error obj with url
var (
	ErrApiConnect = errors.New(`failed to connect to api endpoint`)
)
