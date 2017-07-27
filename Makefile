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
DISTRDIR = $(BUILDDIR)/distr

GOLINTERCFG=.gometalinter.json
GOLINTFLAGS=--deadline=300s -t --vendored-linters
CHECKSTYLE_FILE=$(BUILDDIR)/gometalinter/checkstyle-result.xml
GOLINTER_SUPPRESS_ERR=

GOCOV_COVER_XML=$(BUILDDIR)/gocov/coverage.xml

VERSION      := v1.0.0
BUILD_TIME   := $(shell date +%FT%T%z)
# this variables should be set by a composer
BUILD_COMMIT :=${BUILD_COMMIT}
BUILD_BRANCH :=${BUILD_BRANCH}

PKGS = bitbucket.org/CuredPlumbum/philatelist
COPY_FILES = geoimage.yaml geoimage.swagger.yaml

LDFLAGS	= -ldflags "-X $(PKG)/cmd.buildVersion=${VERSION}  -X $(PKG)/cmd.buildTime=${BUILD_TIME} -X $(PKG)/cmd.buildCommit=${BUILD_COMMIT} -X $(PKG)/cmd.buildBranch=${BUILD_BRANCH}"

.PHONY: clean deps test lint $(GOLINT) distrib install info $(PKGS) $(COPY_FILES)

$(GOLINT):
	$(GOGET) -u github.com/alecthomas/gometalinter
	$(GOLINT) --install --vendored-linters

$(GOCOV):
	$(GOGET) -u github.com/axw/gocov/gocov
	$(GOGET) -u github.com/axw/gocov/...
	$(GOGET) -u github.com/AlekSi/gocov-xml

lint: deps $(GOLINT)
	$(GOLINT) $(GOLINTFLAGS) --config=$(GOLINTERCFG) ./...;$(GOLINTER_SUPPRESS_ERR)

lint-checkstyle: deps $(GOLINT)
	mkdir -p $(dir $(CHECKSTYLE_FILE))
	$(GOLINT) $(GOLINTFLAGS) --checkstyle --config=$(GOLINTERCFG) ./... > $(CHECKSTYLE_FILE);$(GOLINTER_SUPPRESS_ERR)

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
	$(GOBUILD) $(LDFLAGS) -o $(DISTRDIR)/$(OUT) $(PKG)

distr: build $(COPY_FILES)

$(COPY_FILES):
	mkdir -p $(DISTRDIR)
	cp $@ $(DISTRDIR)/

info:
	env
	go version
	go env

gocov-report: $(GOCOV) deps
	$(GOCOV) test -race ./... | $(GOCOV) report

gocov-cover-xml: $(GOCOV) deps
	mkdir -p $(dir $(GOCOV_COVER_XML))
	$(GOCOV) test -race ./... | gocov-xml > $(GOCOV_COVER_XML)
