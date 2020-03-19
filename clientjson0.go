package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"

	curr "currency"
)

const prompt = "currency"

// This program is a client implementation for the currency service
// program. It sends JSON-encoded requests, i.e {"Get":"USD"},
// and receives a JSON-encoded array of currency information directly
// over TCP or unix domain socket.
// 
// Focus:
// This version of the client program highlights the use fo
// IO streaming, data serialization, and client-side error handling.
// 
// Usage: client [options]
// options:
//    -e service endpoint or socket path, default localhost:4040
//    -n network protocol name [tcp, unix], default tcp
// 
// Once started a prompt is provided to interact with service.
//  
func main() {
   // setup flags
   var add string
   var network string
   flag.StringVar(&addr, "e", "localhost:4040", "service endpoint")
   flag.StringVar(&network, "n", "tcp", "network protocol")
   flag.Parse()

   // dial connection
   conn, err := net.Dial(network, addr)
   if err != nil {
      fmt.Println("failed to create socket:", err)
      os.Exit(1)
   }

   defer conn.Close()
   fmt.Println("connected to currency service:", addr)

   var param string

   // start REPL
   for {
      fmt.Println("Enter search string or *")
      fmt.Print(prompt, "> ")
      _, err = fmt.Scanf(%s", &param)
      if err != nil {
         fmt.Println(Usage: <search string or *>")
         continue
      }
      req := curr.CurrencyRequest{Get: param}

      // Send request:
      // use json encoder to encode value of type curr.CurrencyRequest
      // and stream it to the server via net.Conn.
      if err := json.NewEncoder(conn).Encode(&req); err != nil {
         switch err := err.(type) {
         case net.Error:
            fmt.Println("failed to send request:", err)
            os.Exit(1)
         default:
            fmt.Println("failed to encode request:", err)
            continue
         }
      }
