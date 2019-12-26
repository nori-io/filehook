GOCMD=go
GOBUILD=$(GOCMD) build -buildmode=plugin
GOCLEAN=$(GOCMD) clean
BINARY_NAME=nori_hookfile.so


ifndef PLUGIN_DIR # to allow PLUGIN_DIR to be given as args (see CI)
	DIR=/home/anita/.config/nori
	PLUGIN_DIR=$(DIR)/plugins
endif

.PHONY: all build clean

all: build
build:
	mkdir -p $(PLUGIN_DIR);
	$(GOBUILD) -o $(PLUGIN_DIR)/$(BINARY_NAME) -v
clean:
	$(GOCLEAN)
	rm -f $(PLUGIN_DIR)/$(BINARY_NAME)