package utils

import (
	"context"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

func GetAuth0IdFromContext(ctx context.Context) string {
	token := ctx.Value(jwtmiddleware.ContextKey{})

	if token == nil {
		return ""
	}

	claims := token.(*validator.ValidatedClaims)
	return claims.RegisteredClaims.Subject
}
