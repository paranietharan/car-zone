package errorutils

import "errors"

var (
	ErrEngineNotFound      = errors.New("engine not found")
	ErrEngineAlreadyExists = errors.New("engine already exists")
	ErrEngineInvalidID     = errors.New("invalid engine ID")
	ErrEngineCreateFailed  = errors.New("failed to create engine")
	ErrEngineUpdateFailed  = errors.New("failed to update engine")
	ErrEngineDeleteFailed  = errors.New("failed to delete engine")
)
