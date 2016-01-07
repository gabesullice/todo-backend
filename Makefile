PROJECT_NAME=todo-backend
BUILD_ROOT=$(PROJECT_NAME)-$(TRAVIS_BRANCH)-$(GOOS).$(GOARCH)
DIST_DIR=dist/$(BUILD_ROOT)

ifeq ($(GOOS),windows)
  BUILD_NAME=$(BUILD_ROOT).exe
else
	BUILD_NAME=$(BUILD_ROOT)
endif

releases: dist_dirs
	go build -v -o "$(DIST_DIR)/$(BUILD_NAME)" github.com/gabesullice/todo-backend
	cp README.md "$(DIST_DIR)"
	tar -cvzf $(DIST_DIR).tar.gz $(DIST_DIR)

dist_dirs:
	mkdir -p $(DIST_DIR)
