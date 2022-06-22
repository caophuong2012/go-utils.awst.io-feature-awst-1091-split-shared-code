package utils

import (
	"context"
	"net/http"

	"github.com/awst-global/go-utils.awst.io/constants"
)

func GetAuth0IDAndCorrelationIdFromContext(ctx context.Context) (correlationId, userId string) {

	correlationIdContextKey := ctx.Value(constants.CorrelationIdContextKey{})
	if correlationIdContextKey != nil {
		correlationId = correlationIdContextKey.(string)
	}

	auth0IdContextKey := ctx.Value(constants.Auth0IdContextKey{})
	if auth0IdContextKey != nil {
		userId = auth0IdContextKey.(string)
	}

	return correlationId, userId
}

func SetAuth0IDAndCorrelationIdToContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, constants.CorrelationIdContextKey{}, r.Header.Get(constants.CORRELATION_ID_HEADER_KEY))
		ctx = context.WithValue(ctx, constants.Auth0IdContextKey{}, r.Header.Get(constants.AUTH0_ID_HEADER_KEY))
		*r = *r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
