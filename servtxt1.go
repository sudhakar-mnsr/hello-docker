package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strings"

	curr "currency"
)

var currencies = curr.Load("./data.csv")

// This program implements a simple currency lookup service
// over TCP or Unix Data Socket. It loads ISO currency
// information using package lib (see above) and uses a simple
// text-based protocol to interact with the client and send
// the data.
//
// Clients send currency search requests as a textual command in the form:
//
// GET <currency, country, or code>
//
// When the server receives the request, it is parsed and is then used
// to search the list of currencies. The search result is then printed
// line-by-line back to the client.
//
// Focus:
// This version of the currency server focuses on implementing a streaming
// strategy when receiving data from client to avoid dropping data when the
// request is larger than the internal buffer. This relies on the fact that
// net.Conn implements io.Reader which allows the code to stream data.
//
// Testing:
// Netcat or telnet can be used to test this server by connecting and
// sending command using the format described above.
//
// Usage: server0 [options]
// options:
//   -e host endpoint, default ":4040"
//   -n network protocol [tcp,unix], default "tcp"
