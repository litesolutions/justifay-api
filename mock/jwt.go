package mock

import (
	"github.com/litesolutions/justifay-api/model"
)

// JWT mock
type JWT struct {
	GenerateTokenFn func(*model.AuthUser) (string, error)
}

// GenerateToken mock
func (j *JWT) GenerateToken(u *model.AuthUser) (string, error) {
	return j.GenerateTokenFn(u)
}
