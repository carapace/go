package secure

import (
	"github.com/CIP-NL/go/secure/api/v1/proto/generated"
	"google.golang.org/grpc"
)

// Client can be used to verify a server is currently running within an attestated enclave
type Client interface {
	v1.AttestationClient
	Dial(host, port string, callopts ...grpc.DialOption)
	Verify() (bool, error) // Check if the server is running within an enclave
}

type AttestedClient interface {
	Client
	v1.ClientAttestationClient
}