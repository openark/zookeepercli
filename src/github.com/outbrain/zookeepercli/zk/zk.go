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

package zk

import (
	"time"
	"github.com/samuel/go-zookeeper/zk"
)


var servers []string

func SetServers(serversArray []string) {
	servers = serversArray
}


func connect() (*zk.Conn, error) {
	conn, _, err := zk.Connect(servers, time.Second)
	return conn, err
}


func Get(path string) ([]byte, error) {
	connection, err := connect()
	if err != nil { return []byte{}, err }
	defer connection.Close()
	
	data, _, err := connection.Get(path)
	return data, err
}
