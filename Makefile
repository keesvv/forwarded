ENTRY=cmd/forwarded/main.go
TARGET=build/forwarded

all: build strip

run:
	go run $(ENTRY)

build:
	go build -ldflags="-s -w" -o $(TARGET) $(ENTRY)

strip:
	strip $(TARGET)
