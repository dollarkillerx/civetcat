build:
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w"  agent_.go
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w"  backend_.go
