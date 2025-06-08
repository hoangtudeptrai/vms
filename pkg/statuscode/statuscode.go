package statuscode

const (
	StatusSuccess                = 200
	StatusCommonBackendError     = 1000 // Backend error code start from 1000 to 9999
	StatusBindingInputJsonFailed = 400
	StatusCreateItemFailed       = 500
	StatusReadItemFailed         = 500
	StatusUpdateItemFailed       = 500
	StatusDeleteItemFailed       = 500
	StatusSearchItemFailed       = 1005
	StatusItemNotFound           = 1006
	StatusServerError            = 1007
	StatusAuthenticationFailed   = 1008
	StatusUnauthorized           = 401
	StatusInternalServerError    = 500
)
