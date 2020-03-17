package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

// This program implements a simple echo server over
// Unix Domain Socket (streaming).  When the server receives a
// request, it returns its content immediately.
//
// Usage:
// echos2 -e <socket_path-endpoint>
func main() {
   var addr string
   flag.StringVar(&addr, "e", "/tmp/echo2.sock", "service endpoint address")
   flag.Parse()

   // create local unix domain socket address endpoint
   laddr, err := net.ResolveUnixAddr("unix", addr)
   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }

