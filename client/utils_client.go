package client

type WebService string

const (
	CollectionManagementService WebService = "CollectionManagementService"
	CreatorManagementService    WebService = "CreatorManagementService"
	IdentityManagementService   WebService = "IdentityManagementService"
	UploaderManagementService   WebService = "UploaderManagementService"
	CompositeManagementService  WebService = "CompositeManagementService"
	BlockChainManagementService WebService = "BlockChainManagementService"

	WebServiceType = "WebServiceType"
)
