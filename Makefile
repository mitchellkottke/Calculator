TARGET=calculator

build:
	go build -o ./${TARGET}

clean:
	rm -f ./${TARGET}

test: build
	go test ./...