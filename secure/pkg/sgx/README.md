## CGO wrapper for Intel SGX

[![license](https://img.shields.io/github/license/mashape/apistatus.svg?style=flat-square)]()

Sgx allows the running of arbitrary golang binaries in intel SGX enclaves,
and provides some tooling for remote code attestation


Run the example:
```
$ source $SGX_SDK/environment # not needed when you already have it
$ make cgo
```

Featurs:
```
# Features
Following features are demonstrated, running inside enclave.
- Monotonic counter
- ECDSA: private/public key generation, signing, verifying
- SHA256
```

# Reference
- [hello-enclave](https://github.com/digawp/hello-enclave)
- [go-with-sgx](https://github.com/mjkim/go-with-intel-sgx)