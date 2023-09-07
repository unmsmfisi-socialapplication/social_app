
ifeq ($(OS),Windows_NT)
    uname_S := Windows
else
    uname_S := $(shell uname -s)
endif

ifeq ($(uname_S), Windows)
    target = bin\\main.exe
	main = cmd\\server.go
endif
ifeq ($(uname_S), Linux)
    target = bin/main
	main = cmd/server.go
endif

build:
	go build -o $(target) $(main)

run:
	go run cmd/server.go

live:
	air -c .air.toml --build.cmd "go build -o $(target) $(main)" --build.bin "$(target)"
