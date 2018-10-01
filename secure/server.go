package secure

import (
	"github.com/CIP-NL/go/secure/api/v1/proto/generated"
)

type Server interface {
	v1.AttestationServer
	v1.ClientAttestationServer
}

