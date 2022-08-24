package ginfinite

type GinfiniteError struct {
	status  int
	message string
}

func (g *GinfiniteError) Error() string {
	return g.message
}

func InternalServerError(message string) *GinfiniteError {
	return &GinfiniteError{
		status:  500,
		message: message,
	}
}

func BadRequest(message string) *GinfiniteError {
	return &GinfiniteError{
		status:  400,
		message: message,
	}
}

func CustomError(status int, message string) *GinfiniteError {
	return &GinfiniteError{
		status:  status,
		message: message,
	}
}
