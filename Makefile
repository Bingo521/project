app=my_project
build:
	rm -rf output
	mkdir -p output/bin
	mkdir -p output/conf
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${app}
	mv ./${app} ./output/bin
	cp -r ./conf/conf.yml ./output/conf
scp:
	scp -r ./output zhaobingbing@122.51.215.48:~/

gen_proto:
	mkdir -p proto_gen/login
	protoc --proto_path=proto --go_out=./proto_gen/login ./proto/login.proto
	find ./proto_gen -name '*.pb.go' | xargs sed -i '' 's:,omitempty::g'


gen_gen_class_schedule:
	mkdir -p proto_gen/class_schedule
	protoc --proto_path=proto --go_out=./proto_gen/class_schedule ./proto/class_schedule.proto
	find ./proto_gen -name '*.pb.go' | xargs sed -i '' 's:,omitempty::g'

gen_upload:
	mkdir -p proto_gen/upload
	protoc --proto_path=proto --go_out=./proto_gen/upload ./proto/upload.proto
	find ./proto_gen -name '*.pb.go' | xargs sed -i '' 's:,omitempty::g'

gen_message:
	mkdir -p proto_gen/message
	protoc --proto_path=proto --go_out=./proto_gen/message ./proto/message.proto
	find ./proto_gen -name '*.pb.go' | xargs sed -i '' 's:,omitempty::g'

gen_user_info:
	mkdir -p proto_gen/user_info
	protoc --proto_path=proto --go_out=./proto_gen/user_info ./proto/user_info.proto
	find ./proto_gen -name '*.pb.go' | xargs sed -i '' 's:,omitempty::g'
gen_second_hand:
	mkdir -p proto_gen/second_hand
	protoc --proto_path=proto --go_out=./proto_gen/second_hand ./proto/second_hand.proto
	find ./proto_gen -name '*.pb.go' | xargs sed -i '' 's:,omitempty::g'

gen_digg:
	mkdir -p proto_gen/digg
	protoc --proto_path=proto --go_out=./proto_gen/digg ./proto/digg.proto
	find ./proto_gen -name '*.pb.go' | xargs sed -i '' 's:,omitempty::g'
gen_comment:
	mkdir -p proto_gen/comment
	protoc --proto_path=proto --go_out=./proto_gen/comment ./proto/comment.proto
	find ./proto_gen -name '*.pb.go' | xargs sed -i '' 's:,omitempty::g'




