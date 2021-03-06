package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/coreos/locksmith/lock"
)

var (
	cmdSetMax = &Command{
		Name:    "set-max",
		Summary: "Set the maximum number of lock holders",
		Usage:   "UNIT",
		Description: `Set the maximum number of machines that can be rebooting at a given time. This
can be set at any time and will not affect current holders of the lock.`,
		Run: runSetMax,
	}
)

func runSetMax(args []string) (exit int) {
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "New maximum value must be set.")
		return 1
	}

	elc, err := lock.NewEtcdLockClient(nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error initializing etcd client:", err)
		return 1
	}
	l := lock.New("hi", elc)
	max, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invalid maximum value:", args[0])
		return 1
	}

	sem, old, err := l.SetMax(max)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error setting value:", err)
	}

	fmt.Println("Old-Max:", old)
	fmt.Println("Max:", sem.Max)

	return
}
