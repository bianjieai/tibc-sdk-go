package client

import (
	"fmt"
	"regexp"

	"github.com/bianjieai/tibc-sdk-go/modules/types"
)

var _ types.Height = (*Height)(nil)

// IsRevisionFormat checks if a chainID is in the format required for parsing revisions
// The chainID must be in the form: `{chainID}-{revision}
// 24-host may enforce stricter checks on chainID
var IsRevisionFormat = regexp.MustCompile(`^.*[^-]-{1}[1-9][0-9]*$`).MatchString

// ZeroHeight is a helper function which returns an uninitialized height.
func ZeroHeight() Height {
	return Height{}
}

// NewHeight is a constructor for the IBC height type
func NewHeight(revisionNumber, revisionHeight uint64) Height {
	return Height{
		RevisionNumber: revisionNumber,
		RevisionHeight: revisionHeight,
	}
}

// GetRevisionNumber returns the revision-number of the height
func (h Height) GetRevisionNumber() uint64 {
	return h.RevisionNumber
}

// GetRevisionHeight returns the revision-height of the height
func (h Height) GetRevisionHeight() uint64 {
	return h.RevisionHeight
}

// String returns a string representation of Height
func (h Height) String() string {
	return fmt.Sprintf("%d-%d", h.RevisionNumber, h.RevisionHeight)
}

// Decrement will return a new height with the RevisionHeight decremented
// If the RevisionHeight is already at lowest value (1), then false success flag is returend
func (h Height) Decrement() (decremented types.Height, success bool) {
	if h.RevisionHeight == 0 {
		return Height{}, false
	}
	return NewHeight(h.RevisionNumber, h.RevisionHeight-1), true
}

// Increment will return a height with the same revision number but an
// incremented revision height
func (h Height) Increment() types.Height {
	return NewHeight(h.RevisionNumber, h.RevisionHeight+1)
}

// IsZero returns true if height revision and revision-height are both 0
func (h Height) IsZero() bool {
	return h.RevisionNumber == 0 && h.RevisionHeight == 0
}
