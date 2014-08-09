all: godocbook
	./godocbook net/http github.com/nathankerr/pdf

godocbook: *.go
	go fmt
	go build

templates.rice-box.go: rice templates/*
	./rice embed-go

rice: Godeps/_workspace/src/github.com/GeertJohan/go.rice
	go build github.com/nathankerr/godocbook/Godeps/_workspace/src/github.com/GeertJohan/go.rice/rice

.PHONY: vendor
vendor:
	go get github.com/tools/godep
	godep save -r

clean:
	rm -f godocbook rice