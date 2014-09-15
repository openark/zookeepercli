zookeepercli
============

Simple, lightweight, dependable CLI for ZooKeeper

**zookeepercli** is a non-interactive command line client for [ZooKeeper](http://zookeeper.apache.org/). It provides with:

 * Basic CRUD-like operations: `create`, `set`, `delete`, `get`, `ls` (aka `children`).
 * Well formatted and controlled output: supporting either `txt` or `json` format
 * Single, no-dependencies binary file, based on a native Go ZooKeeper library by [github.com/samuel/go-zookeeper](http://github.com/samuel/go-zookeeper)

Usage samples:

    
    
    $ zookeepercli --servers srv-1,srv-2,srv-3 -c create /demo_only some_value
    
    # Default port is 2181. The above is equivalent to:
    $ zookeepercli --servers srv-1:2181,srv-2:2181,srv-3:2181 -c create /demo_only some_value
    
    $ zookeepercli --servers srv-1,srv-2,srv-3 --format=txt -c get /demo_only
    some_value
    
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
    
    
    $ zookeepercli --servers srv-1,srv-2,srv-3 --format=json -c ls /demo_only
    ["key3","key2","key1"]
    
    
    $ zookeepercli --servers srv-1,srv-2,srv-3 -c delete /demo_only
    2014-09-15 08:26:31 FATAL zk: node has children
    
    $ zookeepercli --servers srv-1,srv-2,srv-3 -c delete /demo_only/key1
    $ zookeepercli --servers srv-1,srv-2,srv-3 -c delete /demo_only/key2
    $ zookeepercli --servers srv-1,srv-2,srv-3 -c delete /demo_only/key3
    $ zookeepercli --servers srv-1,srv-2,srv-3 -c delete /demo_only
    

 
The only existing solution known to the author provides output in uncontrolled, not-well-formed, inconsistent format, and is relatively heavyweight to invoke.

 
 
 
 

 
