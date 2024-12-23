package internal

import (
	"github.com/guilhermealegre/go-clean-arch-core-lib/errors"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/errors/config"
)

// Generic error codes
var (
	ErrorGeneric            = config.GetError("1", "%s", errors.Error)
	ErrorInvalidInputFields = config.GetError("2", "The field: %s as invalid value: %v", errors.Info)
	ErrorInvalidLogin       = config.GetError("3", "Invalid login", errors.Info)
	ErrorInvalidEmail       = config.GetError("4", "Invalid User Email", errors.Info)
	ErrorUserIDNotFound     = config.GetError("5", "User ID not found", errors.Info)
	ErrInsufficientFunds    = config.GetError("6", "Insufficient funds", errors.Info)
	ErrorRateLimitExceeded  = config.GetError("7", "Rate limit exceeded", errors.Info)
)
