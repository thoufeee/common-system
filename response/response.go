package response

// response model
type Response[T any] struct {
	Success bool   `json:"success"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    *T     `json:"data,omitempty"`
}

// for successs response
func Success[T any](code, message string, data T) Response[T] {
	return Response[T]{
		Success: true,
		Code:    code,
		Message: message,
		Data:    &data,
	}
}

// for error response
func Error(code, message string) Response[any] {
	return Response[any]{
		Success: true,
		Code:    code,
		Message: message,
		Data:    nil,
	}
}
