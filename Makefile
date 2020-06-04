.PHONY: clean build push deploy all

clean:
	rm -rf build

build:
	GOOS=linux GOARCH=amd64 go build -o build/envserver ./cmd/envserver
	docker build -t telliottio/envserver:tilt ./build -f Dockerfile

deploy:
	kubectl apply -k deployment
	kubectl rollout restart deployment/envserver --namespace envserver

all: clean push deploy