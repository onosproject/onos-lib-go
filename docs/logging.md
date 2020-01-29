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
hierarchical names in finding and updating a logger during runtime. 
This package uses [Adaptive Radix Tree] data structure to implement a  hierarchical 
logger using [Zap]. The library provides a 
logger interface that implements
 [Zap] SugaredLogger but the standard Zap Logger functions are also accessible by default. 

## Usage

There are two methods that a user can use to add a logger to a package or go program:

1. Create a logger using the default configuration by providing just a name and a level for the logger to *AddLogger* function.
For example:

```bash
logging.AddLogger("warn", "controller")  // 
logging.AddLogger("info", "controller", "device")
logging.AddLogger("error", "controller", "device", "change")
```

> if you don't provide the logger level, the logger inherits from its parent
> and if none of its parents exists then it inherits from root logger (i.e. default logger).

2. 












[logging]: https://github.com/onosproject/onos-lib-go/tree/master/pkg/logging
[Zap]: https://godoc.org/go.uber.org/zap
[Adaptive Radix Tree]: https://github.com/plar/go-adaptive-radix-tree