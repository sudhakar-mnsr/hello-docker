package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

// This program implements a simple echo client over TCP or unix domain socket.
// It sends a text content to the server and displays
// the response on the screen.
//
// Usage:
// echoc3 [flags] <text content>
// flags:
//   -e <address-endpoint>
//   -n <network>
func main() {
   var addr string
   var network string
   flag.StringVar(&addr, "e", "localhost:4040", "service address")
   flag.StringVar(&network, "n", "tcp", "network protocol")
   flag.Parse()
   text := flag.Arg(0)

   // Validate network
   switch network {
   case "tcp", "tcp4", "tcp6", "unix":
   default:
      fmt.Println("unsupported network protocol")
      os.Exit(1)
   }

   // Use function Dial to create a generic connection to the 
   // remote address
   conn, err := net.Dial(network, addr)
   if err != nil {
      fmt.Println("failed to connect to server:", err)
      os.Exit(1)
   }
   defer conn.Close()

   // send text to server 
