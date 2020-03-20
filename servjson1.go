package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	curr "currency1"
)

var (
	currencies = curr.Load("./data.csv")
)

// This program implements a simple currency lookup service 
// over TCP or Unix data socket. It loads ISO currency
// information using package curr (see above) and uses a
// simple JSON-encode text-based protocol to exchange data 
// with a client.
// 
// Clients send currency search requests as JSON objects
// as {"GET":"<currency name, code, or country>"}. The request
// data is then unmarshalled to Go type curr.CurrencyRequest 
// using the encoding/json package.
// 
// The request is then used to search the list of currencies
// The search result, a []curr.Currency, is marshalled as JSON
// array of objects and sent to the client.
// 
// Focus:
// This version of the program highlights the use of the encoding
// packages to serialize data to/from Go data types to another
// representation such as JSON. This version uses the encoding/JSON
// package Encoder/Decoder types which are accept in io.Writer and 
// io.Reader respectively. This means they can be used directly with 
// the io.Conn value.
// 
// Testing:
// Netcat can be used for rudimentary testing. However, use clientjsonX
// programs functional tests.
// 
// Usage: server [options]
// options:
//    -e host endpoint, default ":4040"
//    -n network protocol [tcp, unix], default "tcp"
//   
func main() {
   // setup flags
   var addr string
   var network string
   flag.StringVar(&addr, "e", ":4040", "service endpoint")
   flag.StringVar(&network, "n", "tcp", "network protocol")
   flag.Parse()

   // validate supported network protocols
   switch network {
   case "tcp", "tcp4", "tcp6", "unix":
   default:
      fmt.Println("unsupported network protocol")
      os.Exit(1)
   }
  
   // create a listener for provided network and host address
   ln, err := net.Listen(network, addr)
   if err != nil {
      log.Println(err)
      os.Exit(1)
   }
