package errors

import "errors"

var (
	ErrInvalidArgumentCount = errors.New("Expected 2 arguments")
	ErrZeroSteps            = errors.New("Steps must be greater than zero")
	ErrZeroDuration         = errors.New("Duration must be greater than zero")
	ErrZeroWeight           = errors.New("Weight must be greater than zero")
	ErrZeroHeigth           = errors.New("Height must be greater than zero")
	ErrUnknownTraining      = errors.New("Unknown training")
)