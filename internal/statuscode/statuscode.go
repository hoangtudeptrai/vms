package statuscode

const (
	// Success codes
	StatusSuccess = 200

	// Error codes
	StatusBadRequest          = 400
	StatusUnauthorized        = 401
	StatusForbidden           = 403
	StatusNotFound            = 404
	StatusMethodNotAllowed    = 405
	StatusConflict            = 409
	StatusInternalServerError = 500

	// Business error codes
	StatusCreateItemFailed = 1001
	StatusReadItemFailed   = 1002
	StatusUpdateItemFailed = 1003
	StatusDeleteItemFailed = 1004
)
