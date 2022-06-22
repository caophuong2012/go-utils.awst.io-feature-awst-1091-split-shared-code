package core

import (
	"net/http"

	"github.com/awst-global/go-utils.awst.io/client"
	"github.com/awst-global/go-utils.awst.io/constants"
)

var (
	GET_ALL_COLLECTIONS map[string]string = map[string]string{
		client.WebServiceType:        string(client.CollectionManagementService),
		constants.ENDPOINTKEY_URL:    constants.PREFIX_URL_V1 + "/collections",
		constants.ENDPOINTKEY_METHOD: http.MethodGet,
	}
	CREATE_COLLECTION map[string]string = map[string]string{
		client.WebServiceType:        string(client.CollectionManagementService),
		constants.ENDPOINTKEY_URL:    constants.PREFIX_URL_V1 + "/collections",
		constants.ENDPOINTKEY_METHOD: http.MethodPost,
	}
	UPDATE_COLLECTION map[string]string = map[string]string{
		client.WebServiceType:        string(client.CollectionManagementService),
		constants.ENDPOINTKEY_URL:    constants.PREFIX_URL_V1 + "/collections/%v",
		constants.ENDPOINTKEY_METHOD: http.MethodPut,
	}
	DELETE_COLLECTION map[string]string = map[string]string{
		client.WebServiceType:        string(client.CollectionManagementService),
		constants.ENDPOINTKEY_URL:    constants.PREFIX_URL_V1 + "/collections/%v",
		constants.ENDPOINTKEY_METHOD: http.MethodDelete,
	}
	GET_COLLECTION map[string]string = map[string]string{
		client.WebServiceType:        string(client.CollectionManagementService),
		constants.ENDPOINTKEY_URL:    constants.PREFIX_URL_V1 + "/collections/%v",
		constants.ENDPOINTKEY_METHOD: http.MethodGet,
	}
	GET_COLLECTION_BY_CREATOR_ID map[string]string = map[string]string{
		client.WebServiceType:        string(client.CollectionManagementService),
		constants.ENDPOINTKEY_URL:    constants.PREFIX_URL_V1 + "/creators/%v/collections",
		constants.ENDPOINTKEY_METHOD: http.MethodGet,
	}
)
