package main

import (
   "fmt"
   "os"
   "net"
   "time"
)

func main() {
   service := ":1200"
   udpAddr, err := net.ResolveUDPAddr("udp", service)
   checkError(err)

   for {
      handleClient(conn)
   }
}


