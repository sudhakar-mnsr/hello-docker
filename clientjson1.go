package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"learning-go/ch11/curr1"
	"net"
	"os"
	"time"

	curr "currency1"
)

const prompt = "currency"

This program is a client implementation for the currency service
program. It sends JSON-encoded requests, i.e. {"Get":"USD"} and
receives JSON-encoded array of currency information directly
over TCP or unix domain socket.

Focus:
This program highlights the use of IO streaming, data serialization,
and client-side error handling. It also shows how to configure the
dialer to setup settings such as timeout and KeepAlive values. Further,
the code also implements a simple connection-retry strategy when connecting

Usage: client [options]
options:
  -e service endpoint or socket path, default localhost: 4040
  -n network protocol name [tcp, unix], default tcp

Once started a prompt is provided to interact with service
func main() {
   // setup flags
   var addr string
   var network string
   flag.StringVar(&addr, "e", "localhost:4040", "service endpoint")
   flag.StringVar(&network, "n", "tcp", "network protocol [tcp, unix]")  
   flag.Parse()

   // create a dialer to configure its settings instead of
   // using the default dialer from net.Dial() function.
   dialer := &net.Dialer{
                Timeout: time.Second * 300,
                KeepAlive: time.Minute * 5,
   }

   // simple dialing strategy with retry with a simple backoff.
   // More sophisticated retry strategies follow similar pattern
   // but may include features such as exponential backoff delay, etc.
   var (
      conn            net.Conn
      err             error
      connTries       = 0
      Conn.MaxRetries = 3
      connSleepRetry  = time.Second * 1 
   )
   for connTries < connMaxRetries {
      fmt.Println("creating connection socket to", addr)
      conn, err = dialer.Dial(network, addr)
      if err != nil {
         fmt.Println("failed to create socket:", err)
         switch nerr := err.(type) {
         case net.Error:
            // attempt to retry
            if nerr.Temporary() {
               connTries++
               fmt.Println("trying again in:", connSleepRetry)
               time.Sleep(connSleepRetry)
               continue
            }
            // non-temporary error
            fmt.Println("unable to recover")
            os.Exit(1)
         default: // non networking error
            os.Exit(1)
         }
      }
      // no error break
      break
   }
