package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	curr "currency1"
)

var (
	currencies = curr.Load("./data.csv")
)

// This program implements a simple currency lookup service
// over TCP of runix Data Socket. It loads ISO currency
// information using package curr (see above) and uses a simple
// JSON-encode text-nased protocol to exchange data with a 
// client
// 
// Clients send currency search requests as JSON objects as 
// {"Get":"<currency name or country or code>"}. The request
// data is then unmarshalled to Go type curr.CurrencyRequest 
// using the encoding/json package.
// 
// The request is then used to search the list of currencies.
// The search result, a []curr.Currency, is marshalled as JSON 
// array of objects and sent to the client.
// 
// Focus:
// This version of the program highlights the use of the encoding
// packages to serialize data to/from Go data types to another 
// representation such as JSON. The program uses the bufio package
// to stream data to and from the client as was done with the 
// previous version. This time, however, char '}' is used as 
// demarcation instead of '\n'.
// 
// Testing:
// Netcat can be used for rudimentary testing. However use clientjsonX
// programs functional tests.
// 
// Usage: server [options]
// options:
//    -e host endpoint, default ":4040"
//    -n network protocol [tcp, unix], default "tcp"
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
   defer ln.Close()
   log.Println("**** Global Currency Service ****")
   log.Printf("Service started: (%s) %s\n", network, addr)

   // connection loop
   for {
      conn, err := ln.Accept()
      if err != nil {
         log.Println(err)
         conn.Close()
         continue
      }
      log.Println("Connected to ", conn.RemoteAddr())
      go handleConnection(conn)
   }
}

func handleConnection(conn net.Conn) {
   defer func() {
      if err := conn.Close(); err != nil {
         log.Println("error closing connection:", err)
      }
   }()

   reader := bufio.NewReaderSize(conn, 4)

   // command-loop
   for {
      // reader will read bytes until '}' is encounter which
      // should indicate the end of JSON object i.e {"get":"US"}
      buf, err := reader.ReadSlice('}')
      if err != nil {
         if err != io.EOF {
            log.Println("connection read error", err)
            return
         }
      }
      reader.Reset(conn)
      // unmarshal request into value of type curr.CurrencyRequest
      var req curr.CurrencyRequest
      if err := json.Unmarshal(buf, &req); err != nil {
         log.Println("failed to unmarshal request:", err)
         // inform the client of bad request by marshalling
         //curr.CurrencyError value as JSON to the client.
         // This could be wrapped in its own function.
         cerr, jerr := json.Marshal(&curr.CurrencyError{Error: err.Error()})
         if jerr != nil {
            log.Println("failed to marshal CurrencyError:", jerr)
            continue
         }
         if _, werr := conn.Write(cerr); werr != nil {
            log.Println("failed to write to CurrencyError:", werr)
            return
         }
         continue
      }       

      // search currencies, result is []curr.Currency
      result := curr.Find(currencies, req.Get)
  
      // marshal result to JSON array
      rsp, err := json.Marshal(&result)
      if err != nil {
         log.Println("failed to marshal data:", err)
         // Note: fmt.Fprintf prints a raw JSON object directly
         // to the client without encoding.
         if _, err := fmt.Fprintf(conn, `{"currency_error":"internal error"}`); err != nil {
            log.Printf("failed to write to client: %v", err)
            return
         }
         continue
      }

      // write response to client
      if _, err := conn.Write(rsp); err != nil {
         log.Println("failed to write response:", err)
         return
      }
   }
}
