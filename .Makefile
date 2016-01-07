BIN=bin
DIST=dist

PKG=github.com/gabesullice/todo
NAME=todo

WINOS=windows
MACOS=darwin
LNXOS=linux

ARCH=amd64

README=README.md

release: $(DIST)/$(NAME)-$(TODO_VERSION)-windows-$(ARCH).zip $(DIST)/$(NAME)-$(TODO_VERSION)-darwin-$(ARCH).tar.gz $(DIST)/$(NAME)-$(TODO_VERSION)-linux-$(ARCH).tar.gz

clean:
	rm -rf {$(BIN),$(DIST)}/*

$(DIST)/$(NAME)-$(TODO_VERSION)-windows-$(ARCH).zip: $(BIN)/$(NAME)-$(TODO_VERSION)-windows-$(ARCH).exe
	zip $@ $? $(README)

$(DIST)/$(NAME)-$(TODO_VERSION)-darwin-$(ARCH).tar.gz: $(BIN)/$(NAME)-$(TODO_VERSION)-darwin-$(ARCH)
	tar -cvzf $@ $? $(README)

$(DIST)/$(NAME)-$(TODO_VERSION)-linux-$(ARCH).tar.gz: $(BIN)/$(NAME)-$(TODO_VERSION)-linux-$(ARCH)
	tar -cvzf $@ $? $(README)

$(BIN)/$(NAME)-$(TODO_VERSION)-windows-$(ARCH).exe:
	env GOOS=$(WINOS) GOARCH=$(ARCH) go build -v -o $@ $(PKG)

$(BIN)/$(NAME)-$(TODO_VERSION)-darwin-$(ARCH):
	env GOOS=$(MACOS) GOARCH=$(ARCH) go build -v -o $@ $(PKG)

$(BIN)/$(NAME)-$(TODO_VERSION)-linux-$(ARCH):
	env GOOS=$(LNXOS) GOARCH=$(ARCH) go build -v -o $@ $(PKG)
