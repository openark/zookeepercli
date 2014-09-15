/*
   Copyright 2014 Outbrain Inc.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package main 

import (
	"flag"
	"strings"
	"github.com/outbrain/log"
	"github.com/outbrain/zookeepercli/zk"
)

// main is the application's entry point. It will either spawn a CLI or HTTP itnerfaces.
func main() {
	servers := flag.String("servers", "", "srv1[:port1][,srv2[:port2]...]")
	command := flag.String("c", "", "command (get|ls)")
	//format := flag.String("f", "txt", "output format (txt|json)")
	verbose := flag.Bool("verbose", false, "verbose")
	debug := flag.Bool("debug", false, "debug mode (very verbose)")
	stack := flag.Bool("stack", false, "add stack trace upon error")
	flag.Parse();
	
	log.SetLevel(log.ERROR)
	if *verbose {
		log.SetLevel(log.INFO)
	}
	if *debug {
		log.SetLevel(log.DEBUG)
	}
	if *stack {
		log.SetPrintStackTrace(*stack)
	}

	log.Info("starting")

	if *servers == "" {
		log.Fatal("Expected comma delimited list of servers via --servers")
	}
	serversArray := strings.Split(*servers, ",")
	if len(serversArray) == 0 {
		log.Fatal("Expected comma delimited list of servers via --servers")
	}

	if len(*command) == 0 {
		log.Fatal("Expected command (-c) (get|ls)")
	}
	
	if len(flag.Args()) == 0 {
		log.Fatal("Expected path argument")
	}
	path := flag.Arg(0)
	
	zk.SetServers(serversArray)

	switch *command {
		case "get": {
			result, err := zk.Get(path)
		}
		case "ls": {
		}
		default: log.Fatalf("Unknown command: %s", *command) 
	}
}
