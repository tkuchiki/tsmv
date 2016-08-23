package main

import (
	"bufio"
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"runtime"
	"strings"
	"time"
)

func usage() {
	kingpin.Usage()
	os.Exit(1)
}

type Rename struct {
	srcPath   string
	destPath  string
	format    string
	mode      os.FileMode
	createDir bool
	recursive bool
	dryRun    bool
	name      bool
}

func (r *Rename) do() error {
	var finfo os.FileInfo
	var err error

	finfo, err = getFileInfo(r.srcPath)
	if err != nil {
		return err
	}

	if finfo.IsDir() {
		return nil
	}

	if !isExist(r.srcPath) {
		return fmt.Errorf("%s: no such file or directory", r.srcPath)
	}

	var destDir string
	var t time.Time

	if *name {
		timestr := extractTime(r.srcPath)
		if timestr == "" {
			return nil
		}

		t, err = timeParse(timestr)
		if err != nil {
			return err
		}
	} else {
		t = finfo.ModTime()
	}
	destDir = fmt.Sprintf(pathFormat(psep), strings.TrimRight(r.destPath, psep), strftime(r.format, t))

	if r.createDir {
		if !isExist(destDir) {
			if r.dryRun {
				if v, ok := isCreated[destDir]; !(ok && v) {
					if r.recursive {
						fmt.Println("mkdir", "-p", destDir)
					} else {
						fmt.Println("mkdir", destDir)
					}
				}
				isCreated[destDir] = true
			} else {
				err = mkdir(destDir, os.FileMode(r.mode), r.recursive)
			}
			if err != nil {
				return err
			}
		}
	}

	if r.dryRun {
		fmt.Println("mv", r.srcPath, createDestPath(r.srcPath, destDir, psep))
	} else {
		err = rename(r.srcPath, destDir, psep)
	}
	if err != nil {
		return err
	}

	return err
}

var (
	filePaths = kingpin.Arg("filepaths", "some file paths").Strings()

	targetDir = kingpin.Flag("target-directory", "move all source arguments into directory").Short('t').PlaceHolder("DIRECTORY").String()
	format    = kingpin.Flag("format", "strftime format").Short('f').Default("%Y%m%d").String()
	createDir = kingpin.Flag("create-directory", "create target directory").Short('c').Bool()
	recursive = kingpin.Flag("recursive", "create directories recursively").Short('r').Bool()
	mode      = kingpin.Flag("mode", "file mode").Short('m').Default("0755").Int64()
	dryRun    = kingpin.Flag("dry-run", "enable dry-run mode").Bool()
	name      = kingpin.Flag("name", "name based slice (format %Y[._/-]?%m[._/-]?%d[._/-]?%H?)").Short('n').Bool()

	isCreated = make(map[string]bool)
	psep      = "/"
)

func main() {
	kingpin.Version("0.1.1")
	kingpin.Parse()

	if runtime.GOOS == "windows" {
		psep = "\\"
	}

	var err error
	srcPaths := make([]string, 0)
	var destPath string

	if *targetDir == "" {
		if len(*filePaths) <= 1 {
			usage()
		}

		srcPaths, destPath = getSrcDestPaths(*filePaths)
	} else {
		var stdinInfo os.FileInfo
		stdinInfo, err = os.Stdin.Stat()
		if err != nil {
			errorln(err)
		}

		if stdinInfo.Mode()&os.ModeNamedPipe != 0 {
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			line := scanner.Text()
			if strings.Index(line, "\x00") > -1 {
				for _, p := range strings.Split(strings.TrimRight(line, "\x00"), "\x00") {
					srcPaths = append(srcPaths, p)
				}
			} else {
				srcPaths = append(srcPaths, line)
			}
			for scanner.Scan() {
				srcPaths = append(srcPaths, scanner.Text())
			}
		} else {
			srcPaths = *filePaths
		}

		destPath = *targetDir
	}

	for _, srcPath := range srcPaths {
		if srcPath == "" {
			continue
		}

		r := Rename{
			srcPath:   srcPath,
			destPath:  destPath,
			format:    *format,
			mode:      os.FileMode(*mode),
			createDir: *createDir,
			recursive: *recursive,
			dryRun:    *dryRun,
			name:      *name,
		}

		err = r.do()

		if err != nil {
			errorln(err)
		}
	}
}
