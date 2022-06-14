package helpers

import (
	"os"
	"os/exec"
	"path"
)

func ConfigureGit() (err error) {
	err = exec.Command("git", "config", "--global", "user.name", "YOUR USER NAME").Run()
	if err != nil {
		return
	}
	err = exec.Command("git", "config", "--global", "user.email", "YOUR EMAIL").Run()
	return
}

func CreateDir(baseDir string, name string, initGit bool) (err error) {
	dirName := path.Join(baseDir, name)
	err = os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		return
	}

	if !initGit {
		return
	}

	currDir, err := os.Getwd()
	if err != nil {
		return
	}
	defer os.Chdir(currDir)
	os.Chdir(dirName)
	err = exec.Command("git", "init").Run()
	return
}
