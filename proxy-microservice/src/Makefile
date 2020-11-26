
TARGETNAME=proxy-microservice
TARGETPATH=$(GOPATH)/bin/$(TARGETNAME)
TARGETNAMESP=$(GOPATH)/bin/$(TARGETNAME)

packages = \
	./config \
	./services \
	./utils \


test:
	@$(foreach package,$(packages), \
		set -e; \
		go test -coverprofile $(package)/cover.out -covermode=count $(package);)

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

all: dependencies proxy-microservice test cover

proxy-microservice:
	GOOS=linux GOARCH=amd64 go build -o $(GOPATH)/bin/linux_amd64/$(TARGETNAMESP)


local: 
	dependencies localbuild test

dependencies:
	echo "Installing dependencies"
	glide install
	go get -u github.com/jteeuwen/go-bindata/...


