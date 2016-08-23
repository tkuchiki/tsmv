# tsmv
tsmv is time sliced mv.

## Installation

Download from https://github.com/tkuchiki/tsmv/releases


## Usage

```
$ ./tsmv --help
usage: tsmv [<flags>] [<filepaths>...]

Flags:
      --help              Show context-sensitive help (also try --help-long and --help-man).
  -t, --target-directory=DIRECTORY
                          move all source arguments into directory
  -f, --format="%Y%m%d"   strftime format
  -c, --create-directory  create target directory
  -r, --recursive         create directories recursively
  -m, --mode=0755         file mode
      --dry-run           enable dry-run mode
  -n, --name              name based slice (format %Y[._/-]?%m[._/-]?%d[._/-]?%H?)
      --version           Show application version.

Args:
  [<filepaths>]  some file paths
```

## Examples

```
$ ./setup_example.sh

$ ls tmp/
dest    src

$ ls -l tmp/src/
total 0
-rw-r--r--  1 tkuchiki  tkuchiki  0  8 22 00:00 testfile-20160722-01.log
-rw-r--r--  1 tkuchiki  tkuchiki  0  8 22 00:00 testfile-20160722.log
-rw-r--r--  1 tkuchiki  tkuchiki  0  8 22 00:00 testfile1
-rw-r--r--  1 tkuchiki  tkuchiki  0  8 22 00:00 testfile2
-rw-r--r--  1 tkuchiki  tkuchiki  0  8 21 00:00 testfile3
-rw-r--r--  1 tkuchiki  tkuchiki  0  8 20 00:00 testfile4
```

### dry-run

```
$ find ./tmp/src/ -type f | ./tsmv -c -t ./tmp/dest/ --dry-run
mkdir ./tmp/dest/20160822
mv ./tmp/src//testfile1 ./tmp/dest/20160822/testfile1
mv ./tmp/src//testfile2 ./tmp/dest/20160822/testfile2
mkdir ./tmp/dest/20160821
mv ./tmp/src//testfile3 ./tmp/dest/20160821/testfile3
mkdir ./tmp/dest/20160820
mv ./tmp/src//testfile4 ./tmp/dest/20160820/testfile4

$ find ./tmp/src/ -type f | xargs ./tsmv -c -t ./tmp/dest/ -f "%Y-%m-%d" --dry-run
mkdir ./tmp/dest/2016-08-22
mv ./tmp/src//testfile1 ./tmp/dest/2016-08-22/testfile1
mv ./tmp/src//testfile2 ./tmp/dest/2016-08-22/testfile2
mkdir ./tmp/dest/2016-08-21
mv ./tmp/src//testfile3 ./tmp/dest/2016-08-21/testfile3
mkdir ./tmp/dest/2016-08-20
mv ./tmp/src//testfile4 ./tmp/dest/2016-08-20/testfile4

$ find ./tmp/src/ -type f | xargs ./tsmv -c -t ./tmp/dest/ -f "%Y%m%d" -n --dry-run
mkdir ./tmp/dest/20160722
mv ./tmp/src//testfile-20160722-01.log ./tmp/dest/20160722/testfile-20160722-01.log
mv ./tmp/src//testfile-20160722.log ./tmp/dest/20160722/testfile-20160722.log

$ find ./tmp/src/ -type f | xargs ./tsmv -c -t ./tmp/dest/ -f "%Y%m%d%H" -n --dry-run
mkdir ./tmp/dest/2016072201
mv ./tmp/src//testfile-20160722-01.log ./tmp/dest/2016072201/testfile-20160722-01.log
mkdir ./tmp/dest/2016072200
mv ./tmp/src//testfile-20160722.log ./tmp/dest/2016072200/testfile-20160722.log
```

### run

```
$ find ./tmp/src/ -type f | xargs ./tsmv -c -t ./tmp/dest/ -f "%Y%m%d%H" -n

$ ls tmp/dest/
2016072200      2016072201

$ ls tmp/dest/20160722*
tmp/dest/2016072200:
testfile-20160722.log

tmp/dest/2016072201:
testfile-20160722-01.log

$ find ./tmp/src/ -type f | ./tsmv -c -t ./tmp/dest/

$ ls ./tmp/dest/
20160820        20160821        20160822

$ ls -l ./tmp/dest/20160820/
total 0
-rw-r--r--  1 tkuchiki  tkuchiki  0  8 20 00:00 testfile4

$ ls -l ./tmp/dest/20160821
total 0
-rw-r--r--  1 tkuchiki  tkuchiki  0  8 21 00:00 testfile3

$ ls -l ./tmp/dest/20160822
total 0
-rw-r--r--  1 tkuchiki  tkuchiki  0  8 22 00:00 testfile1
-rw-r--r--  1 tkuchiki  tkuchiki  0  8 22 00:00 testfile2
```
