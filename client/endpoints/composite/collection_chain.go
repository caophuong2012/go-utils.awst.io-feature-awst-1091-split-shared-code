package composite

import (
	"net/http"

	"github.com/awst-global/go-utils.awst.io/client"
	"github.com/awst-global/go-utils.awst.io/constants"
)

var (
	DEPLOY_CONTRACT map[string]string = map[string]string{
		client.WebServiceType:        string(client.CompositeManagementService),
		constants.ENDPOINTKEY_URL:    constants.PREFIX_URL_V1 + "/smart-contracts/",
		constants.ENDPOINTKEY_METHOD: http.MethodPost,
	}
)
