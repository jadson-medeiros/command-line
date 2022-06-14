package repo_manager

import (
	"errors"
	"fmt"
	"os"
)

type RepoManager struct {
	repos        []string
	ignoreErrors bool
}

func NewRepoManager(baseDir string,
	repoNames []string,
	ignoreErrors bool) (repoManager *RepoManager, err error) {
	_, err = os.Stat(baseDir)
	if err != nil {
		if os.IsNotExist(err) {
			err = errors.New(fmt.Sprintf("base dir: '%s' doesn't exist", baseDir))
		}
		return
	}

	if baseDir[len(baseDir)-1] != '/' {
		baseDir += "/"
	}

	if len(repoNames) == 0 {
		err = errors.New("repo list can't be empty")
		return
	}

	repoManager = &RepoManager{
		ignoreErrors: ignoreErrors,
	}
	for _, r := range repoNames {
		if r == "" {
			err = errors.New("repo name can't be empty")
			return
		}
		path := baseDir + r
		repoManager.repos = append(repoManager.repos, path)
	}

	return
}
