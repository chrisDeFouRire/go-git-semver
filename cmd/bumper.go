package cmd

import "github.com/Masterminds/semver"

type bumper func(v semver.Version) semver.Version
