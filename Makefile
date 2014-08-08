all: godocbook
	./godocbook net/http github.com/nathankerr/pdf

godocbook: *.go
	go fmt
	go build

templates.rice-box.go: templates/*
	rice embed-go