package core

import (
	"net/http"

	"github.com/awst-global/go-utils.awst.io/client"
	"github.com/awst-global/go-utils.awst.io/constants"
)

var (
	GET_COLLECTION_DEPLOY_CONTRACT_STATUS map[string]string = map[string]string{
		client.WebServiceType:        string(client.BlockChainManagementService),
		constants.ENDPOINTKEY_URL:    constants.PREFIX_URL_V1 + "/chain/contract/collection/%v/status",
		constants.ENDPOINTKEY_METHOD: http.MethodGet,
	}

	DEPLOY_CONTRACT map[string]string = map[string]string{
		client.WebServiceType:        string(client.BlockChainManagementService),
		constants.ENDPOINTKEY_URL:    constants.PREFIX_URL_V1 + "/chain/contract",
		constants.ENDPOINTKEY_METHOD: http.MethodPost,
	}
)
