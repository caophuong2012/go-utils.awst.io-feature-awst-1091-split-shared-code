package core

import (
	"net/http"

	"github.com/awst-global/go-utils.awst.io/client"
	"github.com/awst-global/go-utils.awst.io/constants"
)

var (
	ENDPOINT_UPLOAD_ASSET_TO_S3 map[string]string = map[string]string{
		client.WebServiceType:        string(client.UploaderManagementService),
		constants.ENDPOINTKEY_URL:    constants.PREFIX_URL_V1 + "/assets",
		constants.ENDPOINTKEY_METHOD: http.MethodPost,
	}
)
