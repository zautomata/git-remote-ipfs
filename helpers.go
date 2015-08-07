package main

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/whyrusleeping/ipfs-shell"
	"gopkg.in/errgo.v1"
)

func fetchFullBareRepo(root string) (string, error) {
	// TODO: get host from envvar
	shell := shell.NewShell("localhost:5001")
	tmpPath := filepath.Join("/", os.TempDir(), root)
	_, err := os.Stat(tmpPath)
	switch {
	case os.IsNotExist(err) || err == nil:
		if err := shell.Get(root, tmpPath); err != nil {
			return "", errgo.Notef(err, "shell.Get(%s, %s) failed: %s", root, tmpPath, err)
		}
		return tmpPath, nil
	default:
		return "", errgo.Notef(err, "os.Stat(): unhandled error")
	}
}

func interrupt() error {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	return fmt.Errorf("%s", <-c)
}
