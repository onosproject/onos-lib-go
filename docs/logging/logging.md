<!--
SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
SPDX-License-Identifier: Apache-2.0
-->

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
hierarchical names in finding and updating of a logger at runtime. The library provides a logger
interface that implements [Zap] SugaredLogger but the standard Zap Logger functions are also accessible by default. 

## Usage

Create a logger using the default configuration as follows: 
   
```bash
log := logging.GetLogger("controller", "device")
``` 

To configure the logger, create a `logging.yaml` configuration file. The `logging.yaml` file defines `sinks` and `loggers`.

### Sinks

Sinks define where and how logs are written. The logging framework supports several types of sinks:

* `stdout` - writes logs to stdout
* `stderr` - writes logs to stderr
* `file` - writes logs to a file
* `kafka` - writes logs to a Kafka topic

The logging configuration may define multiple sinks. Sinks are defined as a mapping of names to their configurations:

```yaml
sinks:
  stdout:
    type: stdout
    encoding: console
    stdout: {}
  file:
    type: file
    encoding: json
    file:
      path: app.log
```

Each sink may be configured with an `encoding` defining how logs should be formatted when written to the sink. Current
supported encodings include:

* `console` writes logs a columnar format
* `json` writes logs in JSON format

Additionally, each sink type supports custom configuration options, e.g. `path` for `file` sinks or `topic` for `kafka` sinks.

### Loggers

Specific loggers can be configured via the logging configuration file. Loggers are organized in a hierarchy wherein descendants
inherit configurations from their ancestors.

The root logging configuration provides a default configuration for all loggers:

```yaml
loggers:
  root:
    level: info
    output:
      stdout:
        sink: stdout
      file:
        sink: file
        level: debug
```

Specific loggers and their descendants can be configured with additional logger configurations:

```yaml
loggers:
  root:
    level: info
    output:
      stdout:
        sink: stdout
      file:
        sink: file
        level: debug
  foo:
    level: warn
  foo/bar/baz:
    level: debug
    output:
      stdout:
        level: info
```

Loggers define a set of `output` configurations, allowing loggers to write to multiple sinks. Loggers inherit outputs 
from their ancestors and may override fields of inherited output configurations.

### Change Log Level at Runtime

The log level for a specific logger can be changed at runtime:

```bash
log.SetLevel(logging.WarnLevel)
```

When the log level is changed, the level will be propagated to any descendants that have no explicit log level configured.

[logging]: https://github.com/onosproject/onos-lib-go/tree/master/pkg/logging
[Zap]: https://godoc.org/go.uber.org/zap
[Adaptive Radix Tree]: https://github.com/plar/go-adaptive-radix-tree