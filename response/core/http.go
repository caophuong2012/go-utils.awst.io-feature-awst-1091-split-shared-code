package core

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/awst-global/go-utils.awst.io"
)

type ArgumentHttp struct {
	W                http.ResponseWriter
	HttpStatusCode   int
	OptionStatusCode int
	Headers          map[string]string
	DataResponse
}

type DataResponse struct {
	Data          interface{}    `json:"data"`
	ErrorResponse *ErrorResponse `json:"error"`
	CorrelationID string         `json:"correlation_id"`
}

// ErrorResponse
type ErrorResponse struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message"`
}

// RespondJSON writes JSON as http response
func respondJSON(ctx context.Context, w http.ResponseWriter, parameterResponse ArgumentHttp) {
	correlationId, _ := utils.GetAuth0IDAndCorrelationIdFromContext(ctx)
	parameterResponse.DataResponse.CorrelationID = correlationId

	bytes, err := json.Marshal(parameterResponse.DataResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for key, value := range parameterResponse.Headers {
		w.Header().Set(key, value)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(parameterResponse.HttpStatusCode)
	w.Write(bytes)
}

func SuccessJSON(ctx context.Context, httpStatusCode int, parameterResponse ArgumentHttp, object interface{}) {

	parameterResponse.HttpStatusCode = httpStatusCode
	parameterResponse.DataResponse = DataResponse{
		Data: object,
	}

	respondJSON(ctx, parameterResponse.W, parameterResponse)
}

func ErrorJSON(ctx context.Context, httpStatusCode int, parameterResponse ArgumentHttp, err error) {
	parameterResponse.HttpStatusCode = httpStatusCode
	parameterResponse.DataResponse = DataResponse{
		ErrorResponse: &ErrorResponse{
			Code:    parameterResponse.OptionStatusCode,
			Message: err.Error(),
		},
	}

	respondJSON(ctx, parameterResponse.W, parameterResponse)
}
