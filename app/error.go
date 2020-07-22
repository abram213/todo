package app

type CustomError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (e *CustomError) Error() string {
	return e.Message
}

//todo: ValidationError
