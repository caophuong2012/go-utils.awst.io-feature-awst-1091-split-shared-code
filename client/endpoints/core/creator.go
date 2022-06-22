package core

import (
	"net/http"

	"github.com/awst-global/go-utils.awst.io/client"
	"github.com/awst-global/go-utils.awst.io/constants"
)

var (
	ENDPOINT_GET_STORE_FRONT map[string]string = map[string]string{
		client.WebServiceType:        string(client.CreatorManagementService),
		constants.ENDPOINTKEY_URL:    constants.PREFIX_URL_V1 + "/store-front/%v",
		constants.ENDPOINTKEY_METHOD: http.MethodGet,
	}
	ENDPOINT_CREATE_STORE_FRONT map[string]string = map[string]string{
		client.WebServiceType:        string(client.CreatorManagementService),
		constants.ENDPOINTKEY_URL:    constants.PREFIX_URL_V1 + "/store-front/",
		constants.ENDPOINTKEY_METHOD: http.MethodPost,
	}
	ENDPOINT_UPDATE_STORE_FRONT map[string]string = map[string]string{
		client.WebServiceType:        string(client.CreatorManagementService),
		constants.ENDPOINTKEY_URL:    constants.PREFIX_URL_V1 + "/store-front/%v",
		constants.ENDPOINTKEY_METHOD: http.MethodPut,
	}
)
