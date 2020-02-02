app=my_project
build:
	rm -rf output
	mkdir -p output/bin
	mkdir -p output/conf
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${app}
	mv ./${app} ./output/bin
	cp -r ./conf ./output
scp:
	scp -r ./output zhaobingbing@122.51.215.48:~/

gen_login:
	mkdir -p proto_gen/login
	protoc --proto_path=proto --go_out=./proto_gen/login ./proto/login.proto
