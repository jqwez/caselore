TARGET=Vetshore

all: tests build

build: 
#	@tailwindcss -o ./public/css/styles.css -c ./tailwind.config.js
	@templ generate view
	@go build -o bin/$(TARGET) main.go

nrun: build run

run:
	@bin/$(TARGET)

tests: test

test:
	@go test ./tests/...

testV:
	go test -v ./tests/...

dev:
	@air -c ".air.toml"

clean:
	@rm bin/*
	@rmdir bin

.PHONY: build nrun run test tests dev clean
