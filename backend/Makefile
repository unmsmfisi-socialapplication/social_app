
ifeq ($(OS),Windows_NT)
    uname_S := Windows
else
    uname_S := $(shell uname -s)
endif

ifeq ($(uname_S), Windows)
    target = main.exe
endif
ifeq ($(uname_S), Linux)
    target = main
endif

build:
	go build -o bin/$(target) cmd/server.go

run:
	go run cmd/server.go
