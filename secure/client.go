package secure

import (
	"github.com/CIP-NL/go/secure/api/v1/proto/generated"
)

// Client can be used to verify a server is currently running within an attestated enclave
type Client interface {
	v1.AttestationClient
}
