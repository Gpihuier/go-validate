package go_validate

import "errors"

var (
	errorNotFoundTag = errors.New("tab not found")
	errorNotStruct   = errors.New("param must struct")
)
