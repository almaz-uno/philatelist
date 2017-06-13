# Go parameters
GOCMD     = go
GOBUILD   = $(GOCMD) build -v
GOCLEAN   = $(GOCMD) clean -v
GOINSTALL = $(GOCMD) install -v
GOTEST    = $(GOCMD) test
GODEP     = $(GOTEST) -i
GOFMT     = gofmt -w
GOGET     = $(GOCMD) get -v
GOLINT    = gometalinter
GOCOV     = gocov

SOURCEDIR = .
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')
BUILDDIR = build

GOLINTERCFG=.gometalinter.json
GOLINTFLAGS=--deadline=300s -t --vendored-linters

VERSION      := v1.0.0
BUILD_TIME   := $(shell date +%FT%T%z)
# this variables should be set by a composer
BUILD_COMMIT :=${BUILD_COMMIT}
BUILD_BRANCH :=${BUILD_BRANCH}

PKGS = bitbucket.org/CuredPlumbum/philatelist
COPY_FILES = 

LDFLAGS	= -ldflags "-X $(PKG)/cmd.buildVersion=${VERSION}  -X $(PKG)/cmd.buildTime=${BUILD_TIME} -X $(PKG)/cmd.buildCommit=${BUILD_COMMIT} -X $(PKG)/cmd.buildBranch=${BUILD_BRANCH}"

VERSION_FILES = duet-proxy/cmd/version.go

.PHONY: clean deps test lint $(GOLINT) distrib install info $(PKGS) $(COPY_FILES)

$(GOLINT):
	$(GOGET) -u github.com/alecthomas/gometalinter
	$(GOLINT) --install --vendored-linters

$(GOCOV):
	$(GOGET) -u github.com/axw/gocov/gocov

lint: deps $(GOLINT)
	$(GOLINT) $(GOLINTFLAGS) --config=$(GOLINTERCFG) ./...

clean:
	go clean
	if [ -d $(BUILDDIR) ] ; then rm -rf $(BUILDDIR) ; fi

test: deps
	$(GOTEST) -race ./...

deps:
	$(GOGET) -t ./...

build: $(SOURCES) deps $(PKGS)
 
$(PKGS):
	$(eval PKG := $@)
	$(eval OUT := $(notdir $@))
	$(GOBUILD) $(LDFLAGS) -o $(BUILDDIR)/$(OUT) $(PKG)

distr: build $(COPY_FILES)

$(COPY_FILES):
	cp $@ $(BUILDDIR)/

info:
	env
	go version
	go env

gocov-report: $(GOCOV) deps
	$(GOCOV) test -race ./... | $(GOCOV) report
