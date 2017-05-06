PATH_API_MAIN = main.go

build:
	go build -o bin/hostos/api $(PATH_API_MAIN)

build_linux:
	GOOS=linux GOARCH=amd64 go build -o bin/linux/api $(PATH_API_MAIN)
