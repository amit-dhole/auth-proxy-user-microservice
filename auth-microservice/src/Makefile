
TARGETNAME=auth-microservice
TARGETPATH=$(GOPATH)/bin/$(TARGETNAME)
TARGETNAMESP=$(GOPATH)/bin/$(TARGETNAME)

packages = \
	./config \
	./services \
	./utils \


test:
	@$(foreach package,$(packages), \
		set -e; \
		go test -race -coverprofile $(package)/cover.out -covermode=atomic $(package);)

cover:
	echo "mode: count" > cover-all.out
	@$(foreach package,$(packages), \
		tail -n +2 $(package)/cover.out >> cover-all.out;)
	gocover-cobertura < cover-all.out > cover-cobertura.xml

cover-xml:
	@$(foreach package,$(packages), \
		gocov convert $(package)/cover.out | gocov-xml > $(package)/coverage.xml;)

code-quality:
	golangci-lint --exclude-use-default=false --out-format checkstyle run ./... > static-analysis.xml

all: dependencies auth-microservice test cover 

auth-microservice:
	GOOS=linux GOARCH=amd64 go build -o $(GOPATH)/bin/linux_amd64/$(TARGETNAMESP)


local: 
	dependencies localbuild test

dependencies:
	echo "Installing dependencies"
	glide install

show:
	echo "Launching web browser to show overall coverage..."
	go tool cover -html=cover-all.out
