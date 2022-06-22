package main

import (
	"fmt"
	"log"

	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"

	"github.com/jadson-medeiros/command-line/cmd"
)

var (
	version = "v1.0.0"
)

var (
	gitTag         = ""
	buildTimestamp = ""
)

func main() {
	if gitTag != "" {
		version = gitTag
		fmt.Println("Git tag:", gitTag)
	}

	if buildTimestamp != "" {
		fmt.Println("Built at:", buildTimestamp)
	}

	fmt.Println("Current version is: ", version)

	v := semver.MustParse(version[1:])
	latest, err := selfupdate.UpdateSelf(v, "jadson-medeiros/command-line")
	if err != nil {
		log.Fatalf("Binary update failed: %v", err)
		return
	} else {
		if latest.Version.String() != version {
			fmt.Printf("Updated version to: v%v\n", latest.Version)
		}
	}

	cmd.Execute()
}
