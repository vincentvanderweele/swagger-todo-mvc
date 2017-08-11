PKGS = $(shell go list ./... | grep -v /vendor/)

update-deps:
	godep save $(PKGS)

generate-server:
	go-server-generator ./swagger.yaml

docker-build:
	make generate-server && docker build -t vincentvanderweele/swagger-todo-mvc .

docker-run:
	docker run --rm -it \
	-p 9999:9999 \
	-v $(shell pwd):/go/src/github.com/vincentvanderweele/swagger-todo-mvc \
	vincentvanderweele/swagger-todo-mvc \
	fresh

run-swagger-ui:
	docker run --rm -it \
	-p 80:8080 \
	-e "API_URL=http://localhost:9999/swagger" \
	swaggerapi/swagger-ui

.PHONY: update-deps generate-server docker-build docker-run run-swagger-ui
