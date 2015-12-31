BIN=bin
DIST=dist

PKG=github.com/gabesullice/todo
NAME=todo

WINOS=windows
MACOS=darwin
LNXOS=linux

ARCH=amd64

README=README.md

version=0.2.1

release: $(DIST)/$(NAME)-$(version)-windows-$(ARCH).zip $(DIST)/$(NAME)-$(version)-darwin-$(ARCH).tar.gz $(DIST)/$(NAME)-$(version)-linux-$(ARCH).tar.gz

clean:
	rm -rf {$(BIN),$(DIST)}/*

$(DIST)/$(NAME)-$(version)-windows-$(ARCH).zip: $(BIN)/$(NAME)-$(version)-windows-$(ARCH).exe
	zip $@ $? $(README)

$(DIST)/$(NAME)-$(version)-darwin-$(ARCH).tar.gz: $(BIN)/$(NAME)-$(version)-darwin-$(ARCH)
	tar -cvzf $@ $? $(README)

$(DIST)/$(NAME)-$(version)-linux-$(ARCH).tar.gz: $(BIN)/$(NAME)-$(version)-linux-$(ARCH)
	tar -cvzf $@ $? $(README)

$(BIN)/$(NAME)-$(version)-windows-$(ARCH).exe:
	env GOOS=$(WINOS) GOARCH=$(ARCH) go build -v -o $@ $(PKG)

$(BIN)/$(NAME)-$(version)-darwin-$(ARCH):
	env GOOS=$(MACOS) GOARCH=$(ARCH) go build -v -o $@ $(PKG)

$(BIN)/$(NAME)-$(version)-linux-$(ARCH):
	env GOOS=$(LNXOS) GOARCH=$(ARCH) go build -v -o $@ $(PKG)
