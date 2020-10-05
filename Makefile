all: build

build:
	go build -o autocomplete cmd/autocomplete/main.go
	go build -o autocompleted cmd/autocompleted/main.go

clean:
	rm -f autocomplete autocompleted

.PHONY: all clean
