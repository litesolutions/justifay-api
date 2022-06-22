package mock

import (
	"context"

	"github.com/litesolutions/justifay-api/model"
)

// Auth mock
type Auth struct {
	GetUserFn func(context.Context) *model.AuthUser
}

// GetUser mock
func (s *Auth) GetUser(c context.Context) *model.AuthUser {
	return s.GetUserFn(c)
}
