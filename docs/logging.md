# Hierarchical Logging Library Package 
Package [logging] implements a hierarchical logging package using [Zap] logging library
that is designed for fast, structured, leveled logging but it does not support 
hierarchical logging by default. The main objectives are:

- Provide a named hierarchical leveled logger.
- Minimize the required search time of finding a logger in the list of loggers.
- Provide ability to change log level during runtime and propagate the change to all of children of a logger. 

## Overview

To implement a named hierarchical leveled logger, we need to store 
the loggers in appropriate date structure that we can utilize the 
hierarchical names in finding and updating of a logger at runtime. 
This package uses [Adaptive Radix Tree] data structure to implement a  hierarchical 
logger using [Zap]. The library provides a 
logger interface that implements
 [Zap] SugaredLogger but the standard Zap Logger functions are also accessible by default. 

## Usage

There are two methods that a user can use to add a logger to a package or go program:

Create a logger using the default configuration as follows: 
   
```bash
log, := logging.GetLogger("controller", "device")
 ``` 

Create a logger using a custom configuration as follows:
    
```bash
cfg := logging.Configuration{}
cfg.SetEncoding("json").
   		SetLevel(logging.WarnLevel).
   		SetOutputPaths([]string{"stdout"}).
   		SetName("controller", "device", "change").
   		SetErrorOutputPaths([]string{"stderr"}).
   		SetECMsgKey("Msg").
   		SetECLevelKey("Level").
   		SetECTimeKey("Ts").
   		SetECTimeEncoder(zc.ISO8601TimeEncoder).
   		SetECEncodeLevel(zc.CapitalLevelEncoder).
   		Build()
log := cfg.GetLogger() 
``` 
  
### Change Log Level at Runtime

There are two different ways to change log level of a logger at runtime:

1- By providing the new log level and the name of a logger to the *SetLevel* function in the logging package, e.g.:

```bash
newLogger := logging.SetLevel(logging.FatalLevel, "controller")
```

2- Using *SetLevel* function of each logger, e.g.

```bash
log.SetLevel(logging.WarnLevel)
```

### Change the Logger Sink

Currently, we support enabling one Kafka sink for a logger. To do so, you can run the following command to enable forwarding 
logs to a Kafka cluster:

```bash
sinkUrl := logging.SinkURL{url.URL{Scheme: "kafka", Host: "127.0.0.1:9092", RawQuery: "topic=test_log_topic&key=test-key"}}
log.EnableSink(sinkURL)
```

And to disable writing to the Kafka cluster,

```bash
log.DisableSink()
```



[logging]: https://github.com/onosproject/onos-lib-go/tree/master/pkg/logging
[Zap]: https://godoc.org/go.uber.org/zap
[Adaptive Radix Tree]: https://github.com/plar/go-adaptive-radix-tree