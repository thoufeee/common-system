package response


// http status codes
const (
	StatusOK                  = 200
	StatusCreated             = 201
	StatusNoContent            = 204
	StatusBadRequest           = 400
	StatusUnauthorized         = 401
	StatusForbidden            = 403
	StatusNotFound             = 404
	StatusConflict             = 409
	StatusInternalServerError  = 500
)


// respone codes
const (
	CodeSuccess = "SUCCESS"
	CodeCreated = "CREATED"
	CodeUpdated = "UPDATED"
	CodeDeleted = "DELETED"

	CodeBadRequest   = "BAD_REQUEST"
	CodeUnauthorized = "UNAUTHORIZED"
	CodeForbidden    = "FORBIDDEN"
	CodeNotFound     = "NOT_FOUND"
	CodeConflict     = "CONFLICT"
	CodeInternal     = "INTERNAL_ERROR"
)


// /respone message
const (
	MsgSuccess = "Request successful"
	MsgCreated = "Created successfully"
	MsgUpdated = "Updated successfully"
	MsgDeleted = "Deleted successfully"

	MsgBadRequest   = "Invalid request"
	MsgUnauthorized = "Authentication required"
	MsgForbidden    = "Access denied"
	MsgNotFound     = "Resource not found"
	MsgConflict     = "Resource already exists"
	MsgInternal     = "Something went wrong"
)
