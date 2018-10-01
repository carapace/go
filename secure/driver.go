package secure

import (
	"fmt"
	"github.com/CIP-NL/go/secure/api/v1/proto/generated"
	"io"
	"sync"
)

// only a single driver may be registered, since only a single CPU executes
// the binary.
var driverMu sync.RWMutex
var driver Driver

// Register allows for a CPU specific lib to register itself. It should only
// be called in the init of the lib. Register panics if called twice.
func Register(d Driver)  {
	driverMu.Lock()
	defer driverMu.Unlock()
	if driver != nil {
		panic("secure: cpu specific driver has already been registered")
	}
	driver = d
}

// Validate if a driver is registered, and allow for any initializations by the driver to run.
func init()  {
	if driver == nil {
		panic("secure: forgot to register a CPU specific driver?")
	}
	err := driver.Init()
	if err != nil {
		panic(fmt.Sprintf("secure: %s", err))
	}
}

// Driver implements a set of primitives to interact with CPU specific functionalities
// regarding code attestation and data sealing
type Driver interface {
	// Init allows for initialization functions to run outside of the packages init function.
	Init() error

	// Execute runs a function within the secure enclave. Execute is not guaranteed to be suitable
	// for long running processes and will panic if a driver does not support Execute. Run is the preferred method
	// to launch daemon processes.
	Execute(func() error) error

	// Read takes a data-sealed stream and decrypts it using the CPU specific data seal
	Read(io.Reader) (io.Reader, error)

	// Write takes an stream and unencrypted data, and encrypts it on the fly.
	Write(io.Writer, []byte) error

	// Run is the equivalent of execute, however fit for running processes over a long time, such as servers.
	Run(func() error) error

	// Load receives a cert/private key/secret used in the signing of enclave material
	Load(io.Reader) error

	// Serve starts an attestation server. If no servers are provided, the driver will default to the protobuf
	// implementation.
	Serve(...Server) error
}

// Execute runs a file within the secure enclave. Execute is not guaranteed to be suitable
// for long running processes
func Execute(f func() error) error{
	return driver.Execute(f)
}


// Read takes a data-sealed stream and decrypts it using the CPU specific data seal
func Read(stream io.Reader) (io.Reader, error)  {
	return driver.Read(stream)
}

// Write takes an stream and unencrypted data, and encrypts it on the fly.
func Write(stream io.Writer, data []byte) error  {
	return driver.Write(stream, data)
}

// Run is the equivalent of execute, however fit for running processes over a long time, such as servers.
func Run(f func() error) error {
	return driver.Run(f)
}

// Load receives a cert/private key/secret used in the signing of enclave material
func Load(privkey io.Reader) error {
	return driver.Load(privkey)
}

// Serve starts an attestation server. If no servers are provided, the driver will default to the protobuf
// implementation.
func Serve(server ...Server) error {
	return driver.Serve(server...)
}
