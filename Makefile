appname := yago
BINARYNAME="YaGo"
VERSION=v0.1.0
TARGET=all
BUILD_TIME=$(shell date +%FT%T%z)
BUILD=$(shell git rev-parse HEAD)
LDFLAGS=-ldflags="\
	-s \
	-w \
	-X main.Name=${BINARYNAME} \
	-X main.Version=${VERSION} \
	-X main.BuildID=${BUILD} \
	-X main.BuildDate=${BUILD_TIME}"

sources := $(wildcard *.go)

cmd = GOOS=$(1) GOARCH=$(2) go build ${LDFLAGS} -o build/$(appname)$(3)
tar = cd build && tar -cvzf $(appname)_$(1)_$(2).tar.gz $(appname)$(3) && rm $(appname)$(3)
zip = cd build && zip $(appname)_$(1)_$(2).zip $(appname)$(3) && rm $(appname)$(3)

.PHONY: all windows darwin linux dev clean

all: windows darwin linux

clean:
	rm -rf build/

##### LINUX BUILDS #####
linux: build/linux_arm.tar.gz build/linux_arm64.tar.gz build/linux_386.tar.gz build/linux_amd64.tar.gz

build/linux_386.tar.gz: $(sources)
	$(call cmd,linux,386,)
	$(call tar,linux,386)

build/linux_amd64.tar.gz: $(sources)
	$(call cmd,linux,amd64,)
	$(call tar,linux,amd64)

build/linux_arm.tar.gz: $(sources)
	$(call cmd,linux,arm,)
	$(call tar,linux,arm)

build/linux_arm64.tar.gz: $(sources)
	$(call cmd,linux,arm64,)
	$(call tar,linux,arm64)

##### DARWIN (MAC) BUILDS #####
darwin: build/darwin_amd64.tar.gz

build/darwin_amd64.tar.gz: $(sources)
	$(call cmd,darwin,amd64,)
	$(call tar,darwin,amd64)

##### WINDOWS BUILDS #####
windows: build/windows_386.zip build/windows_amd64.zip

build/windows_386.zip: $(sources)
	$(call cmd,windows,386,.exe)
	$(call zip,windows,386,.exe)

build/windows_amd64.zip: $(sources)
	$(call cmd,windows,amd64,.exe)
	$(call zip,windows,amd64,.exe)

##### DEV BUILDS #####
dev: build-dev/darwin_amd64.tar.gz

build-dev/darwin_amd64.tar.gz:
	$(call cmd,darwin,amd64,)
