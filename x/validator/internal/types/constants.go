package types

import (
	"time"
)

const (
	// Default power every validator created with
	Power int64 = 10

	// Zero power is used to demote validator
	ZeroPower int64 = 0

	// Maximum number of validators
	MaxValidators uint16 = 100

	// Maximum time to accept double-sign evidence
	MaxEvidenceAge = 60 * 2 * time.Second

	// Size (number of blocks) of the sliding window used to track validator liveness.
	SignedBlocksWindow   = int64(100)

	// Minimal number of blocks must have been signed per window
	MinSignedPerWindow = int64(50)
)
