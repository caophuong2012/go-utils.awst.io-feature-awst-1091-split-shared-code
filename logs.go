package utils

import (
	"context"

	"github.com/awst-global/go-utils.awst.io/constants"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func LogInfo(ctx context.Context) *zerolog.Event {
	correlationId, auth0Id := GetAuth0IDAndCorrelationIdFromContext(ctx)

	return log.Info().
		Str(constants.LOG_CORRELATION_ID, correlationId).
		Str(constants.LOG_AUTH0_ID, auth0Id)
}

func LogDebug(ctx context.Context) *zerolog.Event {
	correlationId, auth0Id := GetAuth0IDAndCorrelationIdFromContext(ctx)

	return log.Debug().
		Str(constants.LOG_CORRELATION_ID, correlationId).
		Str(constants.LOG_AUTH0_ID, auth0Id)
}

func LogError(ctx context.Context, err error) *zerolog.Event {
	correlationId, auth0Id := GetAuth0IDAndCorrelationIdFromContext(ctx)

	return log.Error().
		Str(constants.LOG_CORRELATION_ID, correlationId).
		Str(constants.LOG_AUTH0_ID, auth0Id).
		Err(err)
}
