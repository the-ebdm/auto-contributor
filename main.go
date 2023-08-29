package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func main() {
	directory := "."
	// Opens an already existing repository.
	r, err := git.PlainOpen(directory)
	CheckIfError(err)

	w, err := r.Worktree()
	CheckIfError(err)

	filename := filepath.Join(directory, "example-git-file")
	err = os.WriteFile(filename, []byte("hello world!"), 0644)
	CheckIfError(err)

	Info("git add example-git-file")
	_, err = w.Add("example-git-file")
	CheckIfError(err)

	Info("git status --porcelain")
	status, err := w.Status()
	CheckIfError(err)
	Info(status.String())

	Info("git commit -m \"example go-git commit\"")
	commit, err := w.Commit("example go-git commit", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "John Doe",
			Email: "john@doe.org",
			When:  time.Now(),
		},
	})

	CheckIfError(err)

	// Prints the current HEAD to verify that all worked well.
	Info("git show -s")
	obj, err := r.CommitObject(commit)
	CheckIfError(err)

	fmt.Println(obj)
}
