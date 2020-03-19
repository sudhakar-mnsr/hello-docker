package main

import (
	"bufio"
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
// request is larger than the internal buffer. This version uses the bufio
// package to use buffered readers to stream from net.Conn.
//
// Testing:
// Netcat or telnet can be used to test this server by connecting and
// sending command using the format described above.
//
// Usage: server0 [options]
// options:
//   -e host endpoint, default ":4040"
//   -n network protocol [tcp,unix], default "tcp"
func main() {
   var addr string
   var network string
   flag.StringVar(&addr, "e", ":4040", "service endpoint")
   flag.StringVar(&network, "n", "tcp", "network protocol")
   flag.Parse()

   // validate supported network protocols
   switch network {
   case "tcp", "tcp4", "tcp6", "unix":
   default:
      log.Fatalln("unsupported network protocol:", network)
   }
   
   // create a listener for provided network and host address
   ln, err := net.Listen(network, addr)
   if err != nil {
      log.Fatal("failed to create listener:", err)
   }

