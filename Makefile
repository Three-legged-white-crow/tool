.PHONY: build

build:
	mkdir "cmd"
	GOOS=linux GOARCH=amd64
	go build -o cmd/checksum checksum/main.go
	@echo "build tool: checksum"
	go build -o cmd/path_clean path_clean/main.go
	@echo "build tool: path_clean"


build-clean:
	rm -rf ./cmd


docker: docker-centos7 docker-debian-slim


docker-centos7:
	@echo "build image with centos7 base"
	cp Dockerfile-centos7 Dockerfile
	docker build -t tool:centos7 --rm . --no-cache
	rm -f Dockerfile


docker-debian-slim:
	@echo "build image with debian-stable-slim base"
	cp Dockerfile-debian-slim Dockerfile
	docker build -t tool:debian-slim --rm . --no-cache
	rm -f Dockerfile