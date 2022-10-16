generate_go_pb:
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb  

clean_pb:
	rm pb/*.go

build:
	go run main.go