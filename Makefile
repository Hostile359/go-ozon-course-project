.PHONY: run
run:
	go run cmd/bot/main.go

build:
	go build -o bin/bot cmd/bot/main.go

prepare:
	mkdir -p config
	echo 'package config\n\nconst ApiKey = "$(APIKEY)"\n' > config/apikey.go

clean:
	rm -rf bin
	rm config/apikey.go
