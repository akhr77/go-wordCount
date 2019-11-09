GOCMD=go
GOFMT=$(GOCMD) fmt
GOVET=$(GOCMD) vet
GOLINT=golint
GOTEST=$(GOCMD) test

fmt:
	${GOFMT} ./...
	${GOVET} ./...
	${GOLINT} ./...

check:
	${GOTEST} ./... -v
	${GOTEST} ./... -cover
