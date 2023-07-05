# app directories

[![MIT License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/Zarthus/iogo.svg)](https://pkg.go.dev/github.com/Zarthus/appdir)

Small helper library to help you locate useful directories based on the following resources:

- https://stackoverflow.com/a/41598938
- https://stackoverflow.com/a/39868933
- https://apple.stackexchange.com/questions/28928/what-is-the-macos-equivalent-to-windows-appdata-folder

## Add it to your project

```bash
go get github.com/zarthus/appdir
```

## Description

API:

```js
GetApplicationDirectory(); // not for linux (presently)
GetUserConfigDirectory();
```

Example:

```go
package main

import (
	"github.com/zarthus/appdir"
	"path/filepath"
	"os"
)

func main() {
	dir, _ := appdir.GetUserConfigDirectory() // TODO error handling
	apppath := filepath.Join(dir, "myapp")
	_ = os.MkdirAll(apppath, 0750)
	_ = os.WriteFile(filepath.Join(apppath, "config.json"), []byte("{\"foo\": true}"), 0644)
	
	// on linux, this would for example create `~/.config/myapp/config.json` with `{"foo": true}`
}
```
