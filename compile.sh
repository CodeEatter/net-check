export GOOS=windows
export GOARCH=amd64
go build -o bin/app-windows-amd64.exe main.go

export GOOS=windows
export GOARCH=386
go build -o bin/app-windows-386.exe main.go

export GOOS=linux
export GOARCH=amd64
go build -o bin/app-linux-amd64 main.go

export GOOS=linux
export GOARCH=arm64
go build -o bin/app-linux-arm64 main.go

export GOOS=linux
export GOARCH=386
go build -o bin/app-linux-386 main.go

export GOOS=darwin
export GOARCH=amd64
go build -o bin/app-macos-amd64 main.go

export GOOS=darwin
export GOARCH=arm64
go build -o bin/app-macos-arm64 main.go