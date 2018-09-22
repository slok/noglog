# noglog

noglog is a replacement for go [glog] implementation.

## Features

- Replace glog messages with your own logger (implement small interface and done).
- No [flags] modification.

## Why

Glog is a logger used on some important projects (Kubernetes), Glog modifies globals like `flag` of your program and it can't be done nothing to disable, modify...

For example in Kubernetes glog is around all the source code, they don't use a common app logger interface that it can be reimplement with any log implementation, or pass any logger as dependency injection, so the Kubernetes extensions that use [client-go]

## When to use noglog

- Tired of glog logs.
- Every extension/app of Kubernetes I do has anoying logs.
- My app flags have glog flags that I didn't ask for.
- `ERROR: logging before flag.Parse: W0922` WTF.

This package is a "bring you own" logger to glog. It will replace all the glog calls with any logger that implements the noglog `Logger` interface.

## Start

For a complete example, check [this][example]

First we need to add noglog as dependency, but it will need to be placed on `github.com/golang/glog` import path, so `noglog` code replaces `glog`

### go modules

```text
require (
    github.com/google/glog master
)

replace (
    github.com/google/glog => github.com/slok/noglog master
)
```

### dep

```toml
[[override]]
  name = "github.com/golang/glog"
  source = "github.com/slok/noglog"
  branch = "master"
```

Second, import glog and the first thing that your app should do is instantiate your logger and set as the glog logger. Example:

```golang
import (
    "github.com/golang/glog"
)

func main() {
    // Import custom logger that satisfies noglog Logger interface.
    logger := mylogger.New()

    glog.SetLogger(logger)
}
```

For security reasons if you don't set a logger, it will use a dummy logger that will disable Kubernetes logs. Sometimes this is useful also.

## Alternatives

Although noglog can use any logger due to the simple interface, there are other alternatives that replace the logger for a specific logger with defaults, if you don't want to bring your own or customize these loggers here are the alternatives:

- [glog-logrus]
- [istio-glog](zap)

[glog]: github.com/golang/glog
[client-go]: https://github.com/kubernetes/client-go
[flags]: https://golang.org/pkg/flag
[glog-logrus]: https://github.com/kubermatic/glog-logrus
[istio-glog]: https://github.com/istio/glog
[example]: /examples
