# jsonlog
## Description
A structured json logger package

[![Godoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/datewu/jsonlog)

## Usage
```golang
package main

import "github.com/datewu/jsonlog"

func main() {
	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)
	logger.PrintInfo("build info", map[string]string{
		"version":   "v0.0.1",
		"buildTime": "2021/06/30",
	})
}
```

