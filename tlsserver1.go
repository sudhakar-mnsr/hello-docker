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

func main() {
	// setup flags
	var addr, network, cert, key, ca string
	flag.StringVar(&addr, "e", "localhost:4443", "service endpoint [ip addr or socket path]")
	flag.StringVar(&network, "n", "tcp", "network protocol [tcp,unix]")
	flag.StringVar(&cert, "cert", "../certs/client-cert.pem", "public cert")
	flag.StringVar(&key, "key", "../certs/client-key.pem", "private key")
	flag.StringVar(&ca, "ca", "../certs/ca-cert.pem", "root CA certificate")
	flag.Parse()

	cer, err := tls.LoadX509KeyPair(cert, key)
	if err != nil {
		log.Fatal(err)
	}

	// Load our CA certificate
	caCert, err := ioutil.ReadFile(ca)
	if err != nil {
		log.Fatal("failed to read CA cert", err)
	}
