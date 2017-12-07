#! /usr/bin/make
ifeq ($(OS),Windows_NT)
	BUILD_TARGET_FILES = image-changer.exe main.go
else
	BUILD_TARGET_FILES ?= iamge-changer main.go
endif

all: cleandep depend build

prepare: cleandep depend

depend:
	@glide install

cleandep:
	@rm -rf vendor

build:
	@go build -o $(BUILD_TARGET_FILES)

origin:
	@go run main.go origin

extention:
	@go run main.go extention --ext=$(ext)

color:
	@go run main.go color --type=$(type)

mosaic:
	@go run main.go mosaic --particle=$(particle)

clean:
	@rm -rf destination/*