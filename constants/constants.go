package constants

type DatabaseError string
type Permission string

// Context Key
type CorrelationIdContextKey struct{}
type Auth0IdContextKey struct{}

const (
	PREFIX_URL_V1 = "/api/v1"
	CONTENT_TYPE  = "Content-Type"

	// Header Keys
	CORRELATION_ID_HEADER_KEY = "X-Correlation-Id"
	AUTH0_ID_HEADER_KEY       = "X-Auth0-Id"
	CREATOR_ID_HEADER_KEY     = "Creator-Id"

	// Log keys
	LOG_CORRELATION_ID = "Correlation-Id"
	LOG_AUTH0_ID       = "Auth0-Id"

	// DB errors
	NO_RECORD_FOUND DatabaseError = "sql: no rows in result set"

	// Permission
	// collection domain
	CREATE_COLLECTION Permission = "collection_domain:create:collection"
	UPDATE_COLLECTION Permission = "collection_domain:update:collection"
	DELETE_COLLECTION Permission = "collection_domain:delete:collection"
	CREATE_ASSET      Permission = "collection_domain:create:asset"

	// uploader domain
	UPLOAD_ASSET Permission = "uploader_domain:create:assets"

	// creator domain
	CREATOR_CREATE_STOREFRONT Permission = "creator_domain:create:storefront"
	CREATOR_UPDATE_STOREFRONT Permission = "creator_domain:update:storefront"

	// identity domain
	IDENTITY_UPDATE_CREATOR Permission = "identity_domain:update:creator"

	// composite domain
	COMPOSITE_CREATE_CONTRACT Permission = "composite_domain:create:contract"

	// Endpoint Keys
	ENDPOINTKEY_URL    = "url"
	ENDPOINTKEY_METHOD = "method"

	// Domains
	AWSTIO_COLLECTION_DOMAIN  = "AWSTIO_COLLECTION_DOMAIN"
	AWSTIO_CREATOR_DOMAIN     = "AWSTIO_CREATOR_DOMAIN"
	AWSTIO_IDENTITY_DOMAIN    = "AWSTIO_IDENTITY_DOMAIN"
	AWSTIO_UPLOADER_DOMAIN    = "AWSTIO_UPLOADER_DOMAIN"
	AWSTIO_COMPOSITE_DOMAIN   = "AWSTIO_COMPOSITE_DOMAIN"
	AWSTIO_BLOCK_CHAIN_DOMAIN = "AWSTIO_BLOCK_CHAIN_DOMAIN"

	// Asset types
	ASSET_IMAGE = "image"
	ASSER_VIDEO = "video"

	// Alias for tables
	CollectorAlias = "ct"
	UserAlias      = "ur"

	// sql sort type
	SORT_TYPE_DESC = "DESC"
	SORT_TYPE_ASC  = "ASC"
)

var SupportedAssetTypes = map[string]string{
	"image/jpg":  ASSET_IMAGE,
	"image/jpeg": ASSET_IMAGE,
	"image/png":  ASSET_IMAGE,
	"image/gif":  ASSET_IMAGE,
	"image/svg":  ASSET_IMAGE,
	"video/mp4":  ASSER_VIDEO,
	"video/webm": ASSER_VIDEO,
	"video/mp3f": ASSER_VIDEO,
}
