all: build

build:
	go build -o autosh cmd/autosh/main.go
	go build -o autocomplete cmd/autocomplete/main.go
	go build -o autocompleted cmd/autocompleted/main.go

clean:
	rm -f autosh autocomplete autocompleted

.PHONY: all clean
