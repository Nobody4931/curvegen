SRC_FILES := main.go util.go poly.go algo.go
OUT_DIR := bin
OUT_FILE := curvegen
OUT_EXT :=

RMDIR := rm -rfd

ifeq ($(OS),Windows_NT)
	OUT_EXT := .exe

	RMDIR := rmdir /S /Q
endif


all: build

build: $(SRC_FILES)
	go build -o $(OUT_DIR)/$(OUT_FILE)$(OUT_EXT) $^

clean:
	$(RMDIR) $(OUT_DIR)
