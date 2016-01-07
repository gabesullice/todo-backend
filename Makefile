PROJECT_NAME=todo-backend
BUILD_ROOT=$(PROJECT_NAME)-$(TRAVIS_BRANCH)-$(GOOS).$(GOARCH)
DIST_DIR=dist/$(BUILD_ROOT)

ifeq ($(GOOS),windows)
  BUILD_NAME=$(BUILD_ROOT).exe
	COMPRESS_CMD=zip -r
	COMPRESS_EXT=.zip
else
	BUILD_NAME=$(BUILD_ROOT)
	COMPRESS_CMD=tar -cvzf
	COMPRESS_EXT=.tar.gz
endif

releases: dist_dirs
	go build -v -o "$(DIST_DIR)/$(BUILD_NAME)" github.com/gabesullice/todo-backend
	cp README.md "$(DIST_DIR)"
	$(COMPRESS_CMD) $(DIST_DIR)$(COMPRESS_EXT) $(DIST_DIR)

dist_dirs:
	mkdir -p $(DIST_DIR)
