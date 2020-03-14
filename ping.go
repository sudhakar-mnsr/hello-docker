package main

import (
        "bytes"
        "fmt"
        "io"
        "net"
        "os"
)

const myIPAddress = "192.168.1.2"
const ipv4HeaderSize = 20

func main() {
   if len(os.Args) != 2 {
      fmt.Println("Usage: ", os.Args[0], "host")
      os.Exit(1)
   }
   localAddr, err := net.ResolveIPAddr("ip4", myIPAddress)
   if err != nil {
      fmt.Println("Resolution error", err.Error())
      os.Exit(1)
   }

   remoteAddr, err := net.ResolveIPAddr(ip4", os.Args[1])
   if err != nil {
      fmt.Println("Resolution error", err.Error())
      os.Exit(1)
   }

   conn, err := net.DialIP("ip4:icmp", localAddr, remoteAddr)
   checkError(err)
