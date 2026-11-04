package response

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
