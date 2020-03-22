package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	curr "github.com/vladimirvivien/go-networking/currency/lib"
)

const prompt = "currency"

// This porgram is a client implementation for the currency service
// program.  It sends JSON-encoded requests, i.e. {"Get":"USD"},
// and receives a JSON-encoded array of currency information directly
// over TCP or unix domain socket.
//
// Focus:
// This version of the client program highlights the use of
// IO streaming, data serialization, and client-side error handling.
//
// Usage: client [options]
// options:
//  - e service endpoint or socket path, default localhost:4443
//  - n network protocol name [tcp,unix], default tcp
//
// Once started a prompt is provided to interact with service.
