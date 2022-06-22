package constants

type Error string

const (
	EMPTY_BODY        Error = "Request body is required"
	PERMISSION_DENIED Error = "Permission denied"
	NOT_FOUND         Error = "Not found"
)
