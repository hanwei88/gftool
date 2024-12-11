package main

import "github.com/gogf/gf/cmd/gf/v2/internal/cmd"

func init() {

	var (
		appVersion string
		upVersion  string
	)

	cmd.SetVersion(appVersion)
	cmd.SetUpVer(upVersion)
}
