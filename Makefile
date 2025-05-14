export GO111MODULE=on

all: build

# Build the binary
build:
	go build -o bin/digispark cmd/digispark/main.go

# Run the binary
run: build
	sudo ./bin/digispark $(ARG)

# Clean up the build artifacts
clean:
	rm -rf bin
