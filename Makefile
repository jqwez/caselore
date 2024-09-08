TARGET=Vetshore

all: build

build: 
#	@tailwindcss -o ./public/css/styles.css -c ./tailwind.config.js
	@templ generate view
	@go build -o bin/$(TARGET) main.go

nrun: build run

run:
	@bin/$(TARGET)

dev:
	@air -c ".air.toml"

clean:
	@rm bin/*
	@rmdir bin

.PHONY: build nrun run dev clean
