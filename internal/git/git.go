package git

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"walkah.dev/walkah/gotem/internal/util"
)

func PullLatest(repoPath string) error {
	fmt.Println("Pulling the latest changes for", repoPath, "...")
	path, err := util.GetAbsolutePath(repoPath)
	if err != nil {
		return err
	}

	r, err := git.PlainOpen(path)
	if err != nil {
		return err
	}

	w, err := r.Worktree()
	if err != nil {
		return err
	}

	options := &git.PullOptions{
		RemoteName: "origin",
	}

	err = w.Pull(options)
	if err != nil {
		if err == git.NoErrAlreadyUpToDate {
			fmt.Println("Already up to date.")
			return nil
		}
		return err
	}

	fmt.Println("Repository has been updated.")
	return nil
}
