package main

import (
	"github.com/chrisDeFouRire/go-git-semver/cmd"
)

var (
	version = "dev" // overridden at build time
)

func main() {
	cmd.Execute(version)
}
