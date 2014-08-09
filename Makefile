all: godocbook
	./godocbook net/http github.com/nathankerr/pdf

godocbook: templates.rice-box.go *.go
	go fmt
	go build

templates.rice-box.go: rice templates/*.*
	./rice embed-go
	$(MAKE) vendor

rice: Godeps/_workspace/src/github.com/GeertJohan/go.rice
	go build github.com/nathankerr/godocbook/Godeps/_workspace/src/github.com/GeertJohan/go.rice/rice

godep: Godeps/_workspace/src/github.com/tools/godep
	go build github.com/nathankerr/godocbook/Godeps/_workspace/src/github.com/tools/godep

.PHONY: vendor
vendor: godep
	./godep save -r
	go fmt

clean:
	rm -fv godocbook rice godep

clean-all: clean
	rm -fv templates.rice-box.go