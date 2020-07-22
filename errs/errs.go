package errs

type CustomError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (e *CustomError) Error() string {
	return e.Message
}

type ValidationError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (e *ValidationError) Error() string {
	return e.Message
}
