zookeepercli
============

Simple, lightweight, dependable CLI for ZooKeeper

**zookeepercli** is a non-interactive command line client for [ZooKeeper](http://zookeeper.apache.org/). It provides with:

 * Basic CRUD-like operations: `create`, `set`, `delete`, `get`, `ls` (aka `children`).
 * Well formatted and controlled output: supporting either `txt` or `json` format
 * Single, no-dependencies binary file, based on a native Go ZooKeeper library by [github.com/samuel/go-zookeeper](http://github.com/samuel/go-zookeeper)

### Download & Install

There are [pre built binaries](https://github.com/outbrain/zookeepercli/releases) for download.
You can find `RPM` and `deb` packages, as well as pre-compiled, dependency free `zookeepercli` executable binary.
In fact, the only file installed by the pre-built `RPM` and `deb` packages is said executable binary file. 

Otherwise the source code is freely available; you will need `git` installed as well as `go`, and you're on you own.

  
### Usage:

    $ zookeepercli --help
    Usage of zookeepercli:
      -c="": command (get|ls|create|set|delete)
      -debug=false: debug mode (very verbose)
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
    
 
The only existing solution known to the author provides output in uncontrolled, not-well-formed, inconsistent format, and is relatively heavyweight to invoke.

### License

Release under the [Apache 2.0 license](https://github.com/outbrain/zookeepercli/blob/master/LICENSE)

Authored by [Shlomi Noach](https://github.com/shlomi-noach) at [Outbrain](https://github.com/outbrain)
 
 
 
 

 
