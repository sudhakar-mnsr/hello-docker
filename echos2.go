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

   // announce service using ListenUnix
   // which creates a UnixListener listener
   l, err := net.ListenUnix("unix", laddr)
   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }
   defer l.Close()
   fmt.Println("listening at (unix)", laddr.String())

   // req/response loop
   for {
      // use UnixListener to block and wait for UDS
      // connection request using AcceptUnix which then
      // creates a UnixConn
      conn, err := l.AcceptUnix()
      if err != nil {
         fmt.Println("failed to accept conn:", err)
         conn.Close()
         continue
      }
      fmt.Println("connected to : ", conn.RemoteAddr())
      go handleConnection(conn)
   }
}
