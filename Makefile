
.PHONY: build
build:
	docker build -t "docker-go" .

.PHONY: run
run:
	#docker run --name docker-go -d docker-go
	docker run --name docker-go -d --mount type=bind,source="$(CURDIR)",target=/app docker-go

.PHONY: login
login:
	docker exec -it docker-go /bin/bash

.PHONY: stop
stop:
	docker stop docker-go

.PHONY: remove
remove: stop
	docker rm docker-go
