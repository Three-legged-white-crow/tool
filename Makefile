linux64:
	GOOS=linux GOARCH=amd64 go build -o pc main.go

windows64:
	GOOS=windows GOARCH=amd64 go build -o pc.exe main.go