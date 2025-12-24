package application

import "errors"

var ErrPermissionDenied = errors.New("permission denied")
var ErrNotFound = errors.New("not found")
