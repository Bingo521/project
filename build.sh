#! /bin/bash
rm -rf output
mkdir -p output/bin
mkdir -p output/conf
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o my_project
mv my_project output/bin
cp -r conf/conf.yml output/conf
