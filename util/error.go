package util

type HttpError struct {
	code int
	err  error
}

func NewHttpError(code int, err error) HttpError {
	return HttpError{
		code: code,
		err:  err,
	}
}

func (e HttpError) Error() string {
	return "http error"
}

func (e HttpError) Unwrap() error {
	return e.err
}

func (e HttpError) Code() int {
	return e.code
}
