package commitment

import (
	"github.com/bianjieai/tibc-sdk-go/modules/types"
)

// SubModuleName is the error codespace
const SubModuleName string = "commitment"

const moduleName = "tibc" + "-" + SubModuleName

// IBC connection sentinel errors
var (
	ErrInvalidProof       = types.Register(moduleName, 2, "invalid proof")
	ErrInvalidPrefix      = types.Register(moduleName, 3, "invalid prefix")
	ErrInvalidMerkleProof = types.Register(moduleName, 4, "invalid merkle proof")
)
