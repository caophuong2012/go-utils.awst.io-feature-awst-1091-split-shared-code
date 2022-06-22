package bff

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/awst-global/go-utils.awst.io"
	"github.com/awst-global/go-utils.awst.io/client"
)

type ArgumentHttp struct {
	W              http.ResponseWriter
	Headers        map[string]string
	DataResponse   []byte
	HttpStatusCode int
}

type SuccessResponse struct {
	Data          interface{} `json:"data"`
	CorrelationID string      `json:"correlation_id"`
}

type ErrorResponse struct {
	Error         *BFFError `json:"error"`
	CorrelationID string    `json:"correlation_id"`
}

// ErrorResponse
type BFFError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message"`
}

// RespondJSON writes JSON as http response
func respondJSON(ctx context.Context, argumentHttp ArgumentHttp) {

	for key, value := range argumentHttp.Headers {
		argumentHttp.W.Header().Set(key, value)
	}

	argumentHttp.W.Header().Set("Content-Type", "application/json; charset=utf-8")
	argumentHttp.W.WriteHeader(argumentHttp.HttpStatusCode)
	argumentHttp.W.Write(argumentHttp.DataResponse)
}

func SuccessJSON(ctx context.Context, argumentHttp ArgumentHttp, clientDataResponse *client.Response) {
	argumentHttp.DataResponse = clientDataResponse.ResponseByte
	argumentHttp.HttpStatusCode = clientDataResponse.StatusCode

	respondJSON(ctx, argumentHttp)
}

func ErrorJSON(ctx context.Context, argumentHttp ArgumentHttp, httpStatusCode int, err error) {
	utils.LogError(ctx, err).Send()

	correlationId, _ := utils.GetAuth0IDAndCorrelationIdFromContext(ctx)

	errorResponse := ErrorResponse{
		Error: &BFFError{
			Message: err.Error(),
		},
		CorrelationID: correlationId,
	}

	bytes, err := json.Marshal(errorResponse)
	if err != nil {
		argumentHttp.W.WriteHeader(http.StatusInternalServerError)
		return
	}

	argumentHttp.DataResponse = bytes
	argumentHttp.HttpStatusCode = httpStatusCode

	respondJSON(ctx, argumentHttp)
}
