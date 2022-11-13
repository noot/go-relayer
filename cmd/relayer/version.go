package main

import (
	"fmt"
	"runtime/debug"
)

// GetVersion returns our version string for an executable
func getVersion() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "unknown-version"
	}

	commitHash := "???????"
	dirty := ""

	for _, setting := range info.Settings {
		switch setting.Key {
		case "vcs.revision":
			commitHash = setting.Value
		case "vcs.modified":
			if setting.Value == "true" {
				dirty = "-dirty"
			}
		}
	}

	return fmt.Sprintf("%s %.7s%s-%s",
		info.Main.Version, // " (devel)" unless passing a git tagged version to `go install`
		commitHash,        // 7 bytes is what "git rev-parse --short HEAD" returns
		dirty,             // add "-dirty" to commit hash if repo was not clean
		info.GoVersion,
	)
}
