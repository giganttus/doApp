package helpers

import "errors"

var (
	ErrEnvLoad = errors.New("environment variables load failed")
	ErrDbConn  = errors.New("database connection failed")
	ErrRepo    = errors.New("repo error")
	ErrService = errors.New("service error")
	ErrAuth    = errors.New("auth error")
)
