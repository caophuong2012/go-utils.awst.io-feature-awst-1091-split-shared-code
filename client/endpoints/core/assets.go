package core

import (
	"net/http"

	"github.com/awst-global/go-utils.awst.io/client"
	"github.com/awst-global/go-utils.awst.io/constants"
)

var (
	CREATE_ASSET map[string]string = map[string]string{
		client.WebServiceType:        string(client.CollectionManagementService),
		constants.ENDPOINTKEY_URL:    constants.PREFIX_URL_V1 + "/assets",
		constants.ENDPOINTKEY_METHOD: http.MethodPost,
	}
	GET_ASSET map[string]string = map[string]string{
		client.WebServiceType:        string(client.CollectionManagementService),
		constants.ENDPOINTKEY_URL:    constants.PREFIX_URL_V1 + "/assets/%v",
		constants.ENDPOINTKEY_METHOD: http.MethodGet,
	}
	GET_ASSET_BY_COLLECTION_ID map[string]string = map[string]string{
		client.WebServiceType:        string(client.CollectionManagementService),
		constants.ENDPOINTKEY_URL:    constants.PREFIX_URL_V1 + "/collections/%v/assets",
		constants.ENDPOINTKEY_METHOD: http.MethodGet,
	}
)
