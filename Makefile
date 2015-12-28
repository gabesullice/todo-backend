BIN=bin
DIST=dist
PKG=github.com/gabesullice/todo
README=README.md
WINOS=windows
MACOS=darwin
LNXOS=linux
ARCH=amd64

version=0.0.1

release: $(DIST)/$(version)-windows-$(ARCH).zip $(DIST)/$(version)-darwin-$(ARCH).tar.gz $(DIST)/$(version)-linux-$(ARCH).tar.gz

$(DIST)/$(version)-windows-$(ARCH).zip: $(BIN)/$(version)-windows-$(ARCH).exe
	zip $@ $? $(README)

$(DIST)/$(version)-darwin-$(ARCH).tar.gz: $(BIN)/$(version)-darwin-$(ARCH)
	tar -cvzf $@ $? $(README)

$(DIST)/$(version)-linux-$(ARCH).tar.gz: $(BIN)/$(version)-linux-$(ARCH)
	tar -cvzf $@ $? $(README)

$(BIN)/$(version)-windows-$(ARCH).exe:
	env GOOS=$(WINOS) GOARCH=$(ARCH) go build -v -o $@ $(PKG)

$(BIN)/$(version)-darwin-$(ARCH):
	env GOOS=$(MACOS) GOARCH=$(ARCH) go build -v -o $@ $(PKG)

$(BIN)/$(version)-linux-$(ARCH):
	env GOOS=$(LNXOS) GOARCH=$(ARCH) go build -v -o $@ $(PKG)
