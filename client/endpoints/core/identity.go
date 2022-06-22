package core

import (
	"net/http"

	"github.com/awst-global/go-utils.awst.io/client"
	"github.com/awst-global/go-utils.awst.io/constants"
)

var (
	// Identity management service
	ENDPOINT_AUTHENTICATE_CREATOR map[string]string = map[string]string{
		client.WebServiceType:        string(client.IdentityManagementService),
		constants.ENDPOINTKEY_URL:    string(constants.PREFIX_URL_V1 + "/authentications/creator"),
		constants.ENDPOINTKEY_METHOD: http.MethodPost,
	}
	UPDATE_CREATOR map[string]string = map[string]string{
		client.WebServiceType:        string(client.IdentityManagementService),
		constants.ENDPOINTKEY_URL:    string(constants.PREFIX_URL_V1 + "/creators/%v"),
		constants.ENDPOINTKEY_METHOD: http.MethodPut,
	}
	GET_CURRENT_USER map[string]string = map[string]string{
		client.WebServiceType:        string(client.IdentityManagementService),
		constants.ENDPOINTKEY_URL:    string(constants.PREFIX_URL_V1 + "/users/current"),
		constants.ENDPOINTKEY_METHOD: http.MethodGet,
	}
	GET_CREATOR_BY_USER_ID map[string]string = map[string]string{
		client.WebServiceType:        string(client.IdentityManagementService),
		constants.ENDPOINTKEY_URL:    string(constants.PREFIX_URL_V1 + "/users/%v/creators"),
		constants.ENDPOINTKEY_METHOD: http.MethodGet,
	}
	GET_CREATOR_BY_AUTH0_ID map[string]string = map[string]string{
		client.WebServiceType:        string(client.IdentityManagementService),
		constants.ENDPOINTKEY_URL:    string(constants.PREFIX_URL_V1 + "/creators/by-auth0-id/%v"),
		constants.ENDPOINTKEY_METHOD: http.MethodGet,
	}

	// Collector
	GET_COLLECTOR_PROFILE map[string]string = map[string]string{
		client.WebServiceType:        string(client.IdentityManagementService),
		constants.ENDPOINTKEY_URL:    string(constants.PREFIX_URL_V1 + "/collectors/profile/%v"),
		constants.ENDPOINTKEY_METHOD: http.MethodGet,
	}
	UPDATE_COLLECTOR_PROFILE map[string]string = map[string]string{
		client.WebServiceType:        string(client.IdentityManagementService),
		constants.ENDPOINTKEY_URL:    string(constants.PREFIX_URL_V1 + "/collectors/profile/%v"),
		constants.ENDPOINTKEY_METHOD: http.MethodPut,
	}
)
