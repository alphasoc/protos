all: build

protos :=	fqdn/fqdn.proto \
			detection/detection.proto

build:
	protoc --go_out=. --go_opt=paths=source_relative $(protos)

clean:
	find . -name "*.go" | xargs rm
