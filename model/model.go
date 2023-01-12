package model

import "errors"

var (
	// ErrPersonCanNotBeNil the person can not be Nil
	ErrPersonCanNotBeNil = errors.New("Person can't be nil")

	// ErrPersonDoesNotExist the person does not exist
	ErrPersonDoesNotExist = errors.New("Person does not exist")
)
