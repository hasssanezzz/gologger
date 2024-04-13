# gologger

Gologger is a lightweight and customizable logger package for Go (Golang) applications. It provides functionalities to log messages with different log levels, timestamp formatting, and customizable prefixes.

## Features
* Logging with different log levels: INFO, DEBUG, WARNING, CRITICAL, FATAL.
* Customizable log message prefix.
* Support for writing log messages to different outputs (e.g., stdout, files).

## Installation

To use gologger in your Go project, you can install it using the following go get command:

```bash
$ go get github.com/hasssanezzz/gologger
```

## Usage


1. Import gologger into your Go code:
```go
import "github.com/hasssanezzz/gologger"
```

2. Initialize a new logger instance:
```go
logger := gologger.New(os.Stdout, "[MYAPP]:")
```

3. Use the logger to log messages:
```go
logger.Info("This is an informational message.")
logger.Warning("This is a warning message.")
logger.Error("This is an error message.")
logger.Critical("This is a critical message.")
logger.Fatal("This is a fatal message.")
```

### Configuration Options
You can customize the logger's behavior using the following configuration options:

* `Out`: Output destination (io.Writer) for log messages.
* `Prefix`: Custom prefix to prepend to log messages.

### Example:

```go
package main

import (
  "fmt"
  "os"

  "github.com/hasssanezzz/gologger/gologger"
)

func main() {
  // Create a file to log into
  file, err := os.OpenFile("logfile.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
  if err != nil {
    fmt.Println("Failed to open log file")
  }
  defer file.Close()

  logger := gologger.New(file, "[MYAPP]:")

  logger.Info("Hello world")
}
```

### Logging Levels

The logger supports the following log levels:

* **Info**: Informational messages.
* **Debug**: Debugging messages (not enabled by default).
* **Warning**: Warning messages.
* **Critical**: Critical error messages.
* **Fatal**: Fatal error messages (exits the application after logging).