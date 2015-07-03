zookeepercli
============

Simple, lightweight, dependable CLI for ZooKeeper

**zookeepercli** is a non-interactive command line client for [ZooKeeper](http://zookeeper.apache.org/). It provides with:

 * Basic CRUD-like operations: `create`, `set`, `delete` (aka `rm`), `exists`, `get`, `ls` (aka `children`).
 * Extended operations: `lsr` (ls recursive), `creater` (create recursively), `deleter` (aka `rmr`, delete recursively)
 * Well formatted and controlled output: supporting either `txt` or `json` format
 * Single, no-dependencies binary file, based on a native Go ZooKeeper library 
   by [github.com/samuel/go-zookeeper](http://github.com/samuel/go-zookeeper) ([LICENSE](https://github.com/outbrain/zookeepercli/blob/master/go-zookeeper-LICENSE))

### Download & Install

There are [pre built binaries](https://github.com/outbrain/zookeepercli/releases) for download.
You can find `RPM` and `deb` packages, as well as pre-compiled, dependency free `zookeepercli` executable binary.
In fact, the only file installed by the pre-built `RPM` and `deb` packages is said executable binary file. 

Otherwise the source code is freely available; you will need `git` installed as well as `go`, and you're on your own.

  
### Usage:

    $ zookeepercli --help
    Usage of zookeepercli:
      -acls="31": optional, csv list [1|,2|,4|,8|,16|,31]
      -auth_pwd="": optional, digest scheme, pwd
      -auth_usr="": optional, digest scheme, user
      -c="": command (exists|get|ls|lsr|create|creater|set|delete|rm|deleter|rmr|getacl|setacl)
      -debug=false: debug mode (very verbose)
      -force=false: force operation
      -format="txt": output format (txt|json)
      -servers="": srv1[:port1][,srv2[:port2]...]
      -stack=false: add stack trace upon error
      -verbose=false: verbose
    

### Examples:
    
    
    $ zookeepercli --servers srv-1,srv-2,srv-3 -c create /demo_only some_value
    
    # Default port is 2181. The above is equivalent to:
    $ zookeepercli --servers srv-1:2181,srv-2:2181,srv-3:2181 -c create /demo_only some_value
    
    $ zookeepercli --servers srv-1,srv-2,srv-3 --format=txt -c get /demo_only
    some_value
    
    # Same as above, JSON format output:
    $ zookeepercli --servers srv-1,srv-2,srv-3 --format=json -c get /demo_only
    "some_value"
    
    # exists exits with exit code 0 when path exists, 1 when path does not exist 
    $ zookeepercli --servers srv-1,srv-2,srv-3 -c exists /demo_only
    true
    
    $ zookeepercli --servers srv-1,srv-2,srv-3 -c set /demo_only another_value
    
    $ zookeepercli --servers srv-1,srv-2,srv-3 --format=json -c get /demo_only
    "another_value"
    
    $ zookeepercli --servers srv-1,srv-2,srv-3 -c delete /demo_only
    
    $ zookeepercli --servers srv-1,srv-2,srv-3 -c get /demo_only
    2014-09-15 04:07:16 FATAL zk: node does not exist
    
    $ zookeepercli --servers srv-1,srv-2,srv-3 -c create /demo_only "path placeholder"
    $ zookeepercli --servers srv-1,srv-2,srv-3 -c create /demo_only/key1 "value1"
    $ zookeepercli --servers srv-1,srv-2,srv-3 -c create /demo_only/key2 "value2"
    $ zookeepercli --servers srv-1,srv-2,srv-3 -c create /demo_only/key3 "value3"
    
    $ zookeepercli --servers srv-1,srv-2,srv-3 -c ls /demo_only
    key3
    key2
    key1
    
    # Same as above, JSON format output:
    $ zookeepercli --servers srv-1,srv-2,srv-3 --format=json -c ls /demo_only
    ["key3","key2","key1"]
    
    $ zookeepercli --servers srv-1,srv-2,srv-3 -c delete /demo_only
    2014-09-15 08:26:31 FATAL zk: node has children
    
    $ zookeepercli --servers srv-1,srv-2,srv-3 -c delete /demo_only/key1
    $ zookeepercli --servers srv-1,srv-2,srv-3 -c delete /demo_only/key2
    $ zookeepercli --servers srv-1,srv-2,srv-3 -c delete /demo_only/key3
    $ zookeepercli --servers srv-1,srv-2,srv-3 -c delete /demo_only

    # /demo_only path now does not exist.
    
    # Create recursively a path:
    $ zookeepercli --servers=srv-1,srv-2,srv-3 -c creater "/demo_only/child/key1" "val1"
    $ zookeepercli --servers=srv-1,srv-2,srv-3 -c creater "/demo_only/child/key2" "val2"
    
    # "-c creater" is same as "-c create --force"

    $ zookeepercli --servers=srv-1,srv-2,srv-3 -c get "/demo_only/child/key1"
    val1

    # This path was auto generated due to recursive create:
    $ zookeepercli --servers=srv-1,srv-2,srv-3 -c get "/demo_only" 
    zookeepercli auto-generated
    
    # ls recursively a path and all sub children:
    $ zookeepercli --servers=srv-1,srv-2,srv-3 -c lsr "/demo_only" 
    child
    child/key1
    child/key2

    # set value with read and write acl using digest authentication
    $ zookeepercli --servers 192.168.59.103 --auth_usr "someuser" --auth_pwd "pass" --acls 1,2 -c create /secret4 value4
    
    # get value using digest authentication
    $ zookeepercli --servers 192.168.59.103 --auth_usr "someuser" --auth_pwd "pass" -c get /secret4

    # create a value with custom acls
    $ zookeepercli --servers 192.168.59.103 -c create /secret5 value5 world:anyone:rw,digest:someuser:hashedpw:crdwa

    # view the current acl on a path
    $ zookeepercli --servers srv-1,srv-2,srv-3 -c create /demo_acl "some value"
    $ zookeepercli --servers srv-1,srv-2,srv-3 -c getacl /demo_acl
    world:anyone:cdrwa

    # set an acl with world and digest authentication
    $ zookeepercli --servers srv-1,srv-2,srv-3 -c setacl /demo_acl "world:anyone:rw,digest:someuser:hashedpw:crdwa"
    $ zookeepercli --servers srv-1,srv-2,srv-3 -c getacl /demo_acl
    world:anyone:rw
    digest:someuser:hashedpw:cdrwa

    # set an acl with world and digest authentication creating the node if it doesn't exist
    $ zookeepercli --servers srv-1,srv-2,srv-3 -force -c setacl /demo_acl_create "world:anyone:rw,digest:someuser:hashedpw:crdwa"

The tool was built in order to allow with shell scripting seamless integration with ZooKeeper. 
There is another, official command line tool for ZooKeeper that the author found inadequate 
in terms of output format and output control, as well as large footprint. 
**zookeepercli** overcomes those limitations and provides with quick, well formatted output as well as
enhanced functionality. 

### Docker

You can also build and run **zookeepercli** in a Docker container. To build the image:

    $ docker build -t zookeepercli .

Now, you can run **zookeepercli** from a container. Examples:

    $ docker run --rm -it zookeepercli --servers $ZK_SERVERS -c create /docker_demo "test value"
    $ docker run --rm -it zookeepercli --servers $ZK_SERVERS -c get /docker_demo
    test value
    $ docker run --rm -it zookeepercli --servers $ZK_SERVERS -c ls /
    docker_demo
    zookeeper

### License

Release under the [Apache 2.0 license](https://github.com/outbrain/zookeepercli/blob/master/LICENSE)

Authored by [Shlomi Noach](https://github.com/shlomi-noach) at [Outbrain](https://github.com/outbrain)
 
 
 
 

 
