package database

import "errors"

var (
	ErrMissingField                        = errors.New(`field missing on object`)
	ErrExpectingPointer                    = errors.New(`argument must be an address`)
	ErrExpectingSlicePointer               = errors.New(`argument must be a slice address`)
	ErrExpectingSliceMapStruct             = errors.New(`argument must be a slice address of maps or structs`)
	ErrExpectingMapOrStruct                = errors.New(`argument must be either a map or a struct`)
	ErrExpectingPointerToEitherMapOrStruct = errors.New(`expecting a pointer to either a map or a struct`)
)


