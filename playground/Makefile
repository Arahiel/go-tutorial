help:
	echo "Options:"
	echo "build - builds entire application"
	echo "run - run application"
	echo "all - build and run application"

build:
	go build -o bin/

runStandalone:
	bin/playground.exe

run:
	go run main.go

all:
	make build runStandalone