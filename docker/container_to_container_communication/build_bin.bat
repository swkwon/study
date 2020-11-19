@ECHO OFF
SET GOOS=linux
go build -o server server.go
go build -o client client.go
SET GOOS=windows
