# mylog_go

[![Go Version](https://img.shields.io/badge/Go-1.27.3-00ADD8?logo=go)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

A lightweight, zero-dependency leveled logger for Go with optional timestamps, bold prefixes, and automatic terminal coloring.

## Features

- Log levels: Debug, Info, Warn, and Error
- Optional timestamp prefix
- Optional bold level labels
- Automatic color-coded output for terminal usage
- No external dependencies

## Installation

```bash
go get github.com/MinatoAka/mylog_go
```

## Quick Start

```go
package main

import (
    "github.com/MinatoAka/mylog_go"
)

func main() {
    mylog.SetLevel(mylog.INFO)
    mylog.SetTimestamp(true)
    mylog.SetBold(true)

    mylog.Debug("debug message: %s", "detail")
    mylog.Info("service started")
    mylog.Warn("disk space is low")
    mylog.Error("request failed")
}
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
