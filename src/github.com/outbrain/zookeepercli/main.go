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
	"github.com/outbrain/golib/log"
	"github.com/outbrain/zookeepercli/output"
	"github.com/outbrain/zookeepercli/zk"
	"io/ioutil"
	"os"
	"strings"
)

// main is the application's entry point.
func main() {
	servers := flag.String("servers", "", "srv1[:port1][,srv2[:port2]...]")
	command := flag.String("c", "", "command, required (exists|get|ls|lsr|create|creater|set|delete)")
	force := flag.Bool("force", false, "force operation")
	format := flag.String("format", "txt", "output format (txt|json)")
	verbose := flag.Bool("verbose", false, "verbose")
	debug := flag.Bool("debug", false, "debug mode (very verbose)")
	stack := flag.Bool("stack", false, "add stack trace upon error")
	flag.Parse()

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
		log.Fatal("Expected command (-c) (exists|get|ls|lsr|create|creater|set|delete)")
	}

	if len(flag.Args()) < 1 {
		log.Fatal("Expected path argument")
	}
	path := flag.Arg(0)
	if *command == "ls" {
	} else if strings.HasSuffix(path, "/") {
		log.Fatal("Path must not end with '/'")
	}

	zk.SetServers(serversArray)

	if *command == "creater" {
		*command = "create"
		*force = true
	}
	switch *command {
	case "exists":
		{
			if exists, err := zk.Exists(path); err == nil && exists {
				output.PrintString([]byte("true"), *format)
			} else {
				log.Fatale(err)
			}
		}
	case "get":
		{
			if result, err := zk.Get(path); err == nil {
				output.PrintString(result, *format)
			} else {
				log.Fatale(err)
			}
		}
	case "ls":
		{
			if result, err := zk.Children(path); err == nil {
				output.PrintStringArray(result, *format)
			} else {
				log.Fatale(err)
			}
		}
	case "lsr":
		{
			if result, err := zk.ChildrenRecursive(path); err == nil {
				output.PrintStringArray(result, *format)
			} else {
				log.Fatale(err)
			}
		}
	case "create":
		{
			if len(flag.Args()) < 2 {
				log.Fatal("Expected data argument")
			}
			if result, err := zk.Create(path, []byte(flag.Arg(1)), *force); err == nil {
				log.Infof("Created %+v", result)
			} else {
				log.Fatale(err)
			}
		}
	case "set":
		{
			var info []byte
			if len(flag.Args()) > 1 {
				info = []byte(flag.Arg(1))
			} else {
				var err error
				info, err = ioutil.ReadAll(os.Stdin)
				if err != nil {
					log.Fatale(err)
				}
			}
			if result, err := zk.Set(path, info); err == nil {
				log.Infof("Set %+v", result)
			} else {
				log.Fatale(err)
			}
		}
	case "delete":
		{
			if err := zk.Delete(path); err != nil {
				log.Fatale(err)
			}
		}
	default:
		log.Fatalf("Unknown command: %s", *command)
	}
}
