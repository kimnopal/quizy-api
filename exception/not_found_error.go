package exception

type NotFoundError struct {
	ErrorInfo string
}

func (notFound *NotFoundError) Error() string {
	return notFound.ErrorInfo
}

func NewNotFoundError(err string) NotFoundError {
	return NotFoundError{
		ErrorInfo: err,
	}
}
