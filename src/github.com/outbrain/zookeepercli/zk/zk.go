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
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

var servers []string

var flags int32 = int32(0)
var acl []zk.ACL = zk.WorldACL(zk.PermAll)

func SetServers(serversArray []string) {
	servers = serversArray
}

func connect() (*zk.Conn, error) {
	conn, _, err := zk.Connect(servers, time.Second)
	return conn, err
}

func Get(path string) ([]byte, error) {
	connection, err := connect()
	if err != nil {
		return []byte{}, err
	}
	defer connection.Close()

	data, _, err := connection.Get(path)
	return data, err
}

func Children(path string) ([]string, error) {
	connection, err := connect()
	if err != nil {
		return []string{}, err
	}
	defer connection.Close()

	children, _, err := connection.Children(path)
	return children, err
}

func Create(path string, data []byte) (string, error) {
	connection, err := connect()
	if err != nil {
		return "", err
	}
	defer connection.Close()

	return connection.Create(path, data, flags, acl)
}

func Set(path string, data []byte) (*zk.Stat, error) {
	connection, err := connect()
	if err != nil {
		return nil, err
	}
	defer connection.Close()

	return connection.Set(path, data, 0)
}

func Delete(path string) error {
	connection, err := connect()
	if err != nil {
		return err
	}
	defer connection.Close()

	return connection.Delete(path, -1)
}
