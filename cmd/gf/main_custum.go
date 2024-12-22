package main

import "github.com/gogf/gf/cmd/gf/v2/internal/cmd"

var (
	appVersion string
	upVersion  string
)

func _setVersion() {
	cmd.SetVersion(appVersion)
	cmd.SetUpVer(upVersion)
}
