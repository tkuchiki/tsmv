package main

import (
	"fmt"
	gostrftime "github.com/jehiah/go-strftime"
	"os"
	"path"
	"strings"
	"time"
)

func isExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func errorln(msg interface{}) {
	fmt.Println(msg)
	os.Exit(1)
}

func errorf(format string, msgs ...interface{}) {
	fmt.Printf(format, msgs...)
	fmt.Println()
	os.Exit(1)
}

func getFileInfo(filename string) (os.FileInfo, error) {
	var finfo os.FileInfo

	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return finfo, err
	}

	finfo, err = f.Stat()

	return finfo, err
}

func mkdir(dirname string, perm os.FileMode, recursive bool) error {
	var err error
	if recursive {
		err = os.MkdirAll(dirname, perm)
	} else {
		err = os.Mkdir(dirname, perm)
	}

	return err
}

func strftime(format string, t time.Time) string {
	t.In(time.Local)
	return gostrftime.Format(format, t)
}

func getSrcDestPaths(paths []string) ([]string, string) {
	n := len(paths)

	return paths[:n-1], paths[n-1]
}

func pathFormat(psep string) string {
	return fmt.Sprintf("%%s%s%%s", psep)
}

func createDestPath(srcPath, destDir, psep string) string {
	filename := path.Base(srcPath)
	return fmt.Sprintf(pathFormat(psep), strings.TrimRight(destDir, psep), filename)
}

func rename(srcPath, destDir, psep string) error {
	return os.Rename(srcPath, createDestPath(srcPath, destDir, psep))
}
