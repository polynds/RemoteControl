PACKAGE=github.com/polynds/RemoteControl

.PHONY: clean all

all: RemoteControl-linux-amd64 RemoteControl-linux-arm64 RemoteControl-darwin-amd64 RemoteControl-windows-amd64.exe

RemoteControl-linux-amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o RemoteControl-linux-amd64 ${PACKAGE}

RemoteControl-linux-arm64:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o RemoteControl-linux-arm64 ${PACKAGE}

RemoteControl-darwin-amd64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o RemoteControl-darwin-amd64 ${PACKAGE}

RemoteControl-windows-amd64.exe:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o RemoteControl-windows-amd64.exe ${PACKAGE}

RemoteControl :
	go build -o RemoteControl

clean:
	rm -f RemoteControl-linux-amd64 RemoteControl-linux-arm64 RemoteControl-darwin-amd64 RemoteControl-windows-amd64.exe