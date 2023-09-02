package helper

import "quizy-api/exception"

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func PanicIfNotFound(err error) {
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
}
