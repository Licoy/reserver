@echo off
SET name=reserver-%1
go build -o ../bin/%name%-win.exe -tags forceposix ../main.go
echo windows platform program compilation completed: %1
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build -o ../bin/%name%-darwin -tags forceposix ../main.go
echo macos platform program compilation completed: %1
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o ../bin/%name%-linux -tags forceposix ../main.go
echo linux platform program compilation completed: %1
pause