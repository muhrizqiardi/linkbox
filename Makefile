build:
	node build.cjs
	GOOS=linux GOARCH=amd64 go build -o bin/linux-amd64/linkbox-linux-amd64 cmd/main.go
	GOOS=windows GOARCH=amd64 go build -o bin/windows-amd64/linkbox-windows-amd64.exe cmd/main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/darwin-amd64/linkbox-darwin-amd64 cmd/main.go
	GOOS=linux GOARCH=arm64 go build -o bin/linux-arm64/linkbox-linux-arm64 cmd/main.go
	GOOS=windows GOARCH=arm64 go build -o bin/windows-arm64/linkbox-windows-arm64.exe cmd/main.go
	GOOS=darwin GOARCH=arm64 go build -o bin/darwin-arm64/linkbox-darwin-arm64 cmd/main.go

