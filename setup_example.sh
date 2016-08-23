#!/bin/sh

mkdir -p tmp/src
mkdir -p tmp/dest

touch -t 201608220000 tmp/src/testfile1
touch -t 201608220000 tmp/src/testfile2
touch -t 201608210000 tmp/src/testfile3
touch -t 201608200000 tmp/src/testfile4

touch -t 201608220000 tmp/src/testfile-20160722.log
touch -t 201608220000 tmp/src/testfile-20160722-01.log
