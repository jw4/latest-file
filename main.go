package main

import (
	"fmt"
	"os"
	"path"
	"sort"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		usage()
	}

	for _, dir := range args {
		handle(dir)
	}
}

func check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
		usage()
	}
}

func onlyfiles(dir string) ([]os.DirEntry, error) {
	all, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var ret []os.DirEntry
	for _, entry := range all {
		if entry.Type().IsRegular() {
			ret = append(ret, entry)
		}
	}

	return ret, nil
}

func handle(dir string) {
	all, err := onlyfiles(dir)
	check(err)

	if len(all) == 0 {
		fmt.Fprintf(os.Stderr, "no files in %s\n", dir)
		return
	}

	sort.Slice(all, func(i, j int) bool {
		lhs, err := all[i].Info()
		check(err)

		rhs, err := all[j].Info()
		check(err)

		return lhs.ModTime().Before(rhs.ModTime())
	})

	fmt.Printf("%s\n", path.Join(dir, all[len(all)-1].Name()))
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s directory [ directory [ ... directory] ]\n", os.Args[0])

	os.Exit(1)
}
