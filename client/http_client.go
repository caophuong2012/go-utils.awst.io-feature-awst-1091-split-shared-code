package client

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/awst-global/go-utils.awst.io"
	"github.com/awst-global/go-utils.awst.io/constants"
)

type DataRequest struct {
	Url     string
	Method  string
	Body    io.Reader
	Headers map[string]string
}

type HttpClient struct {
	Client *http.Client
	Req    *DataRequest
	Ctx    context.Context
}

type Response struct {
	StatusCode   int
	ResponseByte []byte
	Header       http.Header
	Data         interface{}
}

type CoreServiceResponse struct {
	Data          interface{} `json:"data"`
	Error         interface{} `json:"error"`
	CorrelationID string      `json:"correlation_id"`
}

type InternalService struct {
	Host string
}

var internalServices = map[WebService]InternalService{
	IdentityManagementService: {
		Host: utils.GetWithDefault(constants.AWSTIO_IDENTITY_DOMAIN, "http://localhost:3000"),
	},
	CreatorManagementService: {
		Host: utils.GetWithDefault(constants.AWSTIO_CREATOR_DOMAIN, "http://localhost:3001"),
	},
	UploaderManagementService: {
		Host: utils.GetWithDefault(constants.AWSTIO_UPLOADER_DOMAIN, "http://localhost:3002"),
	},
	CollectionManagementService: {
		Host: utils.GetWithDefault(constants.AWSTIO_COLLECTION_DOMAIN, "http://localhost:3003"),
	},
	CompositeManagementService: {
		Host: utils.GetWithDefault(constants.AWSTIO_COMPOSITE_DOMAIN, "http://localhost:3004"),
	},
	BlockChainManagementService: {
		Host: utils.GetWithDefault(constants.AWSTIO_BLOCK_CHAIN_DOMAIN, "http://localhost:3005"),
	},
}

func NewHttpClient(ctx context.Context, endpoint map[string]string, body io.Reader, headers map[string]string) *HttpClient {
	endpointHostData, ok := internalServices[WebService(endpoint[WebServiceType])]
	if !ok {
		panic("Invalid data host")
	}

	endpointSet := endpointHostData.Host + endpoint[constants.ENDPOINTKEY_URL]
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	if headers == nil {
		headers = make(map[string]string)
	}

	correlationIdContext := ctx.Value(constants.CorrelationIdContextKey{})
	if correlationIdContext != nil {
		headers[constants.CORRELATION_ID_HEADER_KEY] = correlationIdContext.(string)
	}
	auth0IdContext := ctx.Value(constants.Auth0IdContextKey{})
	if auth0IdContext != nil {
		headers[constants.AUTH0_ID_HEADER_KEY] = auth0IdContext.(string)
	}

	req := &DataRequest{
		Url:     endpointSet,
		Method:  endpoint[constants.ENDPOINTKEY_METHOD],
		Body:    body,
		Headers: headers,
	}

	timeout, err := time.ParseDuration(utils.GetWithDefault("AWSTIO_HTTP_CONN_TIMEOUT", "10s"))
	if err != nil {
		log.Panicf(err.Error())
	}

	return &HttpClient{
		Client: &http.Client{Timeout: timeout, Transport: transport},
		Req:    req,
		Ctx:    ctx,
	}
}

func (httpClient *HttpClient) newRequest() (*http.Request, error) {
	req, err := http.NewRequest(httpClient.Req.Method, httpClient.Req.Url, httpClient.Req.Body)
	if err != nil {
		return nil, err
	}

	for key, value := range httpClient.Req.Headers {
		req.Header.Set(key, value)
	}

	return req.WithContext(httpClient.Ctx), nil
}

func (httpClient *HttpClient) do(req *http.Request) (*http.Response, error) {
	return httpClient.Client.Do(req)
}

func (httpClient *HttpClient) RequestAPI() (*Response, error) {
	req, err := httpClient.newRequest()
	if err != nil {
		return nil, err
	}

	res, err := httpClient.do(req)
	if err != nil {
		return nil, err
	}

	responseByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var payload CoreServiceResponse
	if err := json.Unmarshal(responseByte, &payload); err != nil {
		return nil, err
	}

	return &Response{
		StatusCode:   res.StatusCode,
		ResponseByte: responseByte,
		Header:       res.Header,
		Data:         payload.Data,
	}, nil
}

//SetParamEndpoint sets params for endpoint
func (httpClient *HttpClient) SetParamEndpoint(urlparam []interface{}) {
	httpClient.Req.Url = fmt.Sprintf(httpClient.Req.Url, urlparam...)
}

// SetQueryParamEndpoint returns the full URL with query params
func (httpClient *HttpClient) SetQueryParamEndpoint(queryParam url.Values) {
	// Add query params
	if q := queryParam.Encode(); q != "" {
		sep := "?"
		if strings.Contains(httpClient.Req.Url, "?") {
			sep = "&"
		}
		httpClient.Req.Url = httpClient.Req.Url + sep + q
	}
}
