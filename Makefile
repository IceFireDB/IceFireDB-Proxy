PROG=bin/IceFireDB-Proxy


SRCS=./cmd/IceFireDB-Proxy


# git commit hash
COMMIT_HASH=$(shell git rev-parse --short HEAD || echo "GitNotFound")

# Compile the date
BUILD_DATE=$(shell date '+%Y-%m-%d %H:%M:%S')
# compile
CFLAGS = -ldflags "-s -w -X \"main.BuildVersion=${COMMIT_HASH}\" -X \"main.BuildDate=$(BUILD_DATE)\""

all:
	if [ ! -d "./bin/" ]; then \
	mkdir bin; \
	fi
	GOPROXY=$(GOPROXY) go build $(CFLAGS) -o $(PROG) $(SRCS)

# Compiling the RACE version
race:
	if [ ! -d "./bin/" ]; then \
    	mkdir bin; \
    	fi
	GOPROXY=$(GOPROXY) go build $(CFLAGS) -race -o $(PROG) $(SRCS)

clean:
	rm -rf ./bin

run:
	GOPROXY=$(GOPROXY) go run --race cmd/proxy/main.go -c config/config.yaml
